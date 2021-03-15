package s3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	c      *s3.S3
	bucket string
	expire time.Duration
	proxy  bool
}

func NewS3(s3AccessKeyId string, s3SecretAccessKey string, s3Endpoint string, s3Region string, s3Bucket string, s3PresignExpire time.Duration, s3ProxyData bool) *S3 {
	conf := &aws.Config{
		Credentials: credentials.NewStaticCredentials(s3AccessKeyId, s3SecretAccessKey, ""),
		Endpoint:    aws.String(s3Endpoint),
		Region:      aws.String(s3Region),
	}
	return &S3{
		c:      s3.New(session.New(conf)),
		bucket: s3Bucket,
		expire: s3PresignExpire,
		proxy:  s3ProxyData,
	}
}

func (s *S3) keyExists(k string) bool {
	conf := &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(k),
	}

	_, err := s.c.HeadObject(conf)
	return err == nil
}

func (s *S3) Name() string {
	return "S3"
}

func (s *S3) List() ([]string, error) {
	conf := &s3.ListObjectsInput{
		Bucket: aws.String(s.bucket),
	}

	rv := []string{}
	if err := s.c.ListObjectsPages(conf, func(fl *s3.ListObjectsOutput, last bool) bool {
		for _, f := range fl.Contents {
			k := *f.Key
			if filepath.Ext(k) == ".json" {
				rv = append(rv, k[:len(k)-5])
			}
		}
		return true
	}); err != nil {
		return nil, err
	}

	return rv, nil
}

func (s *S3) ReadJSON(id string, v interface{}) error {
	conf := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id + ".json"),
	}

	res, err := s.c.GetObject(conf)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == s3.ErrCodeNoSuchKey {
				return os.ErrNotExist
			}
		}
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(v)
}

func (s *S3) WriteJSON(id string, v interface{}) error {
	if s.keyExists(id + ".json") {
		return os.ErrExist
	}

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	conf := &s3.PutObjectInput{
		Body:   bytes.NewReader(data),
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id + ".json"),
	}

	_, err = s.c.PutObject(conf)
	return err
}

func (s *S3) DeleteJSON(id string) error {
	conf := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id + ".json"),
	}

	_, err := s.c.DeleteObject(conf)
	return err
}

func (s *S3) OpenData(id string) (io.ReadCloser, error) {
	conf := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
	}

	res, err := s.c.GetObject(conf)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == s3.ErrCodeNoSuchKey {
				return nil, os.ErrNotExist
			}
		}
		return nil, err
	}
	return res.Body, nil
}

func (s *S3) WriteData(id string, r io.ReadSeeker) (int64, error) {
	if s.keyExists(id) {
		return 0, os.ErrExist
	}

	conf := &s3.PutObjectInput{
		Body:   r,
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
	}

	if _, err := s.c.PutObject(conf); err != nil {
		return 0, err
	}

	return r.Seek(0, io.SeekCurrent)
}

func (s *S3) DeleteData(id string) error {
	conf := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
	}

	_, err := s.c.DeleteObject(conf)
	return err
}

func (s *S3) serveDataHead(w http.ResponseWriter, r *http.Request, id string, contentType string, filename string, attachment bool) error {
	conf := &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
	}
	if v, ok := r.Header[textproto.CanonicalMIMEHeaderKey("If-Match")]; ok && len(v) > 0 {
		conf.IfMatch = aws.String(v[0])
	}
	if v, ok := r.Header[textproto.CanonicalMIMEHeaderKey("If-Modified-Since")]; ok && len(v) > 0 {
		if t, err := http.ParseTime(v[0]); err == nil {
			conf.IfModifiedSince = aws.Time(t)
		}
	}
	if v, ok := r.Header[textproto.CanonicalMIMEHeaderKey("If-None-Match")]; ok && len(v) > 0 {
		conf.IfNoneMatch = aws.String(v[0])
	}
	if v, ok := r.Header[textproto.CanonicalMIMEHeaderKey("If-Unmodified-Since")]; ok && len(v) > 0 {
		if t, err := http.ParseTime(v[0]); err == nil {
			conf.IfUnmodifiedSince = aws.Time(t)
		}
	}
	if v, ok := r.Header[textproto.CanonicalMIMEHeaderKey("Range")]; ok && len(v) > 0 {
		conf.Range = aws.String(v[0])
	}

	req, o := s.c.HeadObjectRequest(conf)
	if err := req.Send(); err != nil {
		if _, ok := err.(awserr.RequestFailure); !ok {
			return err
		}
	}

	if v := o.AcceptRanges; v != nil && *v != "" {
		w.Header().Set("Accept-Ranges", *v)
	}
	if v := o.CacheControl; v != nil && *v != "" {
		w.Header().Set("Cache-Control", *v)
	}
	if v := o.ContentEncoding; v != nil && *v != "" {
		w.Header().Set("Content-Encoding", *v)
	}
	if v := o.ContentLanguage; v != nil && *v != "" {
		w.Header().Set("Content-Language", *v)
	}
	if v := o.ContentLength; v != nil && *v > 0 {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", *v))
	}
	if v := o.ETag; v != nil && *v != "" {
		w.Header().Set("ETag", *v)
	}
	if v := o.LastModified; v != nil && !(*v).IsZero() {
		w.Header().Set("Last-Modified", (*v).UTC().Format(http.TimeFormat))
	}

	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}
	if filename != "" {
		if attachment {
			w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
		} else {
			w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, filename))
		}
	}

	w.WriteHeader(req.HTTPResponse.StatusCode)
	return nil
}

func (s *S3) serveDataGet(w http.ResponseWriter, r *http.Request, id string, contentType string, filename string, attachment bool) error {
	conf := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
	}
	if v, ok := r.Header[textproto.CanonicalMIMEHeaderKey("If-Match")]; ok && len(v) > 0 {
		conf.IfMatch = aws.String(v[0])
	}
	if v, ok := r.Header[textproto.CanonicalMIMEHeaderKey("If-Modified-Since")]; ok && len(v) > 0 {
		if t, err := http.ParseTime(v[0]); err == nil {
			conf.IfModifiedSince = aws.Time(t)
		}
	}
	if v, ok := r.Header[textproto.CanonicalMIMEHeaderKey("If-None-Match")]; ok && len(v) > 0 {
		conf.IfNoneMatch = aws.String(v[0])
	}
	if v, ok := r.Header[textproto.CanonicalMIMEHeaderKey("If-Unmodified-Since")]; ok && len(v) > 0 {
		if t, err := http.ParseTime(v[0]); err == nil {
			conf.IfUnmodifiedSince = aws.Time(t)
		}
	}
	if v, ok := r.Header[textproto.CanonicalMIMEHeaderKey("Range")]; ok && len(v) > 0 {
		conf.Range = aws.String(v[0])
	}

	req, o := s.c.GetObjectRequest(conf)
	if err := req.Send(); err != nil {
		if _, ok := err.(awserr.RequestFailure); !ok {
			return err
		}
	}

	if v := o.AcceptRanges; v != nil && *v != "" {
		w.Header().Set("Accept-Ranges", *v)
	}
	if v := o.CacheControl; v != nil && *v != "" {
		w.Header().Set("Cache-Control", *v)
	}
	if v := o.ContentEncoding; v != nil && *v != "" {
		w.Header().Set("Content-Encoding", *v)
	}
	if v := o.ContentLanguage; v != nil && *v != "" {
		w.Header().Set("Content-Language", *v)
	}
	if v := o.ContentLength; v != nil && *v > 0 {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", *v))
	}
	if v := o.ContentRange; v != nil && *v != "" {
		w.Header().Set("Content-Range", *v)
	}
	if v := o.ETag; v != nil && *v != "" {
		w.Header().Set("ETag", *v)
	}
	if v := o.Expires; v != nil && *v != "" {
		w.Header().Set("Expires", *v)
	}
	if v := o.LastModified; v != nil && !(*v).IsZero() {
		w.Header().Set("Last-Modified", (*v).UTC().Format(http.TimeFormat))
	}

	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}
	if filename != "" {
		if attachment {
			w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
		} else {
			w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, filename))
		}
	}

	w.WriteHeader(req.HTTPResponse.StatusCode)

	if v := o.ContentLength; v != nil && *v > 0 {
		io.CopyN(w, o.Body, *o.ContentLength)
	} else {
		io.Copy(w, o.Body)
	}

	return o.Body.Close()
}

func (s *S3) redirectDataGet(w http.ResponseWriter, r *http.Request, id string, contentType string, filename string, attachment bool) error {
	conf := &s3.GetObjectInput{
		Bucket:              aws.String(s.bucket),
		Key:                 aws.String(id),
		ResponseContentType: aws.String(contentType),
	}
	if filename != "" {
		if attachment {
			conf.ResponseContentDisposition = aws.String(fmt.Sprintf(`attachment; filename="%s"`, filename))
		} else {
			conf.ResponseContentDisposition = aws.String(fmt.Sprintf(`inline; filename="%s"`, filename))
		}
	}

	req, _ := s.c.GetObjectRequest(conf)
	requrl, err := req.Presign(s.expire)
	if err != nil {
		return err
	}

	http.Redirect(w, r, requrl, http.StatusFound)
	return nil
}

func (s *S3) ServeData(w http.ResponseWriter, r *http.Request, id string, contentType string, filename string, attachment bool) error {
	switch r.Method {
	case http.MethodHead:
		// HEAD requests are always proxied
		return s.serveDataHead(w, r, id, contentType, filename, attachment)

	case http.MethodGet:
		if s.proxy {
			return s.serveDataGet(w, r, id, contentType, filename, attachment)
		}
		return s.redirectDataGet(w, r, id, contentType, filename, attachment)

	default:
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return nil
	}
}

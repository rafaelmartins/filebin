package s3

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"os"
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
			if k := f.Key; k != nil {
				rv = append(rv, *k)
			}
		}
		return true
	}); err != nil {
		return nil, err
	}

	return rv, nil
}

func (s *S3) Read(id string) (io.ReadCloser, error) {
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

func (s *S3) ReadMetadata(id string) (string, string, int64, time.Time, error) {
	conf := &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
	}

	res, err := s.c.HeadObject(conf)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == s3.ErrCodeNoSuchKey {
				return "", "", 0, time.Time{}, os.ErrNotExist
			}
		}
		return "", "", 0, time.Time{}, err
	}

	filename := ""
	if v, ok := res.Metadata[textproto.CanonicalMIMEHeaderKey("filename")]; ok && v != nil {
		filenameB, err := base64.URLEncoding.DecodeString(*v)
		if err != nil {
			return "", "", 0, time.Time{}, err
		}
		filename = string(filenameB)
	}

	mimetype := ""
	if v, ok := res.Metadata[textproto.CanonicalMIMEHeaderKey("mimetype")]; ok && v != nil {
		mimetype = *v
	}

	size := int64(0)
	if v := res.ContentLength; v != nil {
		size = *v
	}

	timestamp := time.Time{}
	if v, ok := res.Metadata[textproto.CanonicalMIMEHeaderKey("timestamp")]; ok && v != nil {
		if err := timestamp.UnmarshalText([]byte(*v)); err != nil {
			return "", "", 0, time.Time{}, err
		}
	}

	return filename, mimetype, size, timestamp, nil
}

func (s *S3) Write(id string, r io.ReadSeeker, filename string, mimetype string) (int64, error) {
	if s.keyExists(id) {
		return 0, os.ErrExist
	}

	filename = base64.URLEncoding.EncodeToString([]byte(filename))

	ts, err := time.Now().UTC().MarshalText()
	if err != nil {
		return 0, err
	}

	conf := &s3.PutObjectInput{
		Body:   r,
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
		Metadata: map[string]*string{
			textproto.CanonicalMIMEHeaderKey("filename"):  aws.String(filename),
			textproto.CanonicalMIMEHeaderKey("mimetype"):  aws.String(mimetype),
			textproto.CanonicalMIMEHeaderKey("timestamp"): aws.String(string(ts)),
		},
	}

	if _, err := s.c.PutObject(conf); err != nil {
		return 0, err
	}

	return r.Seek(0, io.SeekCurrent)
}

func (s *S3) Delete(id string) error {
	conf := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
	}

	if _, err := s.c.DeleteObject(conf); err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == s3.ErrCodeNoSuchKey {
				return os.ErrNotExist
			}
		}
		return err
	}

	return nil
}

func handleError(w http.ResponseWriter, r *http.Request, err error) error {
	if err != nil {
		if aerr, ok := err.(awserr.RequestFailure); ok {
			if h := aerr.StatusCode(); h == http.StatusNotModified {
				w.WriteHeader(h)
				return nil
			}
		}
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == s3.ErrCodeNoSuchKey {
				http.NotFound(w, r)
				return nil
			}
		}
	}

	return err
}

func (s *S3) serveDataHead(w http.ResponseWriter, r *http.Request, id string, filename string, mimetype string, attachment bool) error {
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
		return handleError(w, r, err)
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

	if mimetype != "" {
		w.Header().Set("Content-Type", mimetype)
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

func (s *S3) serveDataGet(w http.ResponseWriter, r *http.Request, id string, filename string, mimetype string, attachment bool) error {
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
		return handleError(w, r, err)
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

	if mimetype != "" {
		w.Header().Set("Content-Type", mimetype)
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

func (s *S3) redirectDataGet(w http.ResponseWriter, r *http.Request, id string, filename string, mimetype string, attachment bool) error {
	conf := &s3.GetObjectInput{
		Bucket:              aws.String(s.bucket),
		Key:                 aws.String(id),
		ResponseContentType: aws.String(mimetype),
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

func (s *S3) Serve(w http.ResponseWriter, r *http.Request, id string, filename string, mimetype string, attachment bool) error {
	switch r.Method {
	case http.MethodHead:
		// HEAD requests are always proxied
		return s.serveDataHead(w, r, id, filename, mimetype, attachment)

	case http.MethodGet:
		if s.proxy {
			return s.serveDataGet(w, r, id, filename, mimetype, attachment)
		}
		return s.redirectDataGet(w, r, id, filename, mimetype, attachment)

	default:
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return nil
	}
}

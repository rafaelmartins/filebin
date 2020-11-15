package highlight

import (
	"errors"
	"mime"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
)

var (
	ErrUnsupported = errors.New("highlight: unsupported")
)

func GetLexer(mimetype string) string {
	l := func() chroma.Lexer {
		mt, _, err := mime.ParseMediaType(mimetype)
		if err != nil {
			return nil
		}

		// exceptions
		if mt == "text/plain" {
			return lexers.Get("plaintext")
		}

		return lexers.MatchMimeType(mt)
	}()

	if l == nil {
		return ""
	}

	return l.Config().Name
}

func DetectMimetype(filename string) (string, error) {
	lexer := lexers.Match(filename)
	if lexer == nil {
		return "", ErrUnsupported
	}
	mimes := lexer.Config().MimeTypes
	if len(mimes) == 0 {
		return "", ErrUnsupported
	}
	return mimes[0], nil
}

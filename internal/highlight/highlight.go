package highlight

import (
	"errors"
	"mime"
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
)

var (
	ErrUnsupported = errors.New("highlight: unsupported")
)

func GetLexer(mimetype string) (chroma.Lexer, error) {
	mt, _, err := mime.ParseMediaType(mimetype)
	if err != nil {
		return nil, ErrUnsupported
	}

	// exceptions
	if mt == "text/plain" {
		return lexers.Get("plaintext"), nil
	}

	l := lexers.MatchMimeType(mt)
	if l != nil {
		return l, nil
	}

	// if we failed to find a lexer for a text type, use plaintext
	if strings.HasPrefix(mt, "text/") {
		return lexers.Get("plaintext"), nil
	}

	return nil, ErrUnsupported
}

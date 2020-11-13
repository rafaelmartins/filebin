package highlight

import (
	"mime"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
)

func GetMimeType(filename string) string {
	lexer := lexers.Match(filename)
	if lexer == nil {
		return ""
	}
	mimes := lexer.Config().MimeTypes
	if len(mimes) == 0 {
		return ""
	}
	return mimes[0]
}

func GetLexer(mimetype string) chroma.Lexer {
	mt, _, err := mime.ParseMediaType(mimetype)
	if err != nil {
		return nil
	}

	// exceptions
	if mt == "text/plain" {
		return lexers.Get("plaintext")
	}

	return lexers.MatchMimeType(mt)
}

func Validate(mimetype string) bool {
	return GetLexer(mimetype) != nil
}

package highlight

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/styles"
	"github.com/rafaelmartins/filebin/internal/models"
)

var (
	ErrNoLexer  = errors.New("highlight: no lexer found")
	ErrNotFound = errors.New("highlight: file not found")
)

func Highlight(w http.ResponseWriter, r *http.Request, stylename string, filename string, fd *models.FileData) error {
	lexer := GetLexer(fd.Mimetype)
	if lexer == nil {
		return ErrNoLexer
	}

	style := styles.Get(stylename)
	if style == nil {
		style = styles.Fallback
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNotFound
		}
		return err
	}

	iterator, err := lexer.Tokenise(nil, string(contents))
	if err != nil {
		return err
	}

	formatter := html.New(
		html.Standalone(true),
		html.WithLineNumbers(true),
		html.LinkableLineNumbers(true, "L"),
		html.LineNumbersInTable(true),
	)

	return formatter.Format(w, style, iterator)
}

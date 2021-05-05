package highlight

import (
	"html/template"
	"io"
	"net/http"

	"github.com/rafaelmartins/filebin/internal/filedata"
	"github.com/rafaelmartins/filebin/internal/highlight"
)

var (
	tmplTitle = template.Must(template.New("title").Parse(
		`<title>filebin — {{.GetFilename}}</title>
`))
	tmplDetails = template.Must(template.New("details").Parse(
		`<strong>File:</strong> {{.Fd.GetFilename}} |
<strong>Language:</strong> {{.Lexer}} |
<strong>Created on:</strong> {{.Timestamp}} |
<a href="/{{.Fd.GetId}}.txt">Plain text</a> |
<a href="/{{.Fd.GetId}}/download">Download</a>
<br>
`))
)

func highlightFile(w http.ResponseWriter, fd *filedata.FileData) error {
	lexer, err := highlight.GetLexer(fd.Mimetype)
	if err != nil {
		return err
	}

	_, err = io.WriteString(w, `<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
`)
	if err != nil {
		return err
	}

	if err := tmplTitle.Execute(w, fd); err != nil {
		return err
	}

	if _, err := io.WriteString(w, "<style type=\"text/css\">\n"); err != nil {
		return err
	}

	if err := highlight.GenerateCSS(w); err != nil {
		return err
	}

	_, err = io.WriteString(w, `</style>
</head>
<body>
`)
	if err != nil {
		return err
	}

	fp, err := fd.Read()
	if err != nil {
		return err
	}
	defer fp.Close()

	if err := highlight.GenerateHTML(w, fp, lexer); err != nil {
		return err
	}

	d := struct {
		Fd        *filedata.FileData
		Lexer     string
		Timestamp string
	}{
		Fd:        fd,
		Lexer:     lexer.Config().Name,
		Timestamp: fd.Timestamp.Format("02-01-2006 15:04:05"),
	}
	if err := tmplDetails.Execute(w, d); err != nil {
		return err
	}

	_, err = io.WriteString(w, `</body>
</html>
`)
	return err
}

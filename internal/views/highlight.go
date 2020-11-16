package views

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
		`<strong>File:</strong> {{.GetFilename}} |
<strong>Language:</strong> {{.GetLexer}} |
<a href="/{{.GetId}}.txt">Plain text</a> |
<a href="/download/{{.GetId}}">Download</a>
<br>
`))
)

func highlightFile(w http.ResponseWriter, fd *filedata.FileData) error {
	_, err := io.WriteString(w, `<!DOCTYPE html>
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

	if _, err := io.WriteString(w, `<style type="text/css">`); err != nil {
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

	if err := fd.GenerateHTML(w); err != nil {
		return err
	}

	if err := tmplDetails.Execute(w, fd); err != nil {
		return err
	}

	_, err = io.WriteString(w, `</body>
</html>
`)
	return err
}

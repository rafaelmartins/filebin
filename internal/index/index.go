package index

import (
	"fmt"
	"net/http"

	"github.com/rafaelmartins/filebin/internal/version"
)

var (
	logo = `  __ _ _      _     _
 / _(_) | ___| |__ (_)_ __
| |_| | |/ _ \ '_ \| | '_ \
|  _| | |  __/ |_) | | | | |
|_| |_|_|\___|_.__/|_|_| |_|
`
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	fmt.Fprintf(w, "%s\n", logo)
	fmt.Fprintf(w, "Version %s, running at %s\n\n", version.Version, r.Host)
	fmt.Fprintf(w, "Source code: https://github.com/rafaelmartins/filebin\n")
}

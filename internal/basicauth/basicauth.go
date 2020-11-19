package basicauth

import (
	"crypto/subtle"
	"fmt"
	"net/http"

	"github.com/rafaelmartins/filebin/internal/settings"
)

func BasicAuth(w http.ResponseWriter, r *http.Request) bool {
	s, err := settings.Get()
	if err != nil {
		return false
	}

	if u, p, ok := r.BasicAuth(); ok {
		c1 := subtle.ConstantTimeCompare([]byte(u), []byte(s.AuthUsername)) == 1
		c2 := subtle.ConstantTimeCompare([]byte(p), []byte(s.AuthPassword)) == 1
		if c1 && c2 && s.AuthUsername != "" && s.AuthPassword != "" {
			return true
		}
	}
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, s.AuthRealm))
	w.WriteHeader(http.StatusUnauthorized)
	return false
}

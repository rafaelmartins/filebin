package basicauth

import (
	"crypto/subtle"
	"fmt"
	"net/http"
)

func BasicAuth(w http.ResponseWriter, r *http.Request, realm string, username string, password string) bool {
	if u, p, ok := r.BasicAuth(); ok {
		c1 := subtle.ConstantTimeCompare([]byte(u), []byte(username)) == 1
		c2 := subtle.ConstantTimeCompare([]byte(p), []byte(password)) == 1
		if c1 && c2 && username != "" && password != "" {
			return true
		}
	}
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
	return false
}

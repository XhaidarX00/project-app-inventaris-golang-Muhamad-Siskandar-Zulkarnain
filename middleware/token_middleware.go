package middlewaree

import (
	"main/library"
	"main/service"
	"net/http"
)

// Middleware untuk validasi token
func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		if token == "" {
			library.StrucToJson(w, library.UnauthorizedRequest)
			return
		}

		if err := service.ServiceF.TokenCheck(token); err != "" {
			library.StrucToJson(w, library.UnauthorizedRequest)
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		next.ServeHTTP(w, r)
	})
}

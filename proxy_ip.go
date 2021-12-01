package handlers

import (
	"net/http"
)

func ProxyIPHeadersHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Set the remote IP with the value passed from the proxy.
		if fwd := getIP(r); fwd != "" {
			r.RemoteAddr = fwd
		}

		// Call the next handler in the chain.
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

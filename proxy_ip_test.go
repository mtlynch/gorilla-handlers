package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test the middleware end-to-end
func TestProxyIPHeadersHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	r := newRequest("GET", "/")

	r.Header.Set(xForwardedFor, "8.8.8.8")
	r.Header.Set(xForwardedProto, "https")
	r.Header.Set(xForwardedHost, "google.com")
	var (
		addr string
	)
	ProxyIPHeadersHandler(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			addr = r.RemoteAddr
		})).ServeHTTP(rr, r)

	if rr.Code != http.StatusOK {
		t.Fatalf("bad status: got %d want %d", rr.Code, http.StatusOK)
	}

	if addr != r.Header.Get(xForwardedFor) {
		t.Fatalf("wrong address: got %s want %s", addr,
			r.Header.Get(xForwardedFor))
	}

}

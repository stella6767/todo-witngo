package util

import (
	"net/http"
)

func isHTMX(r *http.Request) bool {
	// Check for "HX-Request" header
	if r.Header.Get("HX-Request") != "" {
		return true
	}
	return false
}

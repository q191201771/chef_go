package http

import "net/http"

func Post(path string, formData map[string]string, headers map[string]string, cookies map[string]string) (string, map[string]*http.Cookie, error) {
	return fetch("POST", path, formData, headers, cookies)
}

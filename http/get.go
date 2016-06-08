package http

import "net/http"

func Get(path string, queries map[string]string, headers map[string]string, cookies map[string]string) (string, map[string]*http.Cookie, error) {
	return fetch("GET", path, queries, headers, cookies)
}

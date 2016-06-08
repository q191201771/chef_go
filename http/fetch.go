package http

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

/// return body, cookies, error
func fetch(method string, path string, params map[string]string, headers map[string]string, cookies map[string]string) (string, map[string]*http.Cookie, error) {
	var (
		url     string
		reqBody io.Reader
	)

	switch method {
	case "GET":
		url = ComposeURL(path, params)
	case "POST":
		url = path
		reqBody = strings.NewReader(URLEncode(params))
	default:
		return "", nil, fmt.Errorf("method only support GET or POST")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return "", nil, err
	}

	var cookiesKV []string
	if len(cookies) > 0 {
		for k, v := range cookies {
			kv := fmt.Sprintf("%s=%s", k, v)
			cookiesKV = append(cookiesKV, kv)
		}
		req.Header.Set("Cookie", strings.Join(cookiesKV, "; "))
	}

	/// some default header,can be overwrite by param[headers]
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", MOCK_UA_MAC_CHROME)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	//defer resp.Body.Close()
	if err != nil {
		return "", nil, err
	}

	replyCookies := make(map[string]*http.Cookie, len(resp.Cookies()))
	for _, c := range resp.Cookies() {
		replyCookies[c.Name] = c
	}

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		/// TODO
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return "", nil, err
		}
		body, err := ioutil.ReadAll(reader)
		return string(body), replyCookies, err
	default:
		body, err := ioutil.ReadAll(resp.Body)
		return string(body), replyCookies, err
	}

}

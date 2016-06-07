package http

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"
)

func Get(path string, queries map[string]string, headers map[string]string) (string, error) {
	url := ComposeURL(path, queries)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	//defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		/// TODO
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return "", err
		}
		body, err := ioutil.ReadAll(reader)
		return string(body), err
	default:
		body, err := ioutil.ReadAll(resp.Body)
		return string(body), err
	}
}

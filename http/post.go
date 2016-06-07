package http

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"strings"
)

func Post(path string, headers map[string]string, formDatas map[string]string) (string, error) {
	client := &http.Client{}
	reqBody := strings.NewReader(URLEncode(formDatas))
	req, err := http.NewRequest("POST", path, reqBody)
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

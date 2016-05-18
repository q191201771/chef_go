package http

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(path string, queries map[string]string, headers map[string]string) (string, error) {
	url := composeURL(path, queries)

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

func marshalQueries(kv map[string]string) string {
	if len(kv) == 0 {
		return ""
	}

	var ret string
	for k, v := range kv {
		pair := k + "=" + v + "&"
		ret += pair
	}
	return ret[:len(ret)-1]
}

func composeURL(path string, queries map[string]string) string {
	if len(queries) == 0 {
		return path
	}
	return fmt.Sprintf(
		"%s?%s",
		path,
		marshalQueries(queries),
	)
}

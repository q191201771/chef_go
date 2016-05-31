package http

import (
	"log"
	"testing"
)

func TestFetch(t *testing.T) {
	log.Println(Get("http://www.baidu.com", nil, nil))
}

func TestComposeUrl(t *testing.T) {
	path := "http://www.baidu.com"
	queries := map[string]string{
		"aaa": "111",
		"bb":  "22",
	}
	urls := map[string]bool{
		"http://www.baidu.com?aaa=111&bb=22": true,
		"http://www.baidu.com?bb=22&aaa=111": true,
	}
	url := ComposeURL(path, queries)
	if _, exist := urls[url]; !exist {
		t.Fatal(url)
	}
}

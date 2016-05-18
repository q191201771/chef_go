package http

import (
	"log"
	"testing"
)

func TestFetch(t *testing.T) {
	log.Println(Get("http://www.baidu.com", nil, nil))
}

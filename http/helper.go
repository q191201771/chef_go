package http

import (
	"fmt"
	"net/url"
)

func ComposeURL(path string, queries map[string]string) string {
	if len(queries) == 0 {
		return path
	}
	return fmt.Sprintf(
		"%s?%s",
		path,
		marshalQueries(queries),
	)
}

func URLEncode(kv map[string]string) string {
	values := url.Values{}
	for k, v := range kv {
		values.Set(k, v)
	}
	return values.Encode()
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

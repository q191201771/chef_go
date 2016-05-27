package main

import (
	"io/ioutil"

	"github.com/q191201771/chef_go/env"
)

func main() {
	env.SetBaseDirLocal()
	ioutil.WriteFile("test.data", []byte("hello"), 0644)
}

package main

import (
	"errors"
	"fmt"

	clog "github.com/q191201771/chef_go/log"
	cpanic "github.com/q191201771/chef_go/panic"
)

func test() {
	panic("!!!test")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			clog.L().Critical(fmt.Sprintf(
				"[chef_go.panic]\n%v\n%s",
				err,
				cpanic.Stack(3),
			))
		}
	}()
	clog.L().Debug("before panic")
	test()
	panic(errors.New("errrors.New"))
	//panic(55)
	//panic(`I'm panic`)
	clog.L().Debug("after panic")
}

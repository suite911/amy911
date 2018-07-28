package onfail

import "log"

type OnFail interface {
	Fail(error)
}

type OnFailCallFunction func(error)

func (onFail OnFailCallFunction) Fail(err error) {
	onFail(err)
}

var Fatal OnFailCallFunction = func(err error) {
	log.Fatalln(err)
}

var Panic OnFailCallFunction = func(err error) {
	panic(err)
}

var Print OnFailCallFunction = func(err error) {
	log.Println(err)
}

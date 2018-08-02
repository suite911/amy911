package onfail

import "log"

// Helper function for configurable fail behavior
func Fail (err error, arg interface{}, callee OnFail, caller ...OnFail) {
	switch {
	case len(caller) >= 1:
		caller[0].Fail(err, arg)
	case callee != nil:
		callee.Fail(err, arg)
	case Default != nil:
		Default.Fail(err, arg)
	default:
		Panic.Fail(err, arg)
	}
}

// Called by the `Fail` helper function when the behavior is not configured at the call site
var Default OnFail = Panic

// Interface for types which can configure fail behavior
type OnFail interface {
	Fail(error, interface{})
}

// Function type for configure fail behavior by calling the function
type OnFailCallFunction func(error, interface{})

// To satisfy the `OnFail` interface
func (onFail OnFailCallFunction) Fail(err error, arg interface{}) {
	onFail(err, arg)
}

// Built-in fail behavior configuration to log fatally
var Fatal OnFailCallFunction = func(err error, arg interface{}) {
	log.Fatalln(err)
}

// Built-in fail behavior configuration to panic
var Panic OnFailCallFunction = func(err error, arg interface{}) {
	panic(err)
}

// Built-in fail behavior configuration to log and continue
var Print OnFailCallFunction = func(err error, arg interface{}) {
	log.Println(err)
}

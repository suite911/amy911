package onfail

import (
	"log"

	"github.com/pkg/errors"
)

// Helper function for configurable fail behavior
func Fail (err, arg interface{}, calleeConf OnFail, callerConf ...OnFail) {
	e := InterfaceToError(err)
	switch {
	case len(callerConf) >= 1:
		callerConf[0].Fail(e, arg)
	case calleeConf != nil:
		calleeConf.Fail(e, arg)
	case Default != nil: // but you should never set Default to nil
		Default.Fail(e, arg)
	case Panic != nil: // but you should never set Panic to nil
		Panic.Fail(e, arg)
	default:
		panic(e)
	}
}

// An internal function; exported in case it is needed in user code
func InterfaceToError(err interface{}) error {
	switch err.(type) {
	case error:
		return err.(error)
	case string:
		return errors.New(err.(string))
	default:
		panic("`err` must be either `error` or `string`!")
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

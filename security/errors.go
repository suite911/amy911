package security

import "errors"

var errNotSupported = errors.New("Not supported")

func ErrNotSupported() error {
	return errNotSupported
}

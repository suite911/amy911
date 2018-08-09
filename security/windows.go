// +build windows

package security

import (
	"os"

	"golang.org/x/sys/unix"

	"github.com/amy911/amy911/onfail"
)

func chroot(path string, onFail interface{}) error {
	err := errNotSupported
	onfail.Fail(err, path, onfail.Panic, onFail)
	return err
}

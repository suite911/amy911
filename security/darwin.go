// +build darwin

package security

import (
	"os"

	"golang.org/x/sys/unix"

	"github.com/amy911/amy911/onfail"
)

func chroot(path string, onFail onfail.OnFail) error {
	if err := unix.Chroot(path); err != nil {
		onfail.Fail(err, path, onfail.Panic, onFail)
		return err
	}
	if err := os.Chdir("/"); err != nil {
		onfail.Fail(err, path, onfail.Panic, onFail)
		return err
	}
	return nil
}

package security

import (
	"github.com/amy911/amy911/onfail"
)

func Chroot(path string, onFail ...onfail.OnFail) error {
	return chroot(path, onFail)
}

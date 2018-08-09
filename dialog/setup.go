package dialog

import (
	"sync"
)

var (
	fClose, fInit func()

	mutex sync.Mutex
)

func Close() {
	mutex.Lock(); defer mutex.Unlock()
	if iLibrary == nil {
		return
	}
	if fClose != nil {
		fClose()
	}
}

func Init() {
	mutex.Lock(); defer mutex.Unlock()
	if iLibrary != nil {
		return
	}
	if fInit != nil {
		fInit()
	}
}

func setUp(fInit_, fClose_ func()) {
	mutex.Lock(); defer mutex.Unlock()
	if iLibrary != nil {
		return
	}
	fClose = fClose_
	fInit = fInit_
}

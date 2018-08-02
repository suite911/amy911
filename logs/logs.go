package logs

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Logs struct {
	Minor *log.Logger // Minor trace point
	Major *log.Logger // Major trace point
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
	Fatal *log.Logger
}

func New(path string, info, debug, trace bool) (logs *Logs, err error) {
	logFile, TODO := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	err = TODO
	if err != nil {
		return
	}

	logs = new(Logs)

	var wTrace, wDebug, wInfo, wWarn, wError, wFatal io.Writer

	wTrace = ioutil.Discard
	if trace {
		wTrace = logFile
	}

	wDebug = ioutil.Discard
	if debug || trace {
		wDebug = logFile
	}

	wInfo = logFile
	if info || debug || trace {
		wInfo = io.MultiWriter(logFile, os.Stdout)
	}

	wWarn = io.MultiWriter(logFile, os.Stderr)
	wError = io.MultiWriter(logFile, os.Stderr)
	wFatal = io.MultiWriter(logFile, os.Stderr)

	const infoFlags = log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC
	const logsFlags = log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC | log.Lshortfile

	logs.Minor = log.New(wTrace, "-t-\t", logsFlags)
	logs.Major = log.New(wTrace, "-tr-\t", logsFlags)
	logs.Debug = log.New(wDebug, "-dbg-\t", logsFlags)
	logs.Info = log.New(wInfo, "-info-\t", infoFlags)
	logs.Warn = log.New(wWarn, "-WARN-\t", logsFlags)
	logs.Error = log.New(wError, "-ERROR-\t", logsFlags)
	logs.Fatal = log.New(wFatal, "-FATAL-\t", logsFlags)
	return
}

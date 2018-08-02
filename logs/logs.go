package logs

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Logs struct {
	LoggerMinor *log.Logger // Minor trace point
	LoggerMajor *log.Logger // Major trace point
	LoggerDebug *log.Logger // Debug info
	LoggerNote  *log.Logger // Note
	LoggerInfo  *log.Logger // Info
	LoggerWarn  *log.Logger // Warning
	LoggerError *log.Logger // Error
	LoggerFatal *log.Logger // Fatal error
}

func New(path string, info, debug, trace bool) (logs *Logs, err error) {
	logs = new(Logs)
	err = logs.Init(path, info, debug, trace)
	return
}

func (logs *Logs) Init(path string, info, debug, trace bool) error {
	if logs == nil {
		panic("logs was nil")
	}
	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

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

	logs.LoggerMinor = log.New(wTrace, "-t-\t", logsFlags)
	logs.LoggerMajor = log.New(wTrace, "-tr-\t", logsFlags)
	logs.LoggerDebug = log.New(wDebug, "-dbg-\t", logsFlags)
	logs.LoggerNote = log.New(logFile, "-note-\t", infoFlags)
	logs.LoggerInfo = log.New(wInfo, "-info-\t", infoFlags)
	logs.LoggerWarn = log.New(wWarn, "-WARN-\t", logsFlags)
	logs.LoggerError = log.New(wError, "-ERROR-\t", logsFlags)
	logs.LoggerFatal = log.New(wFatal, "-FATAL-\t", logsFlags)
	return nil
}

func (logs *Logs) Minor(v ...interface{}) {
	if logs != nil && logs.LoggerMinor != nil {
		logs.LoggerMinor.Println(v...)
	}
}

func (logs *Logs) Major(v ...interface{}) {
	if logs != nil && logs.LoggerMajor != nil {
		logs.LoggerMajor.Println(v...)
	}
}

func (logs *Logs) Debug(v ...interface{}) {
	if logs != nil && logs.LoggerDebug != nil {
		logs.LoggerDebug.Println(v...)
	}
}

func (logs *Logs) Note(v ...interface{}) {
	if logs != nil && logs.LoggerNote != nil {
		logs.LoggerNote.Println(v...)
	}
}

func (logs *Logs) Info(v ...interface{}) {
	if logs != nil && logs.LoggerInfo != nil {
		logs.LoggerInfo.Println(v...)
	}
}

func (logs *Logs) Warn(v ...interface{}) {
	if logs != nil && logs.LoggerWarn != nil {
		logs.LoggerWarn.Println(v...)
	}
}

func (logs *Logs) Error(v ...interface{}) {
	if logs != nil && logs.LoggerError != nil {
		logs.LoggerError.Println(v...)
	}
}

func (logs *Logs) Fatal(v ...interface{}) {
	if logs != nil && logs.LoggerFatal != nil {
		logs.LoggerFatal.Fatalln(v...)
	}
}

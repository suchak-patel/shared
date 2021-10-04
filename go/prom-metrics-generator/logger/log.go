package clog

import (
	"io"
	"log"
)

var (
	Info    *log.Logger
	Error   *log.Logger
)

func Logger(
	infoHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
			"INFO: ",
			log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile)
}

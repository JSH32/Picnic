package logger

import (
	"fmt"
	"time"
)

type loglevel string

const (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	reset  = "\033[0m"
)

const (
	levelErr     loglevel = "ERROR"
	levelPanic            = "PANIC"
	levelOk               = "INFO"
	levelSuccess          = "SUCCESS"
	levelWarn             = "WARNING"
)

const timeformat = "15:04:05"

func formatdate() string {
	tn := time.Now()
	return fmt.Sprintf("[%s%d:%d:%d%s]", green, tn.Hour(), tn.Minute(), tn.Second(), reset)
}

func log(level loglevel, text string) {
	var color string

	switch level {
	case levelErr:
		color = red
	case levelPanic:
		color = red
	case levelOk:
		color = cyan
	case levelSuccess:
		color = green
	case levelWarn:
		color = yellow
	}

	logstring := fmt.Sprintf("%s [%s%s%s] %s", formatdate(), color, level, reset, text)

	fmt.Println(logstring)
}

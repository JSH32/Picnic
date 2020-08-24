package logger

import "os"

// Info : log some text
func Info(text string) {
	log(levelOk, text)
}

// Error : log some text as error
func Error(text string) {
	log(levelErr, text)
}

// Warn : log some text as warning
func Warn(text string) {
	log(levelWarn, text)
}

// Success : log some text as success
func Success(text string) {
	log(levelSuccess, text)
}

// Panic : log some text as panic and exit program
func Panic(text string) {
	log(levelErr, text)
	os.Exit(0)
}

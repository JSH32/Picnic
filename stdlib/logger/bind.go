package logger

type JSLogger struct {
}

// Info : JS bindings for info log
func (JSLogger) Info(content string) {
	Info(content)
}

// Error : JS bindings for error log
func (JSLogger) Error(content string) {
	Error(content)
}

// Panic : JS bindings for panic log
func (JSLogger) Panic(content string) {
	Panic(content)
}

// Success : JS bindings for success log
func (JSLogger) Success(content string) {
	Success(content)
}

// Warn : JS bindings for warning log
func (JSLogger) Warn(content string) {
	Warn(content)
}

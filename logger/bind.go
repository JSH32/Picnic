package logger

// NewJSLogger : JS bindings for logger
func NewJSLogger() *struct {
	Info    func(string) `json:"info"`
	Error   func(string) `json:"error"`
	Warn    func(string) `json:"warn"`
	Success func(string) `json:"success"`
	Panic   func(string) `json:"panic"`
} {
	return &struct {
		Info    func(string) `json:"info"`
		Error   func(string) `json:"error"`
		Warn    func(string) `json:"warn"`
		Success func(string) `json:"success"`
		Panic   func(string) `json:"panic"`
	}{
		Info:    Info,
		Error:   Error,
		Warn:    Warn,
		Success: Success,
		Panic:   Panic,
	}
}

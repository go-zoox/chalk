package chalk

var (
	//
	Success = func(message string) string {
		return Green.Color("✔ " + message)
	}
	Info = func(message string) string {
		return Blue.Color("ℹ " + message)
	}
	Warn = func(message string) string {
		return Yellow.Color("⚠ " + message)
	}
	Error = func(message string) string {
		return Red.Color("✖ " + message)
	}
)

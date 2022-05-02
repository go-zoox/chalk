package chalk

var (
	// Success is the success color
	Success = func(message string) string {
		return Green("✔ " + message)
	}
	// Info is the info color
	Info = func(message string) string {
		return Blue("ℹ " + message)
	}
	// Warn is the warn color
	Warn = func(message string) string {
		return Yellow("⚠ " + message)
	}
	// Error is the error color
	Error = func(message string) string {
		return Red("✖ " + message)
	}
	// Debug is the debug color
	Debug = func(message string) string {
		return Gray("ℹ " + message)
	}
)

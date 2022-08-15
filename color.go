package chalk

import "fmt"

// Color represents one of the ANSI color escape codes.
// http://en.wikipedia.org/wiki/ANSI_escape_code#Colors
type Color struct {
	value int
	reset int
}

// New creates a new color
func New(value int, reset int) func(string) string {
	color := Color{value, reset}
	return func(val string) string {
		return color.Color(val)
	}
}

// Value returns the individual value for this color
func (c Color) Value() int {
	return c.value
}

// Color colors the foreground of the given string
func (c Color) Color(val string) string {
	return fmt.Sprintf("%s%s%s", c, val, Color{c.reset, -1})
}

// String returns the string representation of this color
func (c Color) String() string {
	return fmt.Sprintf("\u001b[%dm", c.value)
}

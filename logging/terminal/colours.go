package terminal

import "strings"

// RED, GREEN, YELLOW color codes
const (
	RED     = "\x1b[31;1m"
	GREEN   = "\x1b[32;1m"
	YELLOW  = "\x1b[33;1m"
	BLUE    = "\x1b[34;1m"
	MAGENTA = "\x1b[35;1m"
	CYAN    = "\x1b[36;1m"
	WHITE   = "\x1b[37;1m"
	CLEAR   = "\x1b[0m"
)

// Colorize will add color information to a string to print it out in color
func Colorize(str string, color string) string {
	coloredString := []string{color, str, CLEAR}
	return strings.Join(coloredString, "")
}

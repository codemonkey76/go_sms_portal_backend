package ui

import "fmt"

const (
	ColorReset = "\u001b[0m" // Reset to default color

	ColorBlack   = "\u001b[30m"
	ColorRed     = "\u001b[31m"
	ColorGreen   = "\u001b[32m"
	ColorYellow  = "\u001b[33m"
	ColorBlue    = "\u001b[34m"
	ColorMagenta = "\u001b[35m"
	ColorCyan    = "\u001b[36m"
	ColorWhite   = "\u001b[37m"

	// Bright versions
	ColorBrightBlack   = "\u001b[30;1m"
	ColorBrightRed     = "\u001b[31;1m"
	ColorBrightGreen   = "\u001b[32;1m"
	ColorBrightYellow  = "\u001b[33;1m"
	ColorBrightBlue    = "\u001b[34;1m"
	ColorBrightMagenta = "\u001b[35;1m"
	ColorBrightCyan    = "\u001b[36;1m"
	ColorBrightWhite   = "\u001b[37;1m"
)

type Color string

func Colorize(color Color, message string) string {
	return fmt.Sprint(string(color), message, string(ColorReset))
}

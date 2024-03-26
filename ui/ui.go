package ui

import "fmt"

const (
	ColorYellow = "\u001b[33m"
	ColorReset  = "\u001b[0m"
	ColorGreen  = "\u001b[32m"
)

type Color string

func Colorize(color Color, message string) {
	fmt.Print(string(color), message, string(ColorReset))
}

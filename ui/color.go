package ui

import "strings"

// Color constants
const (
	Black   = "black"
	Red     = "red"
	Green   = "green"
	Yellow  = "yellow"
	Blue    = "blue"
	Magenta = "magenta"
	Cyan    = "cyan"
	White   = "white"
	Gray    = "gray"
	Default = "default"
)

// Option constants
const (
	Bold      = "bold"
	Underline = "underscore"
	Blink     = "blink"
	Reverse   = "reverse"
	Conceal   = "conceal"
)

// ANSI escape codes
const (
	EscapeStart = "\033["
	EscapeEnd   = "m"
	Reset       = "0"
)

// Color codes map
var colorCodes = map[string]string{
	Black:   "0",
	Red:     "1",
	Green:   "2",
	Yellow:  "3",
	Blue:    "4",
	Magenta: "5",
	Cyan:    "6",
	Gray:    "7",
	White:   "8",
	Default: "9",
}

// Option codes map
var optionCodes = map[string]string{
	Bold:      "1",
	Underline: "4",
	Blink:     "5",
	Reverse:   "7",
	Conceal:   "8",
}

// Color represents a text style with foreground, background and options.
type Color struct {
	Foreground string
	Background string
	Options    []string
}

func NewColor(foreground, background string, options ...string) *Color {
	return &Color{
		Foreground: foreground,
		Background: background,
		Options:    options,
	}
}

func (c *Color) Apply(text string) string {
	return c.Set() + text + c.Unset()
}

func (c *Color) Set() string {
	codes := []string{}

	for _, option := range c.Options {
		if code, ok := optionCodes[option]; ok {
			codes = append(codes, code)
		}
	}

	if code, ok := colorCodes[c.Background]; ok {
		codes = append(codes, "4"+code)
	}

	if code, ok := colorCodes[c.Foreground]; ok {
		codes = append(codes, "3"+code)
	}

	return EscapeStart + strings.Join(codes, ";") + EscapeEnd
}

func (c *Color) Unset() string {
	return EscapeStart + Reset + EscapeEnd
}

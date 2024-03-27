package ui

import (
	"fmt"
	"regexp"
	"strings"
)

func Info(v ...any) {
	c := NewColor("white", "blue", "bold")
	message := fmt.Sprintln(v...)
	fmt.Printf(c.Apply(" Info  ") + " " + message + "\n")
}

func Warn(format string, v ...any) {
	c := NewColor("black", "yellow", "bold")
	message := fmt.Sprintln(v...)
	fmt.Printf(c.Apply(" Warn  ") + " " + message + "\n")
}

func Error(format string, v ...any) {
	c := NewColor("white", "red", "bold")
	message := fmt.Sprintln(v...)
	fmt.Printf(c.Apply(" Error ") + " " + message + "\n")
}

func Output(c *Color, message string) string {
	return c.Apply(message)
}
func ColorizeUri(method, route string) string {
	return ColorizeMethod(method) + " " + ColorizeRoute(route)
}
func ColorizeMethod(method string) string {

	color := "white"
	if method == "GET" {
		color = "blue"
	} else if method == "POST" || method == "PUT" || method == "PATCH" {
		color = "yellow"
	} else if method == "DELETE" {
		color = "red"
	}
	coloredMethod := Output(NewColor(color, "default"), method)
	escapeCodesLength := len(coloredMethod) - len(method)
	formattingWidth := 8 + escapeCodesLength

	return fmt.Sprintf("%-*s", formattingWidth, coloredMethod)
}

func ColorizeRoute(route string) string {
	parts := GetPartsByField(route)
	var result string
	for _, part := range parts {
		if part.IsField {
			result = result + Output(NewColor("yellow", "default"), part.Part)
		} else {
			result = result + part.Part
		}
	}
	return result
}

type StringPart struct {
	Part    string
	IsField bool
}

func GetPartsByField(input string) []StringPart {
	re := regexp.MustCompile(`{[^}]*}`)

	matches := re.FindAllStringIndex(input, -1)
	var parts []StringPart
	lastIndex := 0
	for _, match := range matches {
		start, end := match[0], match[1]

		if start > lastIndex {
			parts = append(parts, StringPart{Part: input[lastIndex:start], IsField: false})
		}

		parts = append(parts, StringPart{Part: input[start:end], IsField: true})

		lastIndex = end
	}

	if lastIndex < len(input) {
		parts = append(parts, StringPart{Part: input[lastIndex:], IsField: false})
	}

	return parts
}

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

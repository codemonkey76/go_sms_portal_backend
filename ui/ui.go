package ui

import (
	"fmt"
	"regexp"
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

package ui

import "fmt"

func Info(message string) {
	c := NewColor("white", "blue", "bold")
	fmt.Printf(c.Apply(" Info ") + " " + message + "\n")
}

func Warn(message string) {

}

func Error(message string) {

}

func Output(c *Color, message string) string {
	return c.Apply(message)
}

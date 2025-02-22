package greeting

import "fmt"

func Greetings(name string) string {
	var message string

	message = fmt.Sprintln("Happy Diwali " + name + "!")
	return message
}

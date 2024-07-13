package api

import "fmt"

func Greetings(name string) string {
	//return a greeting that embeds
	message := fmt.Sprintf("Hi, %v Welcome!", name)
	return message
}

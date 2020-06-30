package main

import "fmt"

func Greetings(s string) string {
	if len(s) == 0 {
		return "Hello Guys Good morning"
	} else {
		return fmt.Sprintf("Helloo %v\n", s)
	}
}

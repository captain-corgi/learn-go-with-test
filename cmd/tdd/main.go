package main

import (
	"fmt"
)

const (
	englishHelloPrefix = "Hello %s"
)

func main() {
	fmt.Println(Hello("World"))
}

//Hello return "Hello World" string
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf(englishHelloPrefix, name)
}

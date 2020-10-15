package main

import (
	"fmt"
)

const (
	englishHelloPrefix    = "Hello %s"
	vietnameseHelloPrefix = "Chao %s"
)

func main() {
	fmt.Println(Hello("World", "VI"))
}

//Hello return "Hello World" string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return fmt.Sprintf(prefix, name)
}

func greetingPrefix(language string) string {
	prefix := ""
	switch language {
	case "VI":
		prefix = vietnameseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

//Add return sum of two number
func Add(a, b int) (rs int) {
	rs = a + b
	return
}

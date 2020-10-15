package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("World"))
}

//Hello return "Hello World" string
func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

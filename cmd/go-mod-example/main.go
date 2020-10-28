package main

import (
	"fmt"

	"github.com/captain-corgi/go-mod-example/greet"
)

func main() {
	fmt.Println(greet.HelloHuman("Predator"))
	fmt.Println(greet.HelloPredator("Anh Le"))
}

package main

import (
	"fmt"

	greetname "github.com/captain-corgi/go-mod-example/greet"
	"github.com/quii/go-mod-example/greet"
)

func main() {
	fmt.Println(greetname.HelloHuman("Anh Le"))
	fmt.Println(greet.TheWorld())
}

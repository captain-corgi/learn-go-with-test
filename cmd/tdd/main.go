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

//Repeat return a string that repeated `n` times
func Repeat(input string, repeatCount int) (rs string) {
	for i := 0; i < repeatCount; i++ {
		rs += input
	}
	return
}

//Sum return sum of numbers array
func Sum(numbers []int) (rs int) {
	for _, number := range numbers {
		rs += number
	}
	return
}

//SumAll return sum of all inputs array
func SumAll(numbersToSum ...[]int) (rs []int) {
	for _, numbers := range numbersToSum {
		rs = append(rs, Sum(numbers)) // NOTE: Create new slice every loop.
	}
	return rs
}

//SumAllTails return sum of all inputs array except first element
func SumAllTails(numbersToSum ...[]int) (rs []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) <= 0 {
			rs = append(rs, 0)
		} else {
			tail := numbers[1:]
			rs = append(rs, Sum(tail))
		}
	}
	return rs
}

//Rectangle describe a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

//Perimeter calculate Perimeter from given w and h
func Perimeter(rectangle Rectangle) (p float64) {
	if rectangle.Width <= 0 || rectangle.Height <= 0 {
		return 0.0
	}
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area calculate Area from given w and h
func Area(rectangle Rectangle) (a float64) {
	if rectangle.Width <= 0 || rectangle.Height <= 0 {
		return 0.0
	}
	return rectangle.Width * rectangle.Height
}

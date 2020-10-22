package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
)

const (
	englishHelloPrefix    = "Hello %s\n"
	vietnameseHelloPrefix = "Chao %s\n"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(MyGreetHandler))
}

//MyGreetHandler is a simple http handler
func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	GreetVN(w, "Anh")
}

//GreetVN say hello in Vietnamese
func GreetVN(writer io.Writer, name string) {
	str := Hello(name, "VI")
	fmt.Fprintf(writer, str)
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

type (
	//Rectangle describe a rectangle
	Rectangle struct {
		Width  float64
		Height float64
	}
	//Triangle describe a triangle
	Triangle struct {
		Base   float64
		Height float64
	}
	//Circle describe a circle
	Circle struct {
		Radius float64
	}
	//Shape represent for a shape.
	Shape interface {
		Area() float64
	}
)

//Area return Width*Height
func (r *Rectangle) Area() float64 {
	if r.Width <= 0 || r.Height <= 0 {
		return 0.0
	}
	return r.Width * r.Height
}

//Area return Width*Height
func (r *Triangle) Area() float64 {
	if r.Base <= 0 || r.Height <= 0 {
		return 0.0
	}
	return (r.Base * r.Height) / 2
}

//Area return Pi*Radius^2
func (r *Circle) Area() float64 {
	if r.Radius <= 0 {
		return 0.0
	}
	return r.Radius * r.Radius * math.Pi
}

//Perimeter calculate Perimeter from given w and h
func Perimeter(rectangle Rectangle) (p float64) {
	if rectangle.Width <= 0 || rectangle.Height <= 0 {
		return 0.0
	}
	return 2 * (rectangle.Width + rectangle.Height)
}

type (
	//Wallet describe a wallet with balance
	Wallet struct {
		balance Bitcoin
	}
	//Bitcoin currency, data type is int
	Bitcoin int
)

//Stringer override .String() method
type Stringer interface {
	String() string
}

var (
	//ErrInsufficientFunds is error when amount not enough for withdraw
	ErrInsufficientFunds = errors.New("cannot withdraw. Insufficient funds")
)

//Deposit increase amount in wallet
func (r *Wallet) Deposit(amount Bitcoin) {
	r.balance += amount
}

//Widthdraw decrease amount in wallet
func (r *Wallet) Widthdraw(amount Bitcoin) error {
	if r.balance < amount {
		return ErrInsufficientFunds
	}
	r.balance -= amount
	return nil
}

//Balance check amount in wallet
func (r *Wallet) Balance() Bitcoin {
	return r.balance
}

func (r Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", r)
}

type (
	//Dictionary is a map with key string, value string
	Dictionary map[string]string
)

var (
	//ErrWordNotFound is error when a word not found in dictionary
	ErrWordNotFound = fmt.Errorf("word not found")
	//ErrKeyWordEmpty is error when key word is empty
	ErrKeyWordEmpty = fmt.Errorf("word is empty")
	//ErrKeyWordDuplicate is error when key word duplicated
	ErrKeyWordDuplicate = fmt.Errorf("word is duplicated")
	//ErrKeyWordNotExist is error when key word not found in dictionary for update
	ErrKeyWordNotExist = fmt.Errorf("word not exist")
)

//Search find and return a word in a map if exist.
func (r Dictionary) Search(word string) (string, error) {
	if word == "" {
		return "", ErrKeyWordEmpty
	}
	definition, ok := r[word]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

//Add function append a word to dictionary
func (r Dictionary) Add(word string, definition string) error {
	_, err := r.Search(word)
	switch err {
	case ErrWordNotFound:
		r[word] = definition
	case nil:
		return ErrKeyWordDuplicate
	default:
		return err
	}
	return nil
}

//Update function aupdate an existing word in dictionary
func (r Dictionary) Update(word string, definition string) error {
	_, err := r.Search(word)
	switch err {
	case ErrWordNotFound:
		return ErrKeyWordNotExist
	case nil:
		r[word] = definition
	default:
		return err
	}
	return nil
}

//Delete remove a word from dictionary
func (r Dictionary) Delete(word string) {
	delete(r, word)
}

//RomanNumeral describe a roman numeral with key-value
type (
	RomanNumeral struct {
		Value  int
		Symbol string
	}
	RomanNumeralList []RomanNumeral
)

//ValueOf return value of a roman numeral
func (r RomanNumeralList) ValueOf(symbol string) int {
	for _, val := range r {
		if val.Symbol == symbol {
			return val.Value
		}
	}
	return 0
}

var (
	allRomanNumerals = []RomanNumeral{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
)

//RomanNumerals return a roman numeral from a number
func RomanNumerals(num int) string {
	var rs strings.Builder
	for _, n := range allRomanNumerals {
		for num >= n.Value {
			rs.WriteString(n.Symbol)
			num -= n.Value
		}
	}
	return rs.String()
}

//ConvertingToArabic return a number from a roman numeral
func ConvertingToArabic(roman string) int {
	total := 0
	for range roman {
		total++
	}
	return total
}

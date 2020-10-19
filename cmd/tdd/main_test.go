package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "main function can run",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("main function failed to run")
				}
			}()
			main()
		})
	}
}

func TestHello(t *testing.T) {
	type args struct {
		name     string
		language string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1. Function should return `Hello Anh Le`",
			args: args{
				name: "Anh Le",
			},
			want: "Hello Anh Le",
		},
		{
			name: "2. Param is empty, should return `Hello World`",
			args: args{
				name: "",
			},
			want: "Hello World",
		},
		{
			name: "3. Hello in English",
			args: args{
				name: "Anh",
			},
			want: "Hello Anh",
		},
		{
			name: "4. Hello in Vietnamese",
			args: args{
				name:     "Anh",
				language: "VI",
			},
			want: "Chao Anh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name, tt.args.language); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name   string
		args   args
		wantRs int
	}{
		{
			name: "1. 1+1=2",
			args: args{
				a: 1,
				b: 1,
			},
			wantRs: 2,
		},
		{
			name: "2. 2+2=4",
			args: args{
				a: 2,
				b: 2,
			},
			wantRs: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRs := Add(tt.args.a, tt.args.b); gotRs != tt.wantRs {
				t.Errorf("Add() = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	type args struct {
		input       string
		repeatCount int
	}
	tests := []struct {
		name   string
		args   args
		wantRs string
	}{
		{
			name:   "1. `a` repeated 5 times",
			args:   args{"a", 5},
			wantRs: "aaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRs := Repeat(tt.args.input, tt.args.repeatCount); gotRs != tt.wantRs {
				t.Errorf("Repeat() = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	rs := Repeat("a", 5)
	fmt.Println(rs)
	// Output: aaaaa
}

func TestSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name   string
		args   args
		wantRs int
	}{
		{
			name: "1. Empty input",
			args: args{
				numbers: []int{},
			},
			wantRs: 0,
		},
		{
			name: "2. Nil input",
			args: args{
				numbers: nil,
			},
			wantRs: 0,
		},
		{
			name: "3. Array has 1 element",
			args: args{
				numbers: []int{1},
			},
			wantRs: 1,
		},
		{
			name: "4. Array has more than 1 element",
			args: args{
				numbers: []int{1, 2, 3, 4},
			},
			wantRs: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRs := Sum(tt.args.numbers); gotRs != tt.wantRs {
				t.Errorf("Sum() = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}

func TestSumAllTails(t *testing.T) {
	type args struct {
		numbersToSum [][]int
	}
	tests := []struct {
		name   string
		args   args
		wantRs []int
	}{
		{
			name: "1. Empty input",
			args: args{
				numbersToSum: [][]int{},
			},
			wantRs: nil,
		},
		{
			name: "2. Nil input",
			args: args{
				numbersToSum: nil,
			},
			wantRs: nil,
		},
		{
			name: "3. Empty element",
			args: args{
				numbersToSum: [][]int{{}, {}},
			},
			wantRs: []int{0, 0},
		},
		{
			name: "4. Nil element",
			args: args{
				numbersToSum: [][]int{nil, nil},
			},
			wantRs: []int{0, 0},
		},
		{
			name: "5. Normal input",
			args: args{
				numbersToSum: [][]int{{1, 2, 3, 4, 5}, {6, 7}},
			},
			wantRs: []int{14, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRs := SumAllTails(tt.args.numbersToSum...); !reflect.DeepEqual(gotRs, tt.wantRs) {
				t.Errorf("SumAllTails() = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	type args struct {
		numbersToSum [][]int
	}
	tests := []struct {
		name   string
		args   args
		wantRs []int
	}{
		{
			name: "1. Empty input",
			args: args{
				numbersToSum: [][]int{},
			},
			wantRs: nil,
		},
		{
			name: "2. Nil input",
			args: args{
				numbersToSum: nil,
			},
			wantRs: nil,
		},
		{
			name: "3. Empty element",
			args: args{
				numbersToSum: [][]int{{}, {}},
			},
			wantRs: []int{0, 0},
		},
		{
			name: "4. Nil element",
			args: args{
				numbersToSum: [][]int{nil, nil},
			},
			wantRs: []int{0, 0},
		},
		{
			name: "5. Normal input",
			args: args{
				numbersToSum: [][]int{{1, 2, 3, 4, 5}, {6, 7}},
			},
			wantRs: []int{15, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRs := SumAll(tt.args.numbersToSum...); !reflect.DeepEqual(gotRs, tt.wantRs) {
				t.Errorf("SumAll() = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	type args struct {
		rectangle Rectangle
	}
	tests := []struct {
		name  string
		args  args
		wantP float64
	}{
		{
			name: "1. Input is 0.0x0.0",
			args: args{
				Rectangle{0, 0},
			},
			wantP: 0.0,
		},
		{
			name: "2. Input is possitive",
			args: args{
				Rectangle{10.0, 10.0},
			},
			wantP: 40.0,
		},
		{
			name: "3. Input is negative",
			args: args{
				Rectangle{-10.0, -10.0},
			},
			wantP: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := Perimeter(tt.args.rectangle); gotP != tt.wantP {
				t.Errorf("Perimeter() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func TestArea(t *testing.T) {
	type args struct {
		rectangle Rectangle
	}
	tests := []struct {
		name  string
		args  args
		wantA float64
	}{
		{
			name: "1. Input is 0x0",
			args: args{
				Rectangle{0.0, 0.0},
			},
			wantA: 0.0,
		},
		{
			name: "2. Input is possitive",
			args: args{
				Rectangle{10.0, 10.0},
			},
			wantA: 100.0,
		},
		{
			name: "3. Input is negative",
			args: args{
				Rectangle{-10.0, 10.0},
			},
			wantA: 0,
		},
		{
			name: "4. Input is negative",
			args: args{
				Rectangle{10.0, -10.0},
			},
			wantA: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := Area(tt.args.rectangle); gotA != tt.wantA {
				t.Errorf("Area() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

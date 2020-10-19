package main

import (
	"bytes"
	"errors"
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
		shape Shape
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1. Rectangle Input is 0x0",
			args: args{
				&Rectangle{0.0, 0.0},
			},
			want: 0.0,
		},
		{
			name: "2. Rectangle Input is possitive",
			args: args{
				&Rectangle{10.0, 10.0},
			},
			want: 100.0,
		},
		{
			name: "3. Rectangle Input is negative",
			args: args{
				&Rectangle{-10.0, 10.0},
			},
			want: 0,
		},
		{
			name: "4. Rectangle Input is negative",
			args: args{
				&Rectangle{10.0, -10.0},
			},
			want: 0,
		},
		{
			name: "5. Circle Input is 0",
			args: args{
				&Circle{0.0},
			},
			want: 0.0,
		},
		{
			name: "6. Circle Input is possitive",
			args: args{
				&Circle{10.0},
			},
			want: 314.1592653589793,
		},
		{
			name: "7. Circle Input is negative",
			args: args{
				&Circle{-10.0},
			},
			want: 0,
		},
		{
			name: "8. Triangle Input is 0",
			args: args{
				&Triangle{0.0, 0.0},
			},
			want: 0.0,
		},
		{
			name: "9. Triangle Input is possitive",
			args: args{
				&Triangle{10.0, 10.0},
			},
			want: 50.0,
		},
		{
			name: "7. Triangle Input is negative",
			args: args{
				&Triangle{-10.0, 10.0},
			},
			want: 0,
		},
		{
			name: "8. Triangle Input is negative",
			args: args{
				&Triangle{10.0, -10.0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.shape.Area(); got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt.args.shape, got, tt.want)
			}
		})
	}
}

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		gotErr := wallet.Widthdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, gotErr)
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingAmt := Bitcoin(20)
		wallet := Wallet{startingAmt}

		wantAmt := Bitcoin(20)
		wantErr := ErrInsufficientFunds

		gotErr := wallet.Widthdraw(Bitcoin(200))

		assertBalance(t, wallet, wantAmt)
		assertError(t, gotErr, wantErr)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t *testing.T, gotErr error, wantErr error) {
	t.Helper()
	if gotErr == nil {
		t.Errorf("expected error, but got nil")
	}
	if gotErr.Error() != wantErr.Error() {
		t.Errorf("got %s, want %s", gotErr.Error(), wantErr.Error())
	}
}

func assertNoError(t *testing.T, gotErr error) {
	t.Helper()
	if gotErr != nil {
		t.Errorf("expected no error, but got %s", gotErr.Error())
	}
}

func TestBitcoin_String(t *testing.T) {
	tests := []struct {
		name string
		r    Bitcoin
		want string
	}{
		{
			name: "Zero",
			r:    Bitcoin(0),
			want: "0 BTC",
		},
		{
			name: "Positive",
			r:    Bitcoin(10),
			want: "10 BTC",
		},
		{
			name: "Negative",
			r:    Bitcoin(-10),
			want: "-10 BTC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("Bitcoin.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	type args struct {
		dictionary Dictionary
		word       string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "1. TDD TC",
			args: args{
				dictionary: Dictionary{},
			},
			want: "",
		},
		{
			name: "2. Word not found",
			args: args{
				dictionary: Dictionary{
					"a": "this is a",
				},
				word: "b",
			},
			want:    "",
			wantErr: errors.New("word not found"),
		},
		{
			name: "3. Word found",
			args: args{
				dictionary: Dictionary{
					"a": "this is a",
				},
				word: "a",
			},
			want: "this is a",
		},
		{
			name: "4. Word empty",
			args: args{
				dictionary: Dictionary{
					"a": "this is a",
				},
				word: "",
			},
			want:    "",
			wantErr: ErrKeyWordEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := tt.args.dictionary.Search(tt.args.word)
			if got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
			if tt.wantErr != nil {
				if gotErr == nil {
					t.Errorf("Search() got err %v, want err %v", "nil", tt.wantErr.Error())
				} else {
					if gotErr.Error() != tt.wantErr.Error() {
						t.Errorf("Search() got err %v, want err %v", gotErr.Error(), tt.wantErr.Error())
					}
				}
			}
		})
	}
}

func TestDictionary_Add(t *testing.T) {
	type args struct {
		word       string
		definition string
	}
	tests := []struct {
		name    string
		r       *Dictionary
		args    args
		wantErr error
	}{
		{
			name: "1. TDD",
			args: args{
				word:       "a",
				definition: "this is a",
			},
			r: &Dictionary{},
		},
		{
			name: "2. Input word is empty",
			args: args{
				word:       "",
				definition: "this is a",
			},
			r:       &Dictionary{},
			wantErr: ErrKeyWordEmpty,
		},
		{
			name: "3. Input word already existed",
			args: args{
				word:       "a",
				definition: "this is a",
			},
			r: &Dictionary{
				"a": "this is a",
			},
			wantErr: ErrKeyWordDuplicate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.r.Add(tt.args.word, tt.args.definition)
			if tt.wantErr != nil {
				if gotErr == nil {
					t.Errorf("Add() got err %v, want err %v", "nil", tt.wantErr.Error())
				} else {
					if gotErr.Error() != tt.wantErr.Error() {
						t.Errorf("Add() got err %v, want err %v", gotErr.Error(), tt.wantErr.Error())
					}
				}
			} else {
				def, searchErr := tt.r.Search(tt.args.word)
				if def != tt.args.definition || searchErr != nil {
					t.Errorf("Incorrect added definition. Got %s, want %s", def, tt.args.definition)
				}
			}
		})
	}
}

func TestDictionary_Update(t *testing.T) {
	type args struct {
		word       string
		definition string
	}
	tests := []struct {
		name    string
		r       Dictionary
		args    args
		wantErr error
	}{
		{
			name: "1. TDD",
			args: args{
				word:       "",
				definition: "",
			},
			r:       Dictionary{},
			wantErr: ErrKeyWordEmpty,
		},
		{
			name: "2. Word not exist",
			args: args{
				word:       "b",
				definition: "this is b",
			},
			r: Dictionary{
				"a": "this is a",
			},
			wantErr: ErrKeyWordNotExist,
		},
		{
			name: "3. Word exist",
			args: args{
				word:       "a",
				definition: "this is a",
			},
			r: Dictionary{
				"a": "this is a",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Update(tt.args.word, tt.args.definition); err != nil && err != tt.wantErr {
				t.Errorf("Dictionary.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDictionary_Delete(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		r       Dictionary
		args    args
		wantErr error
	}{
		{
			name: "1. TDD",
			args: args{
				word: "",
			},
			r:       Dictionary{},
			wantErr: nil,
		},
		{
			name: "2. Word not exist",
			args: args{
				word: "b",
			},
			r: Dictionary{
				"a": "this is a",
			},
			wantErr: ErrKeyWordNotExist,
		},
		{
			name: "3. Word exist",
			args: args{
				word: "a",
			},
			r: Dictionary{
				"a": "this is a",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Delete(tt.args.word)
			gotDef, _ := tt.r.Search(tt.args.word)
			if gotDef != "" {
				t.Errorf("word still exist, not deleted yet")
			}
		})
	}
}

func TestGreetVN(t *testing.T) {
	buffer := bytes.Buffer{}
	GreetVN(&buffer, "Anh")

	got := buffer.String()
	want := "Chao Anh"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

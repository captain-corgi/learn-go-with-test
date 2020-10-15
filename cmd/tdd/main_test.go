package main

import (
	"testing"
)

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRs := Add(tt.args.a, tt.args.b); gotRs != tt.wantRs {
				t.Errorf("Add() = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}

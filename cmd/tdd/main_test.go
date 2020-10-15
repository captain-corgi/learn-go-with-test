package main

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name string
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

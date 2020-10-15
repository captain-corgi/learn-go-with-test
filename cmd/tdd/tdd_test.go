package main

import (
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := ""
	want := ""

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

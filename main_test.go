package main

import (
	"fmt"
	"testing"
)

func TestMainFunc(t *testing.T) {
	gotGreeting := GreetWorld()
	expectedGreeting := "Hello, world!!"
	if gotGreeting != expectedGreeting {
		t.Fatalf("Expected %s, but got: %s", expectedGreeting, gotGreeting)
	}
}

const a = 3

func TestTypeOfConst(t *testing.T) {
	fmt.Printf("type is: %T for %v\n", a, a)
}

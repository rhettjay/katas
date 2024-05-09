package FizzBuzz

import (
	"testing"
)

func Test_Zero(t *testing.T) {
	num, _ := answer(0)
	if num != 0 {
		t.Error("answer was not 2 as expected")
	}
}

func Test_Two(t *testing.T) {
	num, _ := answer(2)
	if num != 2 {
		t.Error("answer was not 2 as expected")
	}
}

func Test_Three(t *testing.T) {
	_, msg := answer(3)
	if msg != "Fizz" {
		t.Error("answer was not Fizz as expected")
	}
}

func Test_Five(t *testing.T) {
	_, msg := answer(5)
	if msg != "Buzz" {
		t.Error("answer was not Buzz as multiple of five")
	}
}

func Test_Three_Five(t *testing.T) {
	_, msg := answer(15)
	if msg != "FizzBuzz" {
		t.Error("answer was not FizzBuzz as multiple of five")
	}
}

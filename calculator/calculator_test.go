package calculator_test

import (
	"calculator/calculator"
	"testing"
)

func TestCalculatorSize(t *testing.T) {
	var calc calculator.Calculator
	got := calc.Size()
	want := 0
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculatorPush(t *testing.T) {
	var calc calculator.Calculator
	calc.Push(4)
	got := calc.Size()
	want := 1
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculatorPeek(t *testing.T) {
	var calc calculator.Calculator
	calc.Push(4)
	got, _ := calc.Peek()
	var want float64 = 4
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

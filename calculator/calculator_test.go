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
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCalculatorPush(t *testing.T) {
	var calc calculator.Calculator
	calc.Push(4)
	got := calc.Size()
	want := 1
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCalculatorPeek(t *testing.T) {
	var calc calculator.Calculator
	calc.Push(4)
	got, _ := calc.Peek()
	var want float64 = 4
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestCalculatorPopZeroElements(t *testing.T) {
	var calc calculator.Calculator
	_, got := calc.Pop()

	want := false
	if got != want {
		t.Errorf("want %t, got %t", want, got)
	}
}

func TestCalculatorPop(t *testing.T) {
	var calc calculator.Calculator

	calc.Push(1)
	calc.Push(2)
	calc.Push(3)

	var want float64 = 3
	got, ok := calc.Pop()

	if !ok {
		t.Errorf("expected %t.", true)
	}
	if got != want {
		t.Errorf("want %f, got %f.", want, got)
	}

	size := calc.Size()
	expected := 2
	if size != expected {
		t.Errorf("expected %d, got %d.", size, expected)
	}

	want = 2
	got, ok = calc.Pop()

	if !ok {
		t.Errorf("expected %t.", true)
	}
	if got != want {
		t.Errorf("want %f, got %f.", want, got)
	}

	size = calc.Size()
	expected = 1
	if size != expected {
		t.Errorf("expected %d, got %d.", size, expected)
	}
}

func TestCalculatorAddZeroElements(t *testing.T) {
	var calc calculator.Calculator
	want := float64(0)
	want_err := false
	got, got_err := calc.Add()

	if got != want {
		t.Errorf("want %f, got %f.", want, got)
	}
	if got_err != want_err {
		t.Errorf("want %t, got %t.", want_err, got_err)
	}
}

func TestCalculatorAddOneElement(t *testing.T) {
	var calc calculator.Calculator
	want := float64(0)
	want_err := false
	calc.Push(1)
	got, got_err := calc.Add()
	want_size := 1
	got_size := calc.Size()

	if got != want {
		t.Errorf("want %f, got %f.", want, got)
	}
	if got_err != want_err {
		t.Errorf("want %t, got %t.", want_err, got_err)
	}
	if float64(got_size) != float64(want_size) {
		t.Errorf("want %d, got %d.", want_size, got_size)
	}
}

func TestCalculatorAdd(t *testing.T) {
	var calc calculator.Calculator
	calc.Push(2)
	calc.Push(2)
	want := float64(4)
	got, _ := calc.Add()

	if want != got {
		t.Errorf("want %f, got %f.", want, got)
	}
}

func TestCalculatorSubtractZeroElements(t *testing.T) {
	var calc calculator.Calculator
	_, got := calc.Subtract()

	want := false
	if got != want {
		t.Errorf("want %t, got %t", want, got)
	}
}

func TestCalculatorSubtractUnary(t *testing.T) {
	var calc calculator.Calculator
	calc.Push(2)
	want := float64(-2)
	got, _ := calc.Subtract()

	if want != got {
		t.Errorf("want %f, got %f.", want, got)
	}
}

func TestCalculatorSubtract(t *testing.T) {
	var calc calculator.Calculator
	calc.Push(2)
	calc.Push(3)
	want := float64(1)
	got, _ := calc.Subtract()

	if want != got {
		t.Errorf("want %f, got %f.", want, got)
	}
}

package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Multiply(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestUnaryNegativeExpectsNegative(t *testing.T) {
	t.Parallel()
	var want float64 = -1
	got := calculator.UnaryNegative(1)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestUnaryNegativeExpectsPositive(t *testing.T) {
	t.Parallel()
	var want float64 = 1
	got := calculator.UnaryNegative(-1)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestUnaryNegativeExpectsZero(t *testing.T) {
	t.Parallel()
	var want float64 = 0
	got := calculator.UnaryNegative(0)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivideSame(t *testing.T) {
	t.Parallel()
	var want float64 = 1
	got := calculator.Divide(4, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivideLesserThan(t *testing.T) {
	t.Parallel()
	var want float64 = 0.5
	got := calculator.Divide(1, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivideGreaterThan(t *testing.T) {
	t.Parallel()
	var want float64 = 3
	got := calculator.Divide(9, 3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivideByZero(t *testing.T) {
	t.Parallel()
	var want float64 = math.Inf(0)
	got := calculator.Divide(1, 0)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSquareRootInfinite(t *testing.T) {
	t.Parallel()
	var want float64 = math.Inf(0)
	got := calculator.SquareRoot(math.Inf(0))
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSquareRootZero(t *testing.T) {
	t.Parallel()
	var want float64 = 0
	got := calculator.SquareRoot(0)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSquareRootOfFour(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.SquareRoot(4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSquareRootNegative(t *testing.T) {
	t.Parallel()
	got := calculator.SquareRoot(-1)
	if !math.IsNaN(got) {
		t.Errorf("want %f, got %f", math.NaN(), got)
	}
}

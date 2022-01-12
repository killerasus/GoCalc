package operations_test

import (
	"calculator/operations"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	var want float64 = 4
	got := operations.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	var want float64 = 2
	got := operations.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	var want float64 = 4
	got := operations.Multiply(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestUnaryNegativeExpectsNegative(t *testing.T) {
	var want float64 = -1
	got := operations.UnaryNegative(1)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestUnaryNegativeExpectsPositive(t *testing.T) {
	var want float64 = 1
	got := operations.UnaryNegative(-1)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestUnaryNegativeExpectsZero(t *testing.T) {
	var want float64 = 0
	got := operations.UnaryNegative(0)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivideSame(t *testing.T) {
	var want float64 = 1
	got := operations.Divide(4, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivideLesserThan(t *testing.T) {
	var want float64 = 0.5
	got := operations.Divide(1, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivideGreaterThan(t *testing.T) {
	var want float64 = 3
	got := operations.Divide(9, 3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivideByZero(t *testing.T) {
	var want float64 = math.Inf(0)
	got := operations.Divide(1, 0)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSquareRootInfinite(t *testing.T) {
	var want float64 = math.Inf(0)
	got := operations.SquareRoot(math.Inf(0))
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSquareRootZero(t *testing.T) {
	var want float64 = 0
	got := operations.SquareRoot(0)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSquareRootOfFour(t *testing.T) {
	var want float64 = 2
	got := operations.SquareRoot(4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSquareRootNegative(t *testing.T) {
	got := operations.SquareRoot(-1)
	if !math.IsNaN(got) {
		t.Errorf("want %f, got %f", math.NaN(), got)
	}
}

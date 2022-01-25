package calculator

import (
	"calculator/operations"
	"calculator/stack"
	"math"
)

const (
	Sum = iota
	Subtraction
	Multiplication
	Division
	Raditiation
	Mod
	Negative
)

type Calculator struct {
	stack stack.Stack
}

func (c *Calculator) Size() int {
	return c.stack.Size()
}

func (c *Calculator) Push(a float64) {
	c.stack.Push(a)
}

func (c *Calculator) Peek() (float64, bool) {
	return c.stack.Peek()
}

func (c *Calculator) Peeki(i int) (float64, bool) {
	return c.stack.Peeki(i)
}

func (c *Calculator) Pop() (float64, bool) {
	return c.stack.Pop()
}

func (c *Calculator) Add() (float64, bool) {
	if c.Size() < 2 {
		return 0, false
	} else {
		a, _ := c.Pop()
		b, _ := c.Pop()

		resp := operations.Add(a, b)
		c.Push(resp)
		return resp, true
	}
}

func (c *Calculator) Subtract() (float64, bool) {
	if c.Size() < 1 {
		return 0, false
	} else if c.Size() == 1 {
		a, _ := c.Pop()
		resp := operations.UnaryNegative(a)
		c.Push(resp)
		return resp, true
	} else {
		b, _ := c.Pop()
		a, _ := c.Pop()

		resp := operations.Subtract(a, b)
		c.Push(resp)
		return resp, true
	}
}

func (c *Calculator) Multiplication() (float64, bool) {
	if c.Size() < 2 {
		return 0, false
	}

	a, _ := c.Pop()
	b, _ := c.Pop()

	resp := operations.Multiply(a, b)
	c.Push(resp)
	return resp, true
}

func (c *Calculator) Division() (float64, bool) {
	if c.Size() < 2 {
		return 0, false
	}

	a, _ := c.Pop()
	b, _ := c.Pop()

	resp := operations.Divide(a, b)
	if math.IsInf(resp, 0) {
		return resp, false
	}

	c.Push(resp)
	return resp, true
}

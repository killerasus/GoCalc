// Implements a RPN calculator
// (RPN) a b op --> (Arithmetic) a op b
package calculator

import (
	"calculator/operations"
	"calculator/stack"
	"math"
	"strconv"
	"strings"
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

	b, _ := c.Pop()
	a, _ := c.Pop()

	resp := operations.Multiply(a, b)
	c.Push(resp)
	return resp, true
}

func (c *Calculator) Division() (float64, bool) {
	if c.Size() < 2 {
		return 0, false
	}

	divisor, _ := c.Pop()
	dividend, _ := c.Pop()

	resp := operations.Divide(dividend, divisor)
	if math.IsInf(resp, 0) {
		return resp, false
	}

	c.Push(resp)
	return resp, true
}

func (c *Calculator) Evaluate(e string) (v float64, ok bool) {
	exp := strings.Split(e, " ")
	for _, s := range exp {
		if n, err := strconv.ParseFloat(s, 64); err == nil {
			c.Push(n)
			v = n
			ok = true
		} else {
			switch s {
			case "+":
				if v, ok = c.Add(); !ok {
					return
				}
			case "-":
				if v, ok = c.Subtract(); !ok {
					return
				}
			case "*":
				if v, ok = c.Multiplication(); !ok {
					return
				}
			case "/":
				if v, ok = c.Division(); !ok {
					return
				}
			default: // Not recognized
				return 0, false
			}
		}
	}
	return
}

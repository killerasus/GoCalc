package calculator

import "calculator/stack"

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

func (c *Calculator) Pop() (float64, bool) {
	return c.stack.Pop()
}

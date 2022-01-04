package stack

const (
	Sum = iota
	Subtraction
	Multiplication
	Division
	Raditiation
	Mod
	Negative
)

type IStack interface {
	Push(float64)
	Pop() float64
	Peek() float64
	Size() int
	IsEmpty() bool
}

type Stack struct {
	stack []float64
}

func (s *Stack) Push(v float64) {
	s.stack = append(s.stack, v)
}

func (s *Stack) Peek() float64 {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) Pop() float64 {
	v := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return v
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

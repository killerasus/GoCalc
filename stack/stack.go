package stack

type IStack interface {
	Push(float64)
	Pop() float64
	Peek() float64
	Peeki(i int) (float64, bool)
	Size() int
	IsEmpty() bool
}

type Stack struct {
	stack []float64
}

func (s *Stack) Push(v float64) {
	s.stack = append(s.stack, v)
}

func (s *Stack) Peek() (float64, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}
	return s.stack[len(s.stack)-1], true
}

func (s *Stack) Peeki(i int) (v float64, ok bool) {
	if i >= len(s.stack) {
		return
	}
	v = s.stack[len(s.stack)-(i+1)]
	ok = true
	return
}

func (s *Stack) Pop() (float64, bool) {
	if len(s.stack) == 0 {
		return 0, false
	} else {
		v := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return v, true
	}
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

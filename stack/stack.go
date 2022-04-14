package stack

type Item interface{}

type Stack interface {
	Push(item Item)
	Pop() Item
}

// 底层用 slice 实现
type StackSlice struct {
	container []Item
}

func (s *StackSlice) Push(item Item) {
	s.container = append(s.container, item)
}

func (s *StackSlice) Pop() Item {
	if len(s.container) == 0 {
		return nil
	}

	item := s.container[len(s.container)-1]
	s.container = s.container[:len(s.container)-1]
	return item
}

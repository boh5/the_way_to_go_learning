package main

import (
	"fmt"
	"strings"
)

type StackArr struct {
	arr *[4]int
}

func main() {
	sa := &StackArr{new([4]int)}
	sa.Push(1)
	sa.Push(2)
	fmt.Println(sa)
	fmt.Println(sa.Pop())
	fmt.Println(sa)
	sa.Push(3)
	fmt.Println(sa)
}

func (s *StackArr) Push(a int) {
	pushed := false
	for i := 0; i < 4; i++ {
		if s.arr[i] == 0 {
			s.arr[i] = a
			pushed = true
			break
		}
	}
	if !pushed {
		for i := 0; i < 3; i++ {
			s.arr[i] = s.arr[i+1]
		}
		s.arr[len(s.arr)-1] = a
	}
}

func (s *StackArr) Pop() (ret int) {
	for i := len(s.arr) - 1; i >= 0; i-- {
		if s.arr[i] != 0 {
			ret = s.arr[i]
			s.arr[i] = 0
			return
		}
	}
	return 0
}

func (s *StackArr) String() (ret string) {
	for i, v := range s.arr {
		ret += "[" + fmt.Sprint(i) + ":" + fmt.Sprint(v) + "] "
	}
	ret = strings.Trim(ret, " ")
	return
}

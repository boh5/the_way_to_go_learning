package main

import (
	"fmt"
	"strings"
)

const LIMIT = 4

type StackStruct struct {
	curFree int
	data    *[LIMIT]int
}

func main() {
	ss := &StackStruct{0, new([LIMIT]int)}
	ss.Push(1)
	ss.Push(2)
	fmt.Println(ss)
	fmt.Println(ss.Pop())
	fmt.Println(ss)
	ss.Push(3)
	fmt.Println(ss)
}

func (s *StackStruct) Push(val int) {
	if s.curFree == LIMIT {
		for i := 0; i < LIMIT-1; i++ {
			s.data[i] = s.data[i+1]
		}
		s.data[LIMIT-1] = val
	} else {
		s.data[s.curFree] = val
		s.curFree += 1
	}
}

func (s *StackStruct) Pop() (ret int) {
	if s.curFree > 0 {
		ret = s.data[s.curFree-1]
		s.curFree -= 1
		s.data[s.curFree] = 0
	} else {
		ret = 0
	}
	return
}

func (s *StackStruct) String() (ret string) {
	for i, v := range s.data {
		ret += "[" + fmt.Sprint(i) + ":" + fmt.Sprint(v) + "] "
	}
	ret = strings.Trim(ret, " ")
	return
}

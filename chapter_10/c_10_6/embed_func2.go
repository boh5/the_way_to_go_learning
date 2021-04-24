package main

import "fmt"

type Log1 struct {
	msg string
}

type Customer1 struct {
	Name string
	Log1
}

func main() {
	c := &Customer1{"Barak Obama", Log1{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
}

func (l *Log1) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log1) String() string {
	return l.msg
}

func (c *Customer1) String() string {
	return c.Name + "\nLog:\n" + fmt.Sprintln(c.Log1.String())
}

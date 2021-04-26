package main

import (
	"fmt"
	"reflect"
)

var secret interface{} = NotKnownType{"Ada", "Go", "Oberon"}

type NotKnownType struct {
	s1, s2, s3 string
}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)

	fmt.Println(typ)
	knd := value.Kind()
	fmt.Println(knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}

	results := value.Method(0).Call(nil)
	fmt.Println(results)
}

func (n NotKnownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

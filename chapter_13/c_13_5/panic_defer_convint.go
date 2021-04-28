package main

import (
	"fmt"
	"math"
)

func ConvertInt64ToInt(i64 int64) (i32 int32) {
	if math.MinInt32 <= i64 && i64 <= math.MaxInt32 {
		return int32(i64)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", i64))
}

func IntFromInt64(i64 int64) (i32 int32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	i32 = ConvertInt64ToInt(i64)
	return i32, nil
}

func main() {
	i, e := IntFromInt64(12333333333333)
	fmt.Println(i)
	fmt.Println(e)
}

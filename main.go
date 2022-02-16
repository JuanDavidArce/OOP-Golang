package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("l", 0, 64)
	if err != nil {
		fmt.Println("%v", err)
	} else {
		fmt.Println(myValue)
	}

	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])

	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
}

package main

import (
	"fmt"
	"strconv"
	"time"
)

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}

func main() {

	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("l", 0, 64)

	if err != nil {
		fmt.Println(err)
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

	array := [12]int{1}
	fmt.Println(array)

	myInterface := []interface{}{}
	myInterface = append(myInterface, "hola")
	fmt.Println(myInterface...)

	// c := make(chan int)
	// go doSomething(c)
	// <-c

	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)

	var c []int
	d := []int{}
	c = append(c, 25)
	fmt.Print(c, d)

}

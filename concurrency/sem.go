package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Println("id", i, "started")
	time.Sleep(4 * time.Second)
	fmt.Println("id", i, "ended")
	<-c

}

func main() {
	c := make(chan int, 2)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait()

}

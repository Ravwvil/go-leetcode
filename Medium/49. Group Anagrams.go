package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func main() {
	tasks := make(chan int, 1)
	tasks <- 1
	tasks <- 2
	go plus(10, 3)
	go plus(10, 4)
	go plus(10, 5)
	smth := <-tasks
	fmt.Println(smth)
}

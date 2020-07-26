package main

import "fmt"

func producer(c chan int)  {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println("produce: ", i)
	}
	defer close(c)
}

func consumer(c chan int)  {
	have := true
	var p int
	for have {
		if p, have = <-c; have {
			fmt.Println("consumer: ", p)
		}
	}
}

func main()  {
	c := make(chan int)
	go producer(c)
	consumer(c)
}
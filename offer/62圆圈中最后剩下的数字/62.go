package main

import (
	"fmt"
)

func main()  {
	n, m := 10,17
	ans := lastRemaining(n,m)
	fmt.Println(ans)
}

func lastRemaining(n int, m int) int {
    if n==1{
        return 0
    }
    return (lastRemaining(n-1,m)+m)%n
}
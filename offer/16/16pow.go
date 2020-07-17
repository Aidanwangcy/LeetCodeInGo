package main

import "fmt"

func main()  {
	x := 2.00
	n := 3
	ans := myPow(x, n)
	fmt.Println(ans)
}

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	if n % 2 == 0 {
		square := myPow(x, n/2)
		return square * square
	} else {
		square := myPow(x, (n-1)/2)
		return square * square * x
	}
}
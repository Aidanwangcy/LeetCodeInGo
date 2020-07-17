package main

import "fmt"

func main() {
	var x int
	fmt.Print("输入数字：")
	fmt.Scan(&x)
	ans := isPalindrome(x)
	fmt.Printf("%v\n", ans)

}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNum := 0
	for x > revertedNum {
		revertedNum = revertedNum*10 + x%10
		x /= 10
	}

	return x == revertedNum || x == revertedNum/10
}

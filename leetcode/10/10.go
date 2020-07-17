package main

import "fmt"

func main()  {
	var s, p string
	fmt.Print("输入字符串：")
	fmt.Scan(&s)
	fmt.Print("输入字符规律：")
	fmt.Scan(&p)
	ans := isMatch(s, p)
	fmt.Printf("%v\n", ans)
}

func isMatch(s string, p string) bool {
	return false
}
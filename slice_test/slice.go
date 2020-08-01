package main

import "fmt"

func main() {
    a := [3]int{1, 1, 1}
    s := []int{1, 1, 1}
    
    modifyArray(a)
    fmt.Println(a)
    
    modifySlice(s)
    fmt.Println(s)
    
    reslice(s)
    fmt.Println(s)
    
    s1 := make([]int, 2, 3)
    s2 := append(s1, 1)
    s2[0] = 3
    fmt.Println(s1)

    s3 := append(s1, 1, 1)
    s3[0] = 4
    fmt.Println(s1)
}

func modifyArray(a [3]int) {
    a[0] = 2
}

func modifySlice(s []int) {
    s[0] = 2
}

func reslice(s []int) {
    s = s[:2]
}
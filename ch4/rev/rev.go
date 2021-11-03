package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	// 左移可以靠反转实现
	reverse(a[:2])
	reverse(a[2:])
	reverse(a[:])
	fmt.Println(a)
}

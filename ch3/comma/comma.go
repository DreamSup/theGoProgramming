package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma1(s string) string {
	var buf bytes.Buffer
	n := strings.LastIndex(s, ".")
	if n < 0 {
		n = len(s)
	}
	m := len(s)
	for i := 0; i < m; i++ {
		if (n-i)%3 == 0 && i > 0 && i < n {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func isomerism(s string, p string) bool {
	hashMap := make(map[byte]int)
	if len(s) != len(p) {
		return false
	}
	for i := range s {
		hashMap[s[i]]++
		hashMap[p[i]]--
	}
	for _, v := range hashMap {
		if v != 0 {
			return false
		}
	}
	return true.
}

func main() {
	fmt.Println(comma("1123456789"))
	fmt.Println(isomerism("刘宏政", "宏证六"))
}

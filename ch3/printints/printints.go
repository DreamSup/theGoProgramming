package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']') // 添加utf-8,最好用buf.WriteRune
	return buf.String()
}

func main() {
	fmt.Println([]int{1, 2, 3})
	fmt.Println(intsToString([]int{1, 2, 3}))
}

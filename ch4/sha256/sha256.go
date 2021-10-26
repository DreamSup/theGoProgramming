package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func differentBit(x [32]byte, y [32]byte) int {
	var pc [256]byte
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	ans := 0
	for i := range x {
		ans += int(pc[x[i]^y[i]])
	}
	return ans
}

var m = flag.String("m", " ", "hash function")

func getSha(s string) interface{} {
	if s == "sha384" {
		return sha512.Sum384([]byte("x"))
	} else if s == "sha512" {
		return sha512.Sum512([]byte("x"))
	} else {
		return sha256.Sum256([]byte("x"))
	}
}

func main() {
	flag.Parse()
	c1 := getSha(*m)
	fmt.Printf("%x\n%T\n", c1, c1)
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n", c2)
	if _, ok := c1.([32]byte); ok {
		fmt.Printf("%d\n", differentBit(c1.([32]byte), c2))
	}
	zero(&c2)
	fmt.Printf("%x\n", c2)
}

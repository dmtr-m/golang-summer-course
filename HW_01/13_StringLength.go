package main

import (
	"fmt"
	"unsafe"
)

func StringLen(s string) int {
	// Works only for UTF-8 and other 1-byte symbol encoding
	return int((*[2]uintptr)(unsafe.Pointer(&s))[1])
}

func TestStringLength() {
	var s string
	fmt.Scan(&s)
	fmt.Println(StringLen(s))
}

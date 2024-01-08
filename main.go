package main

// #cgo CFLAGS: -g -Wall -Isrc
// #cgo LDFLAGS: -Lbuild -lmytest
// #include <stdlib.h>
// #include "mytest.h"
import "C"
import (
	"fmt"
)

type MyTest struct {
	a uint
	b uint
}

func (t *C.MyTest) toGoStruct() MyTest {
	return MyTest{
		a: uint(t.a),
		b: uint(t.b),
	}
}

func main() {
	x := &C.MyTest{a: 10, b: 20}
	C.display(x)

	y := x.toGoStruct()

	fmt.Printf("value = %v\n", y)
}

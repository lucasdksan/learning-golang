package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	var r strings.Reader

	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("Lucas da Silva Leoncio")
	fmt.Println(b.String())

	test := b.Cap()
	l := b.Len()

	fmt.Println("Cap: ", test)
	fmt.Println("Len: ", l)

	b.Reset()

	fmt.Println("Reset: ", b.String())

	b.Write([]byte("Hello, "))
	b.Write([]byte("World!"))

	red, _ := r.Read([]byte("Hello, adasda"))

	fmt.Println("Read: ", red)

	fmt.Println("Write: ", b.String())
}

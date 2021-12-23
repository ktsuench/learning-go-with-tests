package greet

import (
	"bytes"
	"fmt"
)

func Greet(bytesBuf *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}

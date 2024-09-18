package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type MyReader struct{}

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return (b-'A'+13)%26 + 'A'
	case 'a' <= b && b <= 'z':
		return (b-'a'+13)%26 + 'a'
	default:
		return b
	}
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, nil
}

func TryReader() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 7)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r2 := rot13Reader{s}
	io.Copy(os.Stdout, &r2)
	println()
}

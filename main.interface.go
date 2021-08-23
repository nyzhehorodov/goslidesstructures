package main

import (
	"bytes"
	"fmt"
	"io"
)

func someFunc(r io.Reader) error {
	if r == nil {
		return fmt.Errorf("empty reader")
	}
	fmt.Println("ready to read some data")
	
	_, err := r.Read([]byte("some string"))
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}
	
	return nil
}

func main() {
	var b *bytes.Reader
	
	if err := someFunc(b); err != nil {
		fmt.Printf("got some error: %v", err)
	} else {
		fmt.Println("Hello, playground")
	}
}

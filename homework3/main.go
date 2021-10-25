package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	lw := logWriter{}

	io.Copy(lw, file)
}

func (logWriter) Write (bs []byte) (int, error) {
	fmt.Println("File contents:", string(bs))
	return len(bs), nil
}

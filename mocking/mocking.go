package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(writer io.Writer) {
	fmt.Fprintf(writer, "3")
}

func main() {
	Countdown(os.Stdout)
}

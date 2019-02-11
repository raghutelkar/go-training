package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, fib(7))
	
}

func fib(n int) int{
	if n <=1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

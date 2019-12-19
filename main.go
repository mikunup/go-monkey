package main

import (
	"fmt"

	"github.com/mikunup/go-monkey/lexer"
)

func main() {
	fmt.Println("vim-go")

	l := lexer.New("i")
	fmt.Println(l)
}

package main

import (
	"fmt"
	"os"
	"text/scanner"

	"github.com/abiriadev/hango"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("pass a file to transform")
	}

	d, _ := os.Open(os.Args[1])

	lex := hango.HangoLexNew(d)

	for tok, l, r := lex.Scan(); tok != scanner.EOF; tok, l, r = lex.Scan() {
		fmt.Println(l, r, scanner.TokenString(tok))
	}
}

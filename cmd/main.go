package main

import (
	"fmt"
	"os"

	"github.com/abiriadev/hango"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("pass a file to transform")
	}

	d, _ := os.ReadFile(os.Args[1])

	r := hango.HangoLexNew(string(d)).RegenerateSource()

	fmt.Print(r)
}

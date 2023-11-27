package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/abiriadev/hango"
)

func main() {
	flag.Parse()

	srcs := flag.Args()

	if len(srcs) == 0 {
		d, _ := io.ReadAll(os.Stdin)

		r := hango.HangoLexNew(string(d)).RegenerateSource()

		fmt.Print(r)
	} else {

		for _, src := range srcs {
			ext := filepath.Ext(src)

			if ext != ".hgo" {
				panic("no hgo!")
			}

			target := strings.TrimSuffix(src, ext) + ".go"

			d, e := os.ReadFile(src)

			if e != nil {
				panic(e)
			}

			r := hango.HangoLexNew(string(d)).RegenerateSource()

			os.WriteFile(target, []byte(r), 0644)
		}
	}
}

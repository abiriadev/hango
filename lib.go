package hango

import (
	"strings"
	"text/scanner"
)

type HangoLex struct {
	s   scanner.Scanner
	src string
}

func HangoLexNew(src string) HangoLex {
	var lex HangoLex
	lex.src = src
	lex.s.Init(strings.NewReader(src))
	lex.s.Mode = scanner.GoTokens &^ scanner.SkipComments
	lex.s.Whitespace ^= 1<<' ' | 1<<'\t' | 1<<'\r' | 1<<'\n'
	return lex
}

func (this *HangoLex) Scan() (rune, int, int) {
	tok := this.s.Scan()
	return tok, this.s.Position.Offset, this.s.Pos().Offset
}

func (this HangoLex) RegenerateSource() string {
	var buf strings.Builder

	for tok, l, r := this.Scan(); tok != scanner.EOF; tok, l, r = this.Scan() {
		buf.WriteString(this.src[l:r])
	}

	return buf.String()
}

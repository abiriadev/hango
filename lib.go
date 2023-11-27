package hango

import (
	"strings"
	"text/scanner"
)

type HangoLex struct {
	s scanner.Scanner
}

func HangoLexNew(src string) HangoLex {
	var lex HangoLex
	lex.s.Mode = scanner.GoTokens &^ scanner.SkipComments
	lex.s.Whitespace ^= 1<<' ' | 1<<'\t' | 1<<'\r' | 1<<'\n'
	lex.s.Init(strings.NewReader(src))
	return lex
}

func (this *HangoLex) Scan() (rune, int, int) {
	tok := this.s.Scan()
	return tok, this.s.Position.Offset, this.s.Pos().Offset
}

package hango

import (
	"io"
	"text/scanner"
)

type HangoLex struct {
	s scanner.Scanner
}

func HangoLexNew(src io.Reader) HangoLex {
	var lex HangoLex
	lex.s.Init(src)
	lex.s.Mode = scanner.GoTokens &^ scanner.SkipComments
	lex.s.Whitespace ^= 1<<' ' | 1<<'\t' | 1<<'\r' | 1<<'\n'
	return lex
}

func (this *HangoLex) Scan() (rune, int, int, string) {
	tok := this.s.Scan()
	return tok, this.s.Position.Offset, this.s.Pos().Offset, this.s.TokenText()
}

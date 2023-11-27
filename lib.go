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

var KwdMap = map[string]string{
	"정지":   "break",
	"기본":   "default",
	"함수":   "func",
	"특성":   "interface",
	"선택":   "select",
	"경우":   "case",
	"지연":   "defer",
	"출발":   "go",
	"사전":   "map",
	"구조체":  "struct",
	"통로":   "chan",
	"아니라면": "else",
	"이동":   "goto",
	"묶음":   "package",
	"비교":   "switch",
	"상수":   "const",
	"이어서":  "fallthrough",
	"만약":   "if",
	"범위":   "range",
	"형":    "type",
	"계속":   "continue",
	"반복":   "for",
	"사용":   "import",
	"반환":   "return",
	"선언":   "var",
}

func MapKeyWord(tok string) string {
	r := KwdMap[tok]
	if tok == "" {
		return tok
	} else {
		return r
	}
}

func (this HangoLex) RegenerateSource() string {
	var buf strings.Builder

	for tok, l, r := this.Scan(); tok != scanner.EOF; tok, l, r = this.Scan() {
		buf.WriteString(this.src[l:r])
	}

	return buf.String()
}

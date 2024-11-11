package lexer

import (
	"strings"
	"unicode/utf8"

	"github.com/CharukaK/request-monkey/parser/token"
)

type StateFn func(l *Lexer) StateFn

type Lexer struct {
	input   string
	start   int // starting position of the current text chunk
	pos     int // current reading position of the text chunk
	width   int // width of the rune
	stateFn StateFn
	tokens  chan token.Token
}

func (lx *Lexer) emit(t token.TokenType) {
	lx.tokens <- token.Token{
		Type:    t,
		Literal: lx.input[lx.start:lx.pos],
	}

	lx.start = lx.pos
}

func (lx *Lexer) run() {
	for state := lx.stateFn; state != nil; {
		state = state(lx)
	}
	close(lx.tokens)
}

func (lx *Lexer) next() (ch rune) {
	if lx.pos >= len(lx.input) {
		lx.width = 0
		return 0 // this will represet the EOF
	}

	ch, lx.width = utf8.DecodeRuneInString(lx.input[lx.pos:])

	lx.pos += lx.width
	return
}

func (lx *Lexer) backup() {
	lx.pos -= lx.width
}

func (lx *Lexer) ignore() {
	lx.start = lx.pos
}

func (lx *Lexer) peek() (ch rune) {
	ch = lx.next()
	lx.backup()
	return
}

// consumes the character if it is from a valid set of characters
func (lx *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, lx.next()) > 0 {
		return true
	}
	lx.backup()
	return false
}

// consumes until the lexer come across an invalid string
func (lx *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, lx.next()) > 0 {
	}
	lx.backup()
}

func initState(lx *Lexer) StateFn {
	switch lx.next() {
	case 0:
		lx.emit(token.EOF)
    case '@':
	}

	return nil
}

func New(input string) (lex *Lexer) {
	lex = &Lexer{
		input:   input,
		tokens:  make(chan token.Token, 2),
		stateFn: initState,
	}

	return
}

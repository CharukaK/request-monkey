package lexer

import (
	"fmt"
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
        fmt.Printf("%+v\n", lx.stateFn)
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

func (lx *Lexer) NextItem() token.Token {
	for {
		select {
		case token := <-lx.tokens:
			return token
		default:
			lx.stateFn = lx.stateFn(lx)
		}
	}
	panic("not reached")
}

func initState(lx *Lexer) StateFn {
	switch lx.next() {
	case 0:
		lx.emit(token.EOF)
	case '@':
		lx.emit(token.VAR_DECL_PREFIX)
		return varDeclState
	}

	return nil
}

// terminates lexer and returns a formatted error message to lexer.items
func (l *Lexer) errorf(format string, args ...interface{}) StateFn {
	msg := fmt.Sprintf(format, args...)
	start := l.pos - 10
	if start < 0 {
		start = 0
	}
	l.tokens <- token.NewToken(
		token.ILLEGAL,
		fmt.Sprintf("Error at char %d: '%s'\n%s", l.pos, l.input[start:l.pos+1], msg),
	)
	//panic("PANIC")
	return nil
}

func varDeclState(lx *Lexer) StateFn {
	return initState
}

func New(input string) (lex *Lexer) {
	lex = &Lexer{
		input:   input,
		tokens:  make(chan token.Token, 2),
		stateFn: initState,
	}

	fmt.Println(fmt.Sprintf("lexer: %+v", lex))
	go lex.run()

	return
}

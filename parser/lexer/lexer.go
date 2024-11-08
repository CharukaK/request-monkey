package lexer

import (
	"bytes"

	"github.com/CharukaK/request-monkey/parser/token"
)

type Lexer struct {
	input        string
	position     int  // current position
	readPosition int  // next character position
	ch           byte // current character
}

func (lx *Lexer) readChar() {
	if lx.readPosition >= len(lx.input) {
		lx.ch = 0 // ascii starts with 1 therefore, setting it to 0 means nil i.e. EOF
	} else {
		lx.ch = lx.input[lx.readPosition]
	}

	lx.position = lx.readPosition
	lx.readPosition += 1
}

func (lx *Lexer) NextToken() (tok token.Token) {
	lx.skipWhiteSpaces()

	switch lx.ch {
	case '#':
		lx.skipUntilNewLineOrEof()
		tok = lx.NextToken()
	case '0':
		tok.Literal = ""
		tok.Type = token.EOF
	case '\n':
		tok.Literal = ""
		tok.Type = token.NEW_LINE
	default:
		if isAlphaNumeric(lx.ch) {
			tok.Literal = lx.readAnyText()
			tok.Type = token.ANY_TEXT
		}
	}

	lx.readChar()
	return
}

func (lx *Lexer) skipWhiteSpaces() {
	for lx.ch == ' ' || lx.ch == '\t' {
		lx.readChar()
	}
}

func (lx *Lexer) readAnyText() (val string) {
	buf := bytes.NewBuffer(make([]byte, 0))
	for isAlphaNumeric(lx.ch) {
		if err := buf.WriteByte(lx.ch); err != nil {
			panic(err)
		}
		lx.readChar()
	}
	val = buf.String()
	return
}

func (lx *Lexer) skipUntilNewLineOrEof() {
	for lx.ch != '\r' && lx.ch != '\n' && lx.ch != '0' {
		lx.readChar()
	}
}

func isAlphaNumeric(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func New(input string) (l *Lexer) {
	l = &Lexer{input: input}
	l.readChar()
	return
}

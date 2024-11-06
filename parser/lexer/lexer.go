package lexer

import "github.com/CharukaK/request-monkey/parser/token"

type Lexer struct {
	input        string
	position     int  // points to the position of the ch field
	readPosition int  // next read position
	ch           byte // character currently read
}

func (lx *Lexer) readChar() {
	if lx.readPosition >= len(lx.input) {
		lx.ch = 0
	} else {
		lx.ch = lx.input[lx.readPosition]
	}

	lx.position = lx.readPosition
	lx.readPosition++
}

func (lx *Lexer) NextToken() (tok token.Token) {
	lx.skipWhiteSpaces()
	switch lx.ch {
	case '=':
		tok = newToken(token.ASSIGN, token.LiteralMap[token.ASSIGN])
	case '@':
		tok = newToken(token.IDENT_PREFIX, token.LiteralMap[token.IDENT_PREFIX])
	default:
		if isAlphaNumeric(lx.ch) {
			value := lx.getAlphaNumericValue()

			if len(value) > 0 {
				if tt, ok := token.KeywordMap[value]; ok {
					tok = newToken(tt, value)
				} else {
					tok = newToken(token.VALUE, value)
				}
			}
		}
	}

	lx.readChar()

	return
}

func (lx *Lexer) getAlphaNumericValue() string {
	var val string

	for isAlphaNumeric(lx.ch) {
		val += string(lx.ch)
		lx.readChar()
	}

	return val
}

func newToken(tt token.TokenType, literal string) token.Token {
	return token.Token{
		Type:    tt,
		Literal: literal,
	}
}

func isAlphaNumeric(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func (lx *Lexer) skipWhiteSpaces() {
	for lx.ch == ' ' || lx.ch == '\r' || lx.ch == '\n' || lx.ch == '\t' {
		lx.readChar()
	}
}

func New(input string) (lx *Lexer) {
	lx = &Lexer{
		input: input,
	}

	lx.readChar()

	return
}

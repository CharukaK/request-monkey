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
	switch lx.ch {
	case '=':
		tok = newToken(token.ASSIGN, token.LiteralMap[token.ASSIGN])
	case '@':
		tok = newToken(token.IDENT_PREFIX, token.LiteralMap[token.IDENT_PREFIX])
	default:
	}

	lx.readChar()

	return
}

func newToken(tt token.TokenType, literal string) token.Token {
	return token.Token{
		Type:    tt,
		Literal: literal,
	}
}

func New(input string) (lx *Lexer) {
	lx = &Lexer{
		input: input,
	}

	lx.readChar()

	return
}

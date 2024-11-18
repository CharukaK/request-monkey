package lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/CharukaK/request-monkey/parser/token"
)

const (
	from_url = iota
	from_keyval
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

func (lx *Lexer) ifNotAccept(invalid string) bool {
	if strings.IndexRune(invalid, lx.next()) > 0 {
		lx.backup()
		return false
	}

	return true
}

// consumes until the lexer come across an invalid string
func (lx *Lexer) acceptAndRun(valid string) {
	for strings.IndexRune(valid, lx.next()) > 0 {
	}
	lx.backup()
}

func (lx *Lexer) ifNotAcceptAndRun(invalid string) {
	for !(strings.IndexRune(invalid, lx.next()) > 0) {
	}

	lx.backup()
}

func (lx *Lexer) NextToken() token.Token {
	// the states can't have more than 2 emits it will block the main thread and hang
	for {
		select {
		case val := <-lx.tokens:

			return val
		default:
			lx.stateFn = lx.stateFn(lx)
		}
	}

	panic("Next item not reached")
}

func initState(lx *Lexer) StateFn {
	switch lx.next() {
	case 0:
		lx.emit(token.EOF)
	case '@':
		lx.backup()
		return varDeclState
	case '#':
		lx.backup()
		return commentState
	case ' ', '\t', '\n':
		lx.ignore()
		return initState
	case 'G', 'D', 'P', 'H', 'O':
		lx.backup()
		return requestMethodState
	default:
		return lx.errorf("invalid character found!")
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
	if lx.next() != '@' {
		return lx.errorf("expected character '@'")
	}

	lx.emit(token.VAR_DECL_PREFIX)

	for lx.accept("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789") {
		// generate the name string
	}

	lx.emit(token.VAR_NAME)

	lx.ignoreWhiteSpaces()

	for {
		ch := lx.next()

		if ch == '=' {
			lx.emit(token.ASSIGN)
			break
		} else {
			return lx.errorf("expected symbol '='")
		}
	}

	lx.ignoreWhiteSpaces()

	for {
		ch := lx.next()
		if ch == '\n' || ch == 0 {
			break
		}
	}

	lx.backup()

	if lx.start != lx.pos {
		lx.emit(token.VAR_VALUE)
	}

	return initState
}

func (lx *Lexer) ignoreWhiteSpaces() {
	if lx.next() == ' ' {
		lx.ignore()
		lx.ignoreWhiteSpaces()
	} else {
		lx.backup()
	}
}

func commentState(lx *Lexer) StateFn {
	for {
		ch := lx.next()
		if ch == '\n' || ch == 0 {
			break
		}
	}

	lx.backup()
	return initState
}

func requestMethodState(lx *Lexer) StateFn {
	// process the request method statements

	for {
		ch := lx.next()

		if ch == ' ' || ch == '\n' || ch == 0 {
			lx.backup()
			break
		}
	}

	if strings.Index("POST GET PUT DELETE PATCH HEAD CONNECT OPTIONS TRACE", lx.input[lx.start:lx.pos]) > -1 {
		lx.emit(token.METHOD)
		if lx.peek() != '\n' {
			return urlState
		}
	} else {
		return lx.errorf("invalid method type '%s'", lx.input[lx.start:lx.pos])
	}

	return requestBodyState
}

func urlState(lx *Lexer) StateFn {
	for {
		ch := lx.next()
		if ch == '{' && lx.peek() == '{' {

		} else if ch == ' ' || ch == '\n' || ch == 0 {
			lx.backup()
			break
		}
	}

	return requestBodyState
}

func valueInsertState(lx *Lexer, from int) StateFn {
	for {
		ch := lx.next()

		if ch == '}' && ch == '}' {
		} else if ch == ' ' || ch == '\n' || ch == 0 {
			break
		}
	}
	switch from {
	case from_url:
		return urlState
	case from_keyval:
	}

	return nil
}

func requestBodyState(lx *Lexer) StateFn {
	// process header values
	// process payload

	return initState
}

func New(input string) (lex *Lexer) {
	lex = &Lexer{
		input:   input,
		tokens:  make(chan token.Token, 5),
		stateFn: initState,
	}

	fmt.Println(fmt.Sprintf("lexer: %+v", lex))
	// go lex.run()

	return
}

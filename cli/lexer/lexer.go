package lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/CharukaK/request-monkey/cli/token"
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

	lx.emit(token.IDENTIFIER)

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
		} else {
			lx.next()
			lx.ignore()
		}
	} else {
		return lx.errorf("invalid method type '%s'", lx.input[lx.start:lx.pos])
	}

	return requestBodyState
}

func urlState(lx *Lexer) StateFn {
	lx.ignoreWhiteSpaces()
	for {
		ch := lx.next()
		if ch == '{' && lx.peek() == '{' {
			lx.backup()
			if lx.start != lx.pos {
				lx.emit(token.URL_SEGMENT)
			}
			lx.next()
			lx.next()
			lx.emit(token.LBRACE)
			return valueInsertState(lx, from_url)
		} else if ch == ' ' {
			lx.backup()
			lx.emit(token.URL_SEGMENT)
			return httpVersionState
		} else if ch == '\n' {
			lx.backup()
			lx.emit(token.URL_SEGMENT)
			lx.next()
			lx.ignore()
			break
		} else if ch == 0 {
			lx.backup()
			lx.emit(token.URL_SEGMENT)
			break
		}
	}

	return requestBodyState
}

func httpVersionState(lx *Lexer) StateFn {
	lx.ignoreWhiteSpaces()
	for {
		ch := lx.next()
		if ch == ' ' || ch == '\n' || ch == 0 {
			lx.backup()
			if strings.Index("HTTP/1.1 HTTP/2.0", lx.input[lx.start:lx.pos]) > -1 {
				lx.emit(token.HTTP_VERSION)
			} else {
				return lx.errorf("Expected a valid http version")
			}

			if lx.peek() == '\n' {
				lx.next()
				lx.ignore()
			}

			break
		}
	}

	return requestBodyState
}

func valueInsertState(lx *Lexer, from int) StateFn {
	return func(l *Lexer) StateFn {
		for {
			ch := lx.next()
			if ch == '}' && lx.peek() == '}' {
				lx.backup()
				if lx.start != lx.pos {
					l.emit(token.IDENTIFIER)
				}
				lx.next()
				lx.next()
				l.emit(token.RBRACE)
				break
			} else if ch == '\n' || ch == 0 {
				return lx.errorf("expected closing braces '}}'")
			}
		}

		switch from {
		case from_url:
			return urlState
		case from_keyval:
			return headerValueState
		}

		return nil
	}
}

func requestBodyState(lx *Lexer) StateFn {
	// process header values
	// process payload

	lx.ignoreWhiteSpaces()
	for {
		ch := lx.next()
		if ch == ':' {
			lx.backup()
			lx.emit(token.HEADER_KEY)
			lx.next()
			lx.emit(token.COLON)
			return headerValueState
		} else if ch == '\n' || ch == 0 {
			lx.backup()
			if lx.start != lx.pos {
				lx.emit(token.HEADER_KEY)
				return lx.errorf("expected symbol ':'")
			}
			break
		}
	}

	if lx.next() == '\n' {
		lx.ignore()
		lx.ignoreWhiteSpaces()
		if lx.peek() == '\n' {
			lx.next()
			lx.ignore()
			return payloadState
		}
	}

	return initState
}

func headerValueState(lx *Lexer) StateFn {
	lx.ignoreWhiteSpaces()

	for {
		ch := lx.next()
		if ch == '{' && lx.peek() == '{' {
			lx.backup()
			if lx.start != lx.pos {
				lx.emit(token.HEADER_VAL_SEGMENT)
			}
			lx.next()
			lx.next()
			lx.emit(token.LBRACE)
			return valueInsertState(lx, from_keyval)
		} else if ch == '\n' {
			lx.backup()
			if lx.start != lx.pos {
				lx.emit(token.HEADER_VAL_SEGMENT)
			}
			lx.next()
			lx.ignore()

			lx.ignoreWhiteSpaces()

			if lx.next() == '\n' {
				lx.ignore()
				return payloadState
			}

            return requestBodyState
		} else if ch == 0 {
			lx.backup()
			if lx.start != lx.pos {
				lx.emit(token.HEADER_VAL_SEGMENT)
			}
			break
		}
	}

	return requestBodyState
}

func payloadState(lx *Lexer) StateFn {
	for {
		ch := lx.next()

		if ch == '\n' || ch == 0 {
			lx.backup()
			lx.emit(token.PAYLOAD_SEGMENT)
			lx.next()
			lx.ignore()
			return payloadState
		}

        val := lx.input[lx.start:lx.pos]

        if strings.HasPrefix(val, "POST") ||
        	strings.HasPrefix(val, "GET") ||
        	strings.HasPrefix(val, "PUT") ||
        	strings.HasPrefix(val, "DELETE") ||
        	strings.HasPrefix(val, "PATCH") ||
        	strings.HasPrefix(val, "HEAD") ||
        	strings.HasPrefix(val, "CONNECT") ||
        	strings.HasPrefix(val, "OPTIONS") ||
        	strings.HasPrefix(val, "TRACE") ||
        	strings.HasPrefix(val, "#") ||
        	strings.HasPrefix(val, "@") {
        	lx.pos = lx.start
        	break
        }
	}

	return initState
}

func New(input string) (lex *Lexer) {
	lex = &Lexer{
		input:   input,
		tokens:  make(chan token.Token, 5),
		stateFn: initState,
	}

	return
}

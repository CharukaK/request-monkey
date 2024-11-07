package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

/**
Token types based on the syntax


*/

const (
	ILLEGAL TokenType = iota // tokens or characters that can't be figured out
	EOF                      // let the parser know when to stop
	NEW_LINE
	COMMENT

	ESC_CHAR // escape character `\`

	VAR_DECL_PREFIX // character `@`
	COLON           // character `:`
	ASSIGN          // character `=`

	ANY_TEXT // any type of text url sections, values, keys, payloads

	// methods
	POST
	GET
	PUT
	DELETE
	PATCH
	HEAD
	CONNECT
	OPTIONS
	TRACE
)

var KeywordMap = map[string]TokenType{
	"POST":    POST,
	"GET":     GET,
	"PUT":     PUT,
	"DELETE":  DELETE,
	"PATCH":   PATCH,
	"HEAD":    HEAD,
	"CONNECT": CONNECT,
	"OPTIONS": OPTIONS,
	"TRACE":   TRACE,
}

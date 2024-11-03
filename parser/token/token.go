package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = iota // tokens or characters that can't be figured out
	EOF                      // let the parser know when to stop

	IDENT // identifires (variable names)

	ASSIGN       // assignment operator `=`
	IDENT_PREFIX // variable declaration prefix `@`
	LBRACE       // value insert start `{{`
	RBRACE       // value insert start `}}`

	ESC_CHAR // escape character `\`

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

var LiteralMap = map[TokenType]string{
	ASSIGN:       "=",
	IDENT_PREFIX: "@",
	LBRACE:       "{{",
	RBRACE:       "}}",
	ESC_CHAR:     "\\",
	POST:         "POST",
	GET:          "GET",
	PUT:          "PUT",
	DELETE:       "DELETE",
	PATCH:        "PATCH",
	HEAD:         "HEAD",
	CONNECT:      "CONNECT",
	OPTIONS:      "OPTIONS",
	TRACE:        "TRACE",
}

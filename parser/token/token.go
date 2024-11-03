package token

type TokenType int

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

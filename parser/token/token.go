package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = iota // tokens or characters that can't be figured out
	EOF                      // let the parser know when to stop
	NEW_LINE
	COMMENT

	ESC_CHAR // escape character `\`

	VAR_DECL_PREFIX // character `@`
	COLON           // character `:`
	ASSIGN          // character `=`

	VAR_NAME   // variable name declaration
	VAR_VALUE  // value section of the variables
	IDENTIFIER // value identifiers of the value inserts

	LBRACE // value insert start `{{`
	RBRACE // value insert end `}}`

	METHOD // methods POST GET PUT DELETE PATCH HEAD CONNECT OPTIONS TRACE
    URL_SEGMENT
    HTTP_VERSION
    HEADER_KEY
    HEADER_VAL_SEGMENT
)

func NewToken(tt TokenType, value string) Token {
	return Token{
		Type:    tt,
		Literal: value,
	}
}

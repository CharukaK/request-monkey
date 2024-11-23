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

var tokenTypeToString = map[TokenType]string{
	ILLEGAL:            "ILLEGAL",
	EOF:                "EOF",
	NEW_LINE:           "NEW_LINE",
	COMMENT:            "COMMENT",
	ESC_CHAR:           "ESC_CHAR",
	VAR_DECL_PREFIX:    "VAR_DECL_PREFIX",
	COLON:              "COLON",
	ASSIGN:             "ASSIGN",
	VAR_NAME:           "VAR_NAME",
	VAR_VALUE:          "VAR_VALUE",
	IDENTIFIER:         "IDENTIFIER",
	LBRACE:             "LBRACE",
	RBRACE:             "RBRACE",
	METHOD:             "METHOD",
	URL_SEGMENT:        "URL_SEGMENT",
	HTTP_VERSION:       "HTTP_VERSION",
	HEADER_KEY:         "HEADER_KEY",
	HEADER_VAL_SEGMENT: "HEADER_VAL_SEGMENT",
}

func GetTokenTypeString(tType TokenType) string {
	if val, ok := tokenTypeToString[tType]; ok {
		return val
	}

	return "unknown"
}

func NewToken(tt TokenType, value string) Token {
	return Token{
		Type:    tt,
		Literal: value,
	}
}

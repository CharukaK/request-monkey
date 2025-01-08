package parser

import (
	"fmt"
	"testing"

	"github.com/CharukaK/request-monkey/cli/lexer"
)

func TestVarDecl(t *testing.T) {
	input := `@hello=asdfkjasdlfkj`

	p := NewParser(*lexer.New(input))

    p.Parse()
    
    fmt.Println(p.document)

	// for _, tc := range testcases {
	// 	tok := l.NextToken()
	//
	// 	if tok.Type != tc.expectedType {
	// 		t.Fatalf(`Token type mismatch: expected '%d', got '%d'`, tc.expectedType, tok.Type)
	// 	}
	//
	// 	if tok.Literal != tc.expectedLiteral {
	// 		t.Fatalf(`Token value mismatch: expected '%s', got '%s'`, tc.expectedLiteral, tok.Literal)
	// 	}
	// }

}


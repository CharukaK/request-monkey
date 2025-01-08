package parser

import (
	"testing"

	"github.com/CharukaK/request-monkey/cli/ast"
	"github.com/CharukaK/request-monkey/cli/lexer"
)

func TestVarDecl(t *testing.T) {
	input := `@hello=asdfkjasdlfkj`

	p := NewParser(*lexer.New(input))

    p.Parse()
    
    document := p.document


    if len(document.Statements) != 1 {
        t.Fatal("mismatch on number of statements")
    }

    stmt, ok := document.Statements[0].(*ast.Variable)

    if !ok {
        t.Fatal("invalid type found for statement")
    }

    if stmt.Name.String() != "hello" || stmt.Value.String() != "asdfkjasdlfkj" {
        t.Fatal("invalid name and value combintaion detected")
    }



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


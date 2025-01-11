package parser

import (
	"fmt"
	"testing"

	"github.com/CharukaK/request-monkey/cli/ast"
	"github.com/CharukaK/request-monkey/cli/lexer"
)

func TestVarDecl(t *testing.T) {
	input := `
    @hello=asdfkjasdlfkj
    @test=hello
    `

	p := NewParser(*lexer.New(input))

    p.Parse()
    
    document := p.document


    if len(document.Statements) != 2 {
        t.Fatal("mismatch on number of statements")
    }

    stmt1, ok := document.Statements[0].(*ast.Variable)

    if !ok {
        t.Fatal("invalid type found for statement")
    }

    if stmt1.Name.String() != "hello" || stmt1.Value.String() != "asdfkjasdlfkj" {
        t.Fatal("invalid name and value combintaion detected")
    }

    stmt2, ok := document.Statements[1].(*ast.Variable)

    if !ok {
        t.Fatal("invalid type found for statement")
    }

    if stmt2.Name.String() != "test" || stmt2.Value.String() != "hello" {
        t.Fatal("invalid name and value combintaion detected")
    }
}

func TestRequestDecl(t *testing.T) {
    input := `
    GET asdf/asdfj


    POST asdfjsdfk/asdfjk
    Auth: Hehe
    `

    p := NewParser(*lexer.New(input))

    p.Parse()

    document := p.document

    fmt.Println(document)
}


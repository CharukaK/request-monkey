package parser

import (
	"fmt"

	"github.com/CharukaK/request-monkey/cli/ast"
	"github.com/CharukaK/request-monkey/cli/lexer"
	"github.com/CharukaK/request-monkey/cli/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	document  *ast.Document
	currToken token.Token
	peekToken token.Token
	errors    []token.Token
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lexer.NextToken()

	fmt.Println(">>> curr", p.currToken)
	fmt.Println(">>> peek", p.peekToken)
}

func (p *Parser) Parse() {
	p.document.Statements = make([]ast.Statement, 0)
	for {
		switch p.currToken.Type {
		case token.EOF:
			return
		case token.ILLEGAL:
			p.errors = append(p.errors, p.currToken)
		default:
			if stmt := p.parseStatement(); stmt != nil {
				p.document.Statements = append(p.document.Statements, stmt)
			} else {
				p.nextToken()
			}
		}
	}
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.VAR_DECL_PREFIX:
		return p.parseVariable()
	case token.METHOD:
	default:
	}

	return nil
}

func (p *Parser) parseVariable() *ast.Variable {
	varDecl := &ast.Variable{}

	if p.currToken.Type == token.VAR_DECL_PREFIX {
		p.nextToken()
	}

	if p.currToken.Type == token.IDENTIFIER {
		varDecl.Name.Text = p.currToken.Literal
		p.nextToken()
	}

	for {
		switch p.currToken.Type {
		case token.VAR_DECL_PREFIX:
			p.nextToken()
		case token.VAR_NAME:
			varDecl.Name.Text = p.currToken.Literal
			p.nextToken()
		case token.ASSIGN:
			p.nextToken()
		case token.VAR_VALUE:
			varDecl.Value.Parts = append(varDecl.Value.Parts, &ast.LiteralValue{Text: p.currToken.Literal})
			return varDecl
		default:
			return nil
		}
	}

}

func NewParser(lexer lexer.Lexer) (p Parser) {
	p.lexer = &lexer
	p.document = &ast.Document{}

	p.nextToken()
	p.nextToken()

	return
}

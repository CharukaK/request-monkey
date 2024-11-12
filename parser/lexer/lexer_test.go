package lexer

import (
	"testing"

	"github.com/CharukaK/request-monkey/parser/token"
)

func TestVarDecl(t *testing.T) {
	input := `@hello=asdfkjasdlfkj`

	testcases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{expectedType: token.VAR_DECL_PREFIX, expectedLiteral: "@"},
		// {expectedType: token.VAR_NAME, expectedLiteral: "hello"},
		// {expectedType: token.ASSIGN, expectedLiteral: "="},
		// {expectedType: token.VAR_VALUE, expectedLiteral: "asdfkjasdlfkj"},
	}

	l := New(input)

	for _, tc := range testcases {
		tok := l.NextItem()

		if tok.Type != tc.expectedType {
			t.Fatalf(`Token type mismatch: expected '%d', got '%d'`, tc.expectedType, tok.Type)
		}

		if tok.Literal != tc.expectedLiteral {
			t.Fatalf(`Token value mismatch: expected '%s', got '%s'`, tc.expectedLiteral, tok.Literal)
		}
	}

}

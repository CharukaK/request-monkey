package lexer

import (
	"testing"

	"github.com/CharukaK/request-monkey/parser/token"
)

func TestCommentStrings(t *testing.T) {
	input := `# hello world`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.COMMENT, "#"},
		{token.ANY_TEXT, " hello world"},
	}


	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				`test [%d] - token type is wrong, expected='%s', got='%s'`,
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				`test [%d] - token literal is wrong, expected='%s', got='%s'`,
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}


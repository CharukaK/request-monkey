package lexer

import (
	"testing"

	"github.com/CharukaK/request-monkey/parser/token"
)

// ignore comments
func TestCommentStrings(t *testing.T) {
	input := `
    # hello world
    someText # comment after text
    # test 123
    `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{

		{expectedType: token.NEW_LINE, expectedLiteral: ""},
		{expectedType: token.NEW_LINE, expectedLiteral: ""},
		{expectedType: token.ANY_TEXT, expectedLiteral: "someText"},
		{expectedType: token.NEW_LINE, expectedLiteral: ""},
		{expectedType: token.NEW_LINE, expectedLiteral: ""},
		{expectedType: token.EOF, expectedLiteral: ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				`test [%d] - token type is wrong, expected='%d', got='%d'`,
				i,
				tt.expectedType,
				tok.Type,
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

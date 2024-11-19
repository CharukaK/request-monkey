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
		{expectedType: token.VAR_NAME, expectedLiteral: "hello"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.VAR_VALUE, expectedLiteral: "asdfkjasdlfkj"},
	}

	l := New(input)

	for _, tc := range testcases {
		tok := l.NextToken()

		if tok.Type != tc.expectedType {
			t.Fatalf(`Token type mismatch: expected '%d', got '%d'`, tc.expectedType, tok.Type)
		}

		if tok.Literal != tc.expectedLiteral {
			t.Fatalf(`Token value mismatch: expected '%s', got '%s'`, tc.expectedLiteral, tok.Literal)
		}
	}

}

func TestCommentLine(t *testing.T) {
	input := `
    # below is a variable
    @hello=asdfkjasdlfkj
    `

	testcases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{expectedType: token.VAR_DECL_PREFIX, expectedLiteral: "@"},
		{expectedType: token.VAR_NAME, expectedLiteral: "hello"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.VAR_VALUE, expectedLiteral: "asdfkjasdlfkj"},
	}

	l := New(input)

	for _, tc := range testcases {
		tok := l.NextToken()

		if tok.Type != tc.expectedType {
			t.Fatalf(`Token type mismatch: expected '%d', got '%d'`, tc.expectedType, tok.Type)
		}

		if tok.Literal != tc.expectedLiteral {
			t.Fatalf(`Token value mismatch: expected '%s', got '%s'`, tc.expectedLiteral, tok.Literal)
		}
	}
}

func TestRequestDecl(t *testing.T) {
	input := `
    # Example of a .req file with variables and request declarations

    @host = api.example.com
    @contentType = application/json
    @token = abc123

    # Request 1
    POST http://{{host}}/users HTTP/1.1
    Authorization: Bearer {{token}}
    Content-Type: {{contentType}}

    {
        "name": "John Doe",
        "email": "john.doe@example.com"
    }
    `

	testcases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{expectedType: token.VAR_DECL_PREFIX, expectedLiteral: "@"},
		{expectedType: token.VAR_NAME, expectedLiteral: "host"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.VAR_VALUE, expectedLiteral: "api.example.com"},

		{expectedType: token.VAR_DECL_PREFIX, expectedLiteral: "@"},
		{expectedType: token.VAR_NAME, expectedLiteral: "contentType"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.VAR_VALUE, expectedLiteral: "application/json"},

		{expectedType: token.VAR_DECL_PREFIX, expectedLiteral: "@"},
		{expectedType: token.VAR_NAME, expectedLiteral: "token"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.VAR_VALUE, expectedLiteral: "abc123"},

		{expectedType: token.METHOD, expectedLiteral: "POST"},
		{expectedType: token.URL_SEGMENT, expectedLiteral: "http://"},
		{expectedType: token.LBRACE, expectedLiteral: "{{"},
		{expectedType: token.IDENTIFIER, expectedLiteral: "host"},
		{expectedType: token.RBRACE, expectedLiteral: "}}"},
		{expectedType: token.URL_SEGMENT, expectedLiteral: "/users"},

		{expectedType: token.HTTP_VERSION, expectedLiteral: "HTTP/1.1"},
		//
		// {expectedType: token.HEADER_KEY, expectedLiteral: "Authorization"},
		// {expectedType: token.HEADER_VAL_SEGMENT, expectedLiteral: "Bearer {{token}}"},
		//
		// {expectedType: token.HEADER_KEY, expectedLiteral: "Content-Type"},
		// {expectedType: token.HEADER_VAL_SEGMENT, expectedLiteral: "{{contentType}}"},

	}

	l := New(input)

	for _, tc := range testcases {
		tok := l.NextToken()

		if tok.Type != tc.expectedType {
			t.Fatalf(`Token type mismatch: expected '%d', got '%d'`, tc.expectedType, tok.Type)
		}

		if tok.Literal != tc.expectedLiteral {
			t.Fatalf(`Token value mismatch: expected '%s', got '%s'`, tc.expectedLiteral, tok.Literal)
		}
	}
}

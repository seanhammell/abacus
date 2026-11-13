package calculator

import (
	"testing"
)

func TestTokens(t *testing.T) {
	input := "0+-x/"
	tests := []struct {
		tokenType    string
		tokenLiteral string
	}{
		{OPERAND, "0"},
		{PLUS, "+"},
		{MINUS, "-"},
		{TIMES, "x"},
		{SLASH, "/"},
	}

	lexer := NewLexer(input)

	for i, test := range tests {
		token := lexer.NextToken()

		if token.Type != test.tokenType {
			t.Fatalf("tests[%d]: Wrong token type; expected %q, got %q", i, test.tokenType, token.Type)
		}

		if token.Literal != test.tokenLiteral {
			t.Fatalf("tests[%d]: Wrong token literal; expected %q, got %q", i, test.tokenLiteral, token.Literal)
		}
	}
}

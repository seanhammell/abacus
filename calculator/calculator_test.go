package calculator

import (
	"testing"
)

func TestTokens(t *testing.T) {
	input := "abc 01 + !@#$% 23 - 45 x 67 / 89"
	tests := []struct {
		tokenType    string
		tokenLiteral string
	}{
		{OPERAND, "01"},
		{PLUS, "+"},
		{OPERAND, "23"},
		{MINUS, "-"},
		{OPERAND, "45"},
		{TIMES, "x"},
		{OPERAND, "67"},
		{SLASH, "/"},
		{OPERAND, "89"},
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

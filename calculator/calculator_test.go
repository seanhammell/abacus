package calculator

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	input := "a 01 + ! 23 - 45 x 67 / 89"
	tests := []struct {
		tokenType    string
		tokenLiteral string
	}{
		{INVALID, "a"},
		{OPERAND, "01"},
		{PLUS, "+"},
		{INVALID, "!"},
		{OPERAND, "23"},
		{MINUS, "-"},
		{OPERAND, "45"},
		{TIMES, "x"},
		{OPERAND, "67"},
		{SLASH, "/"},
		{OPERAND, "89"},
		{EOF, ""},
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

func TestParser(t *testing.T) {
	input := []Token{
		{Type: OPERAND, Literal: "4"},
		{Type: PLUS, Literal: "+"},
		{Type: OPERAND, Literal: "4"},
		{Type: TIMES, Literal: "x"},
		{Type: OPERAND, Literal: "3"},
		{Type: MINUS, Literal: "-"},
		{Type: OPERAND, Literal: "2"},
	}
	expected := &ASTNode{
		Token{Type: MINUS, Literal: "-"},
		&ASTNode{
			Token{Type: PLUS, Literal: "+"},
			&ASTNode{Token{Type: OPERAND, Literal: "4"}, nil, nil},
			&ASTNode{
				Token{Type: TIMES, Literal: "x"},
				&ASTNode{Token{Type: OPERAND, Literal: "4"}, nil, nil},
				&ASTNode{Token{Type: OPERAND, Literal: "3"}, nil, nil},
			},
		},
		&ASTNode{Token{Type: OPERAND, Literal: "2"}, nil, nil},
	}

	ast := Parse(input)
	if !checkAst(ast, expected, 0) {
		t.Fatal("Parsed AST does not match expected AST")
	}
}

func checkAst(ast *ASTNode, expected *ASTNode, depth int) bool {
	if ast == nil || expected == nil {
		if ast != expected {
			fmt.Printf("    Non-matching nil at depth %v; expected: %q, got: %q\n", depth, expected.Token, ast.Token)
			return false
		} else {
			return true
		}
	}

	if ast.Token != expected.Token {
		fmt.Printf("    Failure at depth %v; expected: %q, got: %q\n", depth, expected.Token, ast.Token)
		return false
	}

	return checkAst(ast.Left, expected.Left, depth+1) && checkAst(ast.Right, expected.Right, depth+1)
}

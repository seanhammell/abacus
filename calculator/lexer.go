package calculator

import (
	"strings"
)

type Lexer struct {
	input    string
	position int
	next     int
	current  byte
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.input = strings.ReplaceAll(lexer.input, " ", "")
	lexer.read()
	return lexer
}

func (lexer *Lexer) NextToken() Token {
	var token Token

	switch lexer.current {
	case '+':
		token = newToken(PLUS, lexer.current)
	case '-':
		token = newToken(MINUS, lexer.current)
	case 'x':
		token = newToken(TIMES, lexer.current)
	case '/':
		token = newToken(SLASH, lexer.current)
	case 0:
		token.Type = EOF
		token.Literal = ""
	default:
		if isDigit(lexer.current) {
			token = newToken(OPERAND, lexer.current)
			for isDigit(lexer.peek()) {
				lexer.read()
				token.Literal += string(lexer.current)
			}
		} else {
			token = newToken(INVALID, lexer.current)
		}
	}

	lexer.read()
	return token
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func newToken(tokenType string, tokenLiteral byte) Token {
	return Token{Type: tokenType, Literal: string(tokenLiteral)}
}

func (lexer *Lexer) peek() byte {
	if lexer.next >= len(lexer.input) {
		return 0
	}

	return lexer.input[lexer.next]
}

func (lexer *Lexer) read() {
	if lexer.next >= len(lexer.input) {
		lexer.current = 0
	} else {
		lexer.current = lexer.input[lexer.next]
	}

	lexer.position = lexer.next
	lexer.next += 1
}

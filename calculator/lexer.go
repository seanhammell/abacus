package calculator

import "strings"

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

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
	token.Literal = string(lexer.current)
	switch lexer.current {
	case '+':
		token.Type = PLUS
	case '-':
		token.Type = MINUS
	case 'x':
		token.Type = TIMES
	case '/':
		token.Type = SLASH
	default:
		if isDigit(lexer.current) {
			token.Type = OPERAND
			for isDigit(lexer.peek()) {
				lexer.read()
				token.Literal += string(lexer.current)
			}
		} else {
			token.Type = INVALID
		}
	}

	lexer.read()
	return token
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

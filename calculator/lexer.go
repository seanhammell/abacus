package calculator

type Lexer struct {
	input    string
	position int
	next     int
	current  byte
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.read()
	return lexer
}

func (lexer *Lexer) NextToken() Token {
	var token Token

	switch lexer.current {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		token.Type = OPERAND
	case '+':
		token.Type = PLUS
	case '-':
		token.Type = MINUS
	case 'x':
		token.Type = TIMES
	case '/':
		token.Type = SLASH
	}

	token.Literal = string(lexer.current)
	lexer.read()
	return token
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

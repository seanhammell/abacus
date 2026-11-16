package calculator

type Token struct {
	Type    string
	Literal string
}

const (
	OPERAND = "OPERAND"

	PLUS  = "ADD"
	MINUS = "SUBTRACT"
	TIMES = "MULTIPLY"
	SLASH = "DIVIDE"

	INVALID = "INVALID"
	EOF     = "EOF"
)

func (token *Token) Precedence() int {
	switch token.Type {
	case PLUS, MINUS:
		return 1
	case TIMES, SLASH:
		return 2
	default:
		return 0
	}
}

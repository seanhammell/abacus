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
)

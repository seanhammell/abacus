package calculator

type ASTNode struct {
	Token Token
	Left  *ASTNode
	Right *ASTNode
}

func Parse(tokens []Token) *ASTNode {
	var ast []*ASTNode
	var operators []Token
	for _, token := range tokens {
		switch token.Type {
		case OPERAND:
			ast = append(ast, &ASTNode{token, nil, nil})
		case PLUS, MINUS, TIMES, SLASH:
			if len(operators) > 0 {
				top := operators[len(operators)-1]
				if token.Precedence() < top.Precedence() {
					ast = clearOperatorStack(ast, &operators)
				}
			}
			operators = append(operators, token)
		default:
		}
	}

	ast = clearOperatorStack(ast, &operators)

	return ast[0]
}

func clearOperatorStack(ast []*ASTNode, operators *[]Token) []*ASTNode {
	copy := *operators
	*operators = nil
	for len(copy) > 0 {
		node := &ASTNode{
			Token: copy[len(copy)-1],
			Left:  ast[len(ast)-2],
			Right: ast[len(ast)-1],
		}
		ast = ast[:len(ast)-2]
		ast = append(ast, node)
		copy = copy[:len(copy)-1]
	}
	return ast
}

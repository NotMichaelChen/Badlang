package ExpressionEvaluator

/**
 * TokenType is an enum that represents what a possible token could be in an expression
 */
type TokenType int

const (
	Character TokenType = iota
	Number
	Operator
	Err
)

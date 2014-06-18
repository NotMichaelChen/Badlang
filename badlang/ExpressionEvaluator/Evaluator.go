package ExpressionEvaluator

import (
	"badlang/Structures"
)

func Eval(expression string) string {
	return RPNEval(ShuntingYard(ParseExpression(expression)))
}

/**
 * Evaluates a(n) RPN expression; algorithm courtesy of wikipedia
 * Number rounding also occurs here
 *
 * @expression: The array of rpn string tokens
 * @string: The result of the evaluation
 */
func RPNEval(expression []string) string {

	var token string
	var stack Structures.Stack

	for i := 0; i < len(expression); i++ {
		token = expression[i]

		if len(token) == 0 {
			continue
		}

		if GetType(token[0]) == Number || GetType(token[0]) == Character {
			stack.Push(token)
		} else if GetType(token[0]) == Operator {
			if stack.Size() == 0 || (stack.Size() == 1 && token != "!") {
				panic("Error: Not enough values in expression (RPNEval)")
			}
			if token == "!" {
				var tempval string
				tempval = basiccalc(token, stack.Pop(), "")
				stack.Push(tempval)
			} else {
				tempval1 := stack.Pop()
				tempval2 := stack.Pop()
				//Reverse the numbers: first number popped is second in the expression; idk why
				result := basiccalc(token, tempval2, tempval1)
				stack.Push(result)
			}
		} else {
			panic("Error: Token type is invalid (RPNEval)")
		}
	}

	stack.List = Trim(stack.List) //and yet another band-aid solution, which is even worse than before

	if stack.Size() == 1 {
		return stack.Pop()
	} else {
		panic("Error: Too many values (RPNEval)")
	}
}

/**
 * Converts an infix expression (divided up into sections) into a postfix expression to be evaluated later
 * Implementation of the Shunting-yard algorithm; taken from wikipedia
 * Function calls are not implemented (yet)
 *
 * @items: The list of infix tokens to be evaluated
 * @[]string: The list of tokens in reverse polish notation
 */
func ShuntingYard(items []string) []string {
	var stack Structures.Stack
	var token string
	var output []string

	for i := 0; i < len(items); i++ {
		token = items[i]

		if GetType(token[0]) == Character || GetType(token[0]) == Number {
			output = append(output, token)
		} else if token[0] == '(' {
			stack.Push(token)
		} else if token[0] == ')' {
			for stack.Peek() != "(" {
				output = append(output, stack.Pop())
			}
			stack.Pop() //popping the left parenthesis from the stack
		} else if GetType(token[0]) == Operator {

			for stack.IsValid() && stack.Peek() != "" && (GetType(stack.Peek()[0]) == Operator && !OperatorPrecedence(token, stack.Peek())) {
				output = append(output, stack.Pop())
			}

			stack.Push(token)
		}
	}
	for stack.IsValid() {

		if stack.Peek() == ")" || stack.Peek() == "(" {
			panic("Error: Mismatched parenthesis in expression (ShuntingYard)")
		}
		output = append(output, stack.Pop())
	}

	return Trim(output) //What a band-aid solution...
}

/**
 * Takes a string expression and divides it up into numbers, operators, and variable
 * tokens.
 *
 * @expression: The string expression to be tokenized
 * @[]string: The string tokenized on numbers, operators, and variables
 */
func ParseExpression(expression string) []string {
	var thinglist []string

	expression = spaceremover(expression)

	for i := 0; i < len(expression); {

		var lettertype TokenType = GetType(expression[i])

		if expression[i] == '(' || expression[i] == ')' {
			thinglist = append(thinglist, string(expression[i]))
			i++
		} else if lettertype == Number {
			tempstring, end := CollectBytes(expression, i, Number)
			thinglist = append(thinglist, tempstring)
			i = end
		} else if lettertype == Character {
			tempstring, end := CollectBytes(expression, i, Character)
			thinglist = append(thinglist, tempstring)
			i = end
		} else if lettertype == Operator {
			if (i != len(expression)-1) &&
				(expression[i] == '<' || expression[i] == '>' || expression[i] == '=' || expression[i] == '!') &&
				(expression[i+1] == '=') {

				thinglist = append(thinglist, string(expression[i])+string(expression[i+1]))
				i += 2

			} else {
				thinglist = append(thinglist, string(expression[i]))
				i++
			}

			/*tempstring, end := CollectBytes(expression, i, 'o')
			thinglist = append(thinglist, tempstring)
			i = end*/
		} else {
			panic("Something is wrong with the expression at someplace (ParseExpression)") //finish this panic call
		}
	}

	return thinglist
}

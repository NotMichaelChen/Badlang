package ExpressionEvaluator

import (
	"badlang/Global"
	"math"
	"strconv"
)

//I need to rethink the public/private functions idea. Should they still be seperate,
//or should i just throw them all into public

/*******************************************
 *
 * 	         Public Functions
 *
 ******************************************/

/**
 * Given a certain token type, pull that type of data from the string
 *
 * @expression: The string to pull data from
 * @start: Where to begin pulling data from
 * @typefind: The specified token type (c, n, o)
 *
 * @string: The data that was pulled from the string
 * @int: The location of the end of the data
 */
func CollectBytes(expression string, start int, typefind TokenType) (string, int) { //collection, end
	var tempstring string
	var end int

	if typefind == Number {
		for i := start; i < len(expression) && (48 <= expression[i] && expression[i] <= 57); i++ {
			tempstring += string(expression[i])
			end = i
		}
	} else if typefind == Character {
		for i := start; i < len(expression) && (65 <= expression[i] && expression[i] <= 90) || (97 <= expression[i] && expression[i] <= 122); i++ {
			tempstring += string(expression[i])
			end = i
		}
	} else if typefind == Operator {
		for i := start; i < len(expression) && isoperator(expression[i]); i++ {
			tempstring += string(expression[i])
			end = i
		}
	}

	return tempstring, end + 1
}

/**
 * Determines whether the testing character is representative of a variable, a number,
 * or an operator
 *
 * @tester: The character being tested
 *
 * @TokenType: What kind of token the character is
 */
func GetType(tester byte) TokenType {
	if '0' <= tester && tester <= '9' {
		return Number //number
	} else if (65 <= tester && tester <= 90) || (97 <= tester && tester <= 122) {
		return Character //character (variable)
	} else if isoperator(tester) {
		return Operator //operator
	} else {
		return Err //something else
	}
}

/**
 * Determines which operator has a higher precedence
 * If o1 has a higher precedence, then return True
 * Otherwise if o2 has the higher precedence return False
 * Think of it as asking the question "is o1 a higher precedence than o2?"
 *
 * @o1: The first operator to test
 * @02: The second operator to test against
 *
 * @bool: Whether the first operator has a higher precedence or not
 */
func OperatorPrecedence(o1 string, o2 string) bool {
	var (
		num1 = 0
		num2 = 0
	)

	if o1 == "!" {
		num1 = 10
	} else if o1 == "/" || o1 == "*" || o1 == "%" {
		num1 = 9
	} else if o1 == "+" || o1 == "-" {
		num1 = 8
	} else if o1 == "^" {
		num1 = 7
	} else if o1 == ">" || o1 == "<" || o1 == "<=" || o1 == ">=" {
		num1 = 6
	} else if o1 == "==" || o1 == "!=" {
		num1 = 5
	} else if o1 == "&" {
		num1 = 4
	} else if o1 == "`" {
		num1 = 3
	} else if o1 == "|" {
		num1 = 2
	}

	if o2 == "!" {
		num2 = 10
	} else if o2 == "/" || o2 == "*" || o2 == "%" {
		num2 = 9
	} else if o2 == "+" || o2 == "-" {
		num2 = 8
	} else if o2 == "^" {
		num2 = 7
	} else if o2 == ">" || o2 == "<" || o2 == "<=" || o2 == ">=" {
		num2 = 6
	} else if o2 == "==" || o2 == "!=" {
		num2 = 5
	} else if o2 == "&" {
		num2 = 4
	} else if o2 == "`" {
		num2 = 3
	} else if o2 == "|" {
		num2 = 2
	}

	return num1 > num2
}

/**
 * Given an array of strings, trims off any empty entries in the front or back
 * I'm using this to code poorly ^^
 *
 * @items: The array of strings to look at
 * @[]string: The new array of strings with empty entries in front/back removed
 */
func Trim(items []string) []string {
	var (
		begin = 0
		end   = len(items)
	)

	if items[0] == "" {
		begin = 1
	} else if items[len(items)-1] == "" {
		end = len(items) - 1
		return items[0 : len(items)-1]
	}

	return items[begin:end]
}

/*******************************************
 *
 * 	         Private Functions
 *
 ******************************************/

//Does basic calculations for us
//We test if the variable exists in this function
//We also round the number in this function
func basiccalc(operator string, item1 string, item2 string) string {
	var num1, num2 float64
	//Just a holder variable, doesn't really do anything
	var asdf error

	//are item1 or item2 variables?
	//Checking if the first char in item1/item2 are letters; then we can tell if they are variables or not
	//In the end, fills num1/num2 with numbers to calculate
	if (65 <= item1[0] && item1[0] <= 90) || (97 <= item1[0] && item1[0] <= 122) {
		num1 = Global.Variables[item1] //Not afraid to throw an error if the var doesn't exist: maybe a better idea to panic instead?
	} else {
		num1, asdf = strconv.ParseFloat(item1, 64)
	}
	if (65 <= item2[0] && item2[0] <= 90) || (97 <= item2[0] && item2[0] <= 122) {
		num2 = Global.Variables[item2]
	} else {
		num2, asdf = strconv.ParseFloat(item2, 64)
	}

	if asdf != nil {
		panic("Something went wrong with converting the floats! (basiccalc)")
	}

	//Nums are rounded here
	num1 = roundprec(num1, 6)
	num2 = roundprec(num2, 6)

	var result float64

	//Special and simple case
	if operator == "!" {
		if item1 == "0" {
			return "1"
		} else {
			return "0"
		}

	} else if operator == "/" {
		result = num1 / num1
	} else if operator == "*" {
		result = num1 * num2
	} else if operator == "%" {
		result = floatmodulus(num1, num2)
	} else if operator == "+" {
		result = num1 + num2
	} else if operator == "-" {
		result = num1 - num2
	} else if operator == "^" {
		result = math.Pow(num1, num2)
	} else if operator == ">" {
		if num1 > num2 {
			result = 1
		} else {
			result = 0
		}
	} else if operator == "<" {
		if num1 < num2 {
			result = 1
		} else {
			result = 0
		}
	} else if operator == "<=" {
		if num1 <= num2 {
			result = 1
		} else {
			result = 0
		}
	} else if operator == ">=" {
		if num1 >= num2 {
			result = 1
		} else {
			result = 0
		}
	} else if operator == "==" {
		if num1 == num2 {
			result = 1
		} else {
			result = 0
		}
	} else if operator == "!=" {
		if num1 != num2 {
			result = 1
		} else {
			result = 0
		}
	} else if operator == "&" {
		if num1 != 0 && num2 != 0 {
			result = 1
		} else {
			result = 0
		}
	} else if operator == "`" {
		if (num1 != 0 && num2 != 0) || (num1 == 0 && num2 == 0) {
			result = 1
		} else {
			result = 0
		}
	} else if operator == "|" {
		if num1 != 0 || num2 != 0 {
			result = 1
		} else {
			result = 0
		}
	} else {
		panic("Operator not found in expression (RPN)")
	}

	//re-round the answer
	result = roundprec(result, 6)

	return strconv.FormatFloat(result, 'g', -1, 64)
}

func floatmodulus(num1 float64, num2 float64) float64 {
	if num1 != math.Ceil(num1) || num2 != math.Ceil(num2) {
		panic("Trying to do modulus operator on floats (floatmodulus)")
	}

	var temp1, temp2 int

	temp1 = int(num1)
	temp2 = int(num2)

	result := temp1 % temp2

	return float64(result)
}

func isoperator(tester byte) bool {
	switch tester {
	case '=':
		fallthrough
	case '<':
		fallthrough
	case '>':
		fallthrough
	case '!':
		fallthrough
	case '&':
		fallthrough
	case '|':
		fallthrough
	case '%':
		fallthrough
	case '(':
		fallthrough
	case ')':
		fallthrough
	case '*':
		fallthrough
	case '+':
		fallthrough
	case '-':
		fallthrough
	case '/':
		return true
	default:
		return false
	}
}

func roundprec(x float64, prec int) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	sign := 1.0
	if x < 0 {
		sign = -1
		x *= -1
	}

	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow * sign
}

func spaceremover(sentence string) string {
	var tempstring string

	for i := 0; i < len(sentence); i++ {
		if sentence[i] != ' ' {
			tempstring += string(sentence[i])
		}
	}

	return tempstring
}

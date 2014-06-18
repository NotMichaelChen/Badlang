// Hello project Hello.go
package main

import "fmt"

import (
	"badlang/ExpressionEvaluator"
)

/* TODO
 * Organize functions in Helperfuncs and document them

 *
 */

func main() {
	fmt.Println("Hello World!")

	fmt.Println(ExpressionEvaluator.ParseExpression("1+1"))

}

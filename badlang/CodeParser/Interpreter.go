package CodeParser

import (
	"badlang/ExpressionEvaluator" //Used for the gettype and collectbytes functions
	"badlang/Global"
	"strings"
)

//Based off the assumption that each line of code is either
//an expression or a command
func ExecutionLoop(codelist []string) {

	for _, line := range codelist {

		lineiscommand, command := isCommand(line)

		if lineiscommand {

		} else { //not a command, must be expression
			isequalexpression, _ := hasEquals(line)
			if isequalexpression {
				//runEquals()
			} //If there's no equals, then it's just an expression sitting there
			//evaluating it doesn't do us any good, so we just ignore it
			//However, the opportunity to do something is still there
		}
	}
}

func isCommand(line string) (bool, string) {

	//list of available commands
	commands := [4]string{"if", "while", "end"}
	//list of commands that should make the interpreter blow up
	invalcommands := [10]string{"then", "else"}

	var word string

	//checks if the start of the line is a char, that means there's a chance of
	//the line being a command
	if ExpressionEvaluator.GetType(line[0]) == ExpressionEvaluator.Character {
		word, _ = ExpressionEvaluator.CollectBytes(line, 0, ExpressionEvaluator.Character)

		//checks if the obtained word is a valid command
		for _, com := range commands {
			if strings.ToLower(word) == com {
				return true, word
			}
		}

		//makes sure the obtained word isn't an invalid command
		for _, invcom := range invalcommands {
			if strings.ToLower(word) == invcom {
				panic("Invalid command in code (isCommand)")
			}
		}

		return false, word
	} else {
		return false, ""
	}
}

//@return: where to continue execution of code from
func runIf(codelist []string, begin int) int {

	var execute bool

	expression := codelist[begin][2:len(codelist[begin])]
	expression = ExpressionEvaluator.Eval(expression)

}

//use this to figure out if there is an assignment we have to be aware of
//makes sure it's not a comparison equals, also makes sure there's only one
func hasEquals(sentence string) (bool, int) {
	hasequal := false
	var location int

	for i := 0; i < len(sentence); {

		//skip if ==
		if sentence[i] == '=' && sentence[i+1] == '=' {
			i += 2

			//checks if there's an actual = in the expression
		} else if sentence[i] == '=' && i < len(sentence)-1 && sentence[i+1] != '=' {

			//check to make sure there's only one
			if hasequal == true {
				panic("More than one equal in expression (runIf)")
			} else {
				hasequal = true
				location = i
			}
			i++
		} else {
			i++
		}
	}

	return hasequal, location
}

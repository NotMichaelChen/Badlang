package Structures

/*
 * [0,0,0,0]<-
 *          ->
 */

/**
 * @List: the array holding all of the data in the stack
 * @isvalid: Determines whether the stack is ok to pop/peek at
 */
type Stack struct {
	//Actual Stack
	List []string
	//Is the stack valid/is ok to pop?
	isvalid bool
}

/**
 * Pushes a number onto the stack. Should not matter whether the stack is valid
 * or not
 *
 * @num: The number to push onto the stack
 */
func (s *Stack) Push(num string) {
	//pushing something into an invalid stack makes it valid again
	if !s.isvalid {
		s.isvalid = true
	}
	s.List = append(s.List, num)
}

/**
 * Pops the number off the top of the stack.
 *
 * @string: The number that got popped off the stack
 */
func (s *Stack) Pop() string {
	var num string

	if len(s.List) == 1 {
		num = s.List[0]
		//because the last value doesn't disappear, we never have to check for an empty slice
		//it's a feature not a bug kay?
		s.List[0] = ""

		if s.isvalid {
			s.isvalid = false
		} else {
			panic("Error: trying to pop from empty stack")
		}

	} else {
		num = s.List[len(s.List)-1]
		s.List = s.List[:len(s.List)-1]
	}
	return num
}

/**
 * Looks at the top of the stack without removing the number
 *
 * @string: The number at the top of the stack
 */
func (s *Stack) Peek() string {
	if !s.isvalid {
		panic("Error: trying to peek from empty stack")
	}

	return s.List[len(s.List)-1]
}

/**
 * Checks whether the stack is valid to pop/peek off of
 *
 * @bool: Whether the stack is valid or not
 */
func (s *Stack) IsValid() bool {
	return s.isvalid
}

/**
 * Returns the size of the stack. This isn't an actual feature of stacks, but is
 * included for convenience sake
 *
 * @int: The size of the stack
 */
func (s *Stack) Size() int {
	return len(s.List)
}

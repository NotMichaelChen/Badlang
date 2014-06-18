package Structures

//<-[0,0,0,0]<-
/**
 * So right now the queue uses floats while the stack uses strings. This is because
 * stack is used in more cases than queue, and needs to be more flexible. While Queue
 * could do with using strings, there's no real reason to do so right now, so you
 * could consider this legacy code of some sort
 */

/**
 * @List: the array holding the data
 * @isvalid: tells whether the Queue is valid to peek/remove from
 */
type Queue struct {
	List    []float64
	isvalid bool
}

/**
 * Inserts a number into the Queue. This action should be possible regardless if
 * the queue is valid or not
 *
 * @num: The number to insert into the queue
 */
func (q *Queue) Insert(num float64) {
	if !q.isvalid {
		q.isvalid = true
	}
	q.List = append(q.List, num)
}

/**
 * Removes the first number from the queue, then shifts the rest of the array over
 *
 * @float64: The number that got removed from the queue
 */
func (q *Queue) Remove() float64 {
	num := q.List[0]

	if len(q.List) == 1 {
		//same deal here as in stack
		q.List[0] = 0
		if q.isvalid {
			q.isvalid = false
		} else {
			panic("Error: trying to remove from empty queue")
		}
	} else {
		//This effectively shifts the array over
		q.List = append(q.List[1:len(q.List)-1], q.List[len(q.List)-1])
	}

	return num
}

/**
 * Returns whether the queue is valid to remove from or not
 *
 * @bool: Whether the queue is valid or not
 */
func (q *Queue) IsValid() bool {
	return q.isvalid
}

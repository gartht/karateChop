package karatechop

func Chop(searchVal int, input []int) (current int) {

	//guard clause
	if len(input) == 0 {
		return -1
	}

	top := len(input) - 1
	bottom := 0
	current = len(input) / 2
	done := false

	for done == false {

		if input[current] == searchVal {
			return
		}

		if top == bottom {
			//indices have collapsed - the searchvalue is either at that index or it's not in the slice
			current = collapsedIndices(top, searchVal, input)
			return
		}

		if top-bottom == 1 {
			//indices are adjacent - the searchval is either in one of these or not in any
			current = adjacentIndices(top, bottom, searchVal, input)
			return
		}

		if input[current] > searchVal {
			top = current
		} else {
			bottom = current
		}
		current = (top + bottom) / 2
	}
	return
}

func collapsedIndices(i, searchVal int, input []int) (index int) {
	index = i
	if input[i] != searchVal {
		index = -1
	}
	return
}

func adjacentIndices(top, bottom, searchVal int, input []int) (index int) {
	if input[top] == searchVal {
		index = top
	} else {
		if input[bottom] == searchVal {
			index = bottom
		} else {
			index = -1
		}
	}
	return
}

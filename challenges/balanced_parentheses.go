package challenges

import "github.com/chucheka/data-structure-algorithm/stacks"

// BalancedParentheses
// This can also be solved using two pointer pattern
func BalancedParentheses(str string) bool {

	symbolsMap := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	stack := &stacks.LlStack{}

	for _, char := range str {

		if _, contains := symbolsMap[char]; contains {
			stack.Push(char)
		} else {

			last, err := stack.Pop()

			if err != nil {
				return false
			}

			c, _ := last.(rune)

			if closing := symbolsMap[c]; closing != char {
				return false
			}
		}

	}

	return true

}

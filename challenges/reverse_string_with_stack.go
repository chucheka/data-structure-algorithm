package challenges

import (
	"github.com/chucheka/data-structure-algorithm/stacks"
)

func ReverseString(str string) (string, error) {

	st := &stacks.LlStack{}

	for _, char := range str {
		st.Push(string(char))
	}

	result := ""

	for st.Size() > 0 {
		char, err := st.Pop()
		if err != nil {
			return "", err
		}

		s := char.(string)

		result += s
	}

	return result, nil
}

package patterns

func SameSquared(nums []int, squares []int) bool {

	lookUp := map[int]int{}

	if len(nums) != len(squares) {
		return false
	}

	for _, num := range nums {

		square := num * num

		if count, found := lookUp[square]; found {

			count += 1
			lookUp[square] = count
			continue
		}

		lookUp[square] = 1
	}

	for _, square := range squares {

		if count, found := lookUp[square]; found {

			count -= 1

			if count == 0 {
				delete(lookUp, square)
			} else {
				lookUp[square] = count
			}

			continue

		}

		return false

	}

	return true

}

func ValidAnagram(str1 string, str2 string) bool {

	runes1 := []rune(str1)

	runes2 := []rune(str2)

	if len(runes1) != len(runes2) {
		return false
	}

	lookUp := map[rune]int{}

	for _, char := range runes1 {

		if count, found := lookUp[char]; found {

			count += 1
			lookUp[char] = count

		} else {
			lookUp[char] = 1
		}

	}

	for _, char := range runes2 {

		if count, found := lookUp[char]; found {

			count -= 1

			if count == 0 {
				delete(lookUp, char)
			} else {
				lookUp[char] = count
			}

			continue

		}

		return false

	}

	return true
}

package sorting

func findSmallest(nums []int) int {

	smallest := nums[0]

	smallestIndex := 0

	for i, num := range nums {
		if num < smallest {
			smallest = num
			smallestIndex = i
		}
	}

	return smallestIndex
}

func SelectionSort(nums []int) int {

	newArr := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		smallestIndex := findSmallest(nums)
		newArr[i] = nums[smallestIndex]
		nums = append(nums[:smallestIndex], nums[smallestIndex+1])
	}

	return newArr[0]
}

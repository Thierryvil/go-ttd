package array_slice

func Sum(nums []int) int {
	var total = 0
	for _, num := range nums {
		total += num
	}
	return total
}

func SumAll(nums ...[]int) []int {
	var sumAll []int

	for _, num := range nums {
		sumAll = append(sumAll, Sum(num))
	}

	return sumAll
}

func SumAllRest(nums ...[]int) []int {
	var sumAll []int
	for _, num := range nums {
		if len(num) == 0 {
			sumAll = append(sumAll, 0)
		} else {
			final := num[1:]
			sumAll = append(sumAll, Sum(final))
		}
	}
	return sumAll
}

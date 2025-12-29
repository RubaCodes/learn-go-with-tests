package slicesranges

func Sum(numbers []int) int {
	var sum int
	for _, e := range numbers {
		sum += e
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	summed := make([]int, lengthOfNumbers)
	// var summed []int

	for j, v := range numbersToSum {
		summed[j] = Sum(v)
		// summed = append(summed,Sum(v))
	}

	return summed
}
func SumAllTails(numbersToSum ...[]int) []int {
	var summed []int
	for _, v := range numbersToSum {
		//slice[low:high] => [1:] "1 to the end"
		if len(v) > 1 {
			summed = append(summed, Sum(v[1:]))
		}else{
			summed = append(summed, 0)
		}
	}
	return summed
}

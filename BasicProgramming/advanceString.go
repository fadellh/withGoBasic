package BasicProgramming

import "strings"

func Compare(a, b string) string {
	// your code here
	if len(a) < len(b) {
		res := strings.Contains(b, a)
		if res {
			return a
		}
	}

	res := strings.Contains(a, b)

	if res {
		return b
	}
	return " "
}

func caesar(offset int, input string) string {
	result := ""
	for _, item := range input {
		if offset > 26 {
			offset = offset % 26
		}
		asci := item + rune(offset)
		if asci > 122 {
			// fmt.Println(asci - 122)
			asci = 96 + (asci - 122)
		}
		result += string(asci)
	}
	return result
}

func ArrayUnique(arrayA, arrayB []int) []int {
	// your code here
	collection := map[int]int{}
	output := []int{}

	for _, item := range arrayB {
		collection[item] = 1
	}

	for _, itemA := range arrayA {
		_, exist := collection[itemA]

		if !exist {
			output = append(output, itemA)
		}
	}

	return output
}

func findMaxSumSubArray(k int, arr []int) int {
	// your code here
	maxVal := 0
	firstIdx := 0
	lastIdx := k - 1

	for lastIdx < len(arr)-1 {
		tempSum := sum(arr[firstIdx : lastIdx+1]...)
		// fmt.Println(tempSum)
		if maxVal < tempSum {
			maxVal = tempSum
		}
		lastIdx++
		firstIdx++
	}

	return maxVal
}

func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func RemoveDuplicates(array []int) int {
	// your code here
	output := 1
	for i := 1; i < len(array); i++ {
		if array[i] != array[i-1] {
			output++
			continue
		}
	}
	return output
}

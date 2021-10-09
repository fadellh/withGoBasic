package BasicProgramming

import "strconv"

func primeNumber(input int) bool {
	if input == 1 {
		return false
	}

	if input == 2 {
		return true
	}
	factor := 0
	if input%1 == 0 && input%input == 0 {
		factor += 2
	}

	for i := 2; i < 10; i++ {
		if input%i == 0 && i != input {
			return false
		}
	}
	return true
}

func pow(x, n int) int {
	if n == 1 {
		return x
	}
	return x * pow(x, n-1)
}

// func pow2(x, n int) int {
// 	if n == 1 {
// 		return x
// 	}
// 	num := int(math.Floor(float64(n / 2)))
// 	return pow2(x, num) * pow2(x, num)
// }

func JoinArrayRemoveDuplicate(arrayA, arrayB []string) []string {

	joinArray := append(arrayA, arrayB...)
	fruit := map[string]int{}

	for _, item := range joinArray {
		_, exist := fruit[string(item)]

		if !exist {
			fruit[string(item)] = 1
		}
	}

	joinArray = []string{}
	for key := range fruit {
		joinArray = append(joinArray, key)
	}

	return joinArray
}

func munculSekali(angka string) []int {
	// your code here
	var numberData = map[string]int{}
	output := []int{}

	for _, v := range angka {
		value, exist := numberData[string(v)]
		if exist {
			numberData[string(v)] = value + 1
		} else {
			numberData[string(v)] = 1
		}
	}

	for key, value := range numberData {
		if value == 1 {
			str, _ := strconv.Atoi(key)
			output = append(output, str)
		}
	}
	return output
}

func PairSum(arr []int, target int) []int {
	// your code here
	//1, 2, 3, 4, 6}, 6
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]
		// fmt.Println(sum)

		if sum > target {
			right--
			continue
		}

		if sum < target {
			left++
			continue
		}
		// return []int{arr[left], arr[right]}
		return []int{left, right}
	}
	return []int{}
}

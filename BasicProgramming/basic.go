package BasicProgramming

import (
	"fmt"
	"strconv"
)

func test() {
	prime := 17
	fmt.Println(prime % 1)
}

func FaktorBilangan(input int) {
	for i := 1; i <= input; i++ {
		if input%i == 0 {
			fmt.Println(i)
		}
	}
}

//primeNumber determine prime Number
func PrimeNumber(input int) bool {
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

func Palindrome(input string) bool {

	front := 0
	back := len(input) - 1

	for front <= back {
		if string(input[front]) != string(input[back]) {
			return false
		}
		front++
		back--
	}

	return true
	//return handlePalindrome(input, 0, len-1)
}

func Pangkat(base, pangkat int) int {
	result := 1
	for i := pangkat; i > 0; i-- {
		result *= base
	}
	return result
}

func FullPrima(N int) bool {
	if PrimeNumber(N) {
		numberStr := strconv.Itoa(N)
		for _, element := range numberStr {
			char := string(element)
			var num, _ = strconv.Atoi(char)
			if !PrimeNumber(num) {

				return false
			}
		}
		return true
	}
	return false
}

func HandlePalindrome(input string, frontIndex, backIndex int) bool {
	if frontIndex == backIndex {
		return true
	}

	if string(input[backIndex]) != string(input[frontIndex]) {
		return false
	}

	if frontIndex < backIndex+1 {
		return HandlePalindrome(input, frontIndex+1, backIndex+1)
	}

	return true
}

func PlayWithAsterik(n int) string {
	var result string

	//wide := 2*(n-1) + n
	high := n
	pointer := 0
	leftBlank := n - 1
	midleStar := 1
	midleStar2 := 0
	rightBlank := n - 1
	//first triangle
	for pointer < high {
		for i := leftBlank; i > 0; i-- {
			result += " "
		}

		for i := midleStar; i > 0; i-- {
			result += "*"
		}

		for i := midleStar2; i > 0; i-- {
			result += "*"
		}

		for i := rightBlank; i > 0; i-- {
			result += " "
		}

		leftBlank--
		midleStar++
		rightBlank--
		midleStar2++
		result += "\n"
		pointer++
	}

	return result
}

func CetakTablePerkalian(number int) {
	pointer := 1
	for pointer <= number {
		var result string
		for i := 1; i <= number; i++ {
			result += strconv.Itoa(pointer * i)
			result += " "

		}
		fmt.Println(result)
		result += "\n"
		pointer++
	}
}

func UbahHuruf(sentence string) string {

	var result string

	for _, v := range sentence {
		encrypt := v + 10
		asci := encrypt
		if encrypt > 90 {
			asci = 64 + (encrypt - 90)
		}

		if string(v) == " " {
			result += " "
		} else {
			result += string(asci)
		}
	}
	return result
}

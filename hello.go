package main

import "fmt"

// type Student struct {
// 	name  []string
// 	score []int
// }

// func (s Student) Average() float64 {
// 	var total float64
// 	for _, v := range s.score {
// 		total += float64(v)
// 	}

// 	return total / float64(len(s.score))
// }

// func (s Student) Min() (min int, name string) {

// 	for idx, v := range s.score {
// 		if idx == 0 {
// 			min = v
// 			name = s.name[idx]
// 		}
// 		if v < min {
// 			min = v
// 			name = s.name[idx]
// 		}
// 	}

// 	return min, name
// }

// func (s Student) Max() (max int, name string) {

// 	for idx, v := range s.score {
// 		if v > max {
// 			max = v
// 			name = s.name[idx]
// 		}
// 	}

// 	return max, name
// }

type students struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *students) Encode() string {
	var nameEncode string
	offset := 3
	for _, item := range s.name {
		if offset > 26 {
			offset = offset % 26
		}
		asci := item + rune(offset)
		if asci > 122 {
			// fmt.Println(asci - 122)
			asci = 96 + (asci - 122)
		}
		nameEncode += string(asci)
	}
	return nameEncode
}

func (s *students) Decode() string {
	var nameDecode = ""
	offset := 23
	for _, item := range s.name {
		if offset > 26 {
			offset = offset % 26
		}
		asci := item + rune(offset)
		if asci > 122 {
			// fmt.Println(asci - 122)
			asci = 96 + (asci - 122)
		}
		nameDecode += string(asci)
	}
	return nameDecode
}

func main() {
	// a := 10
	// b := 20

	// swap(&a, &b)
	// var name string = "John"
	// var nameAddress *string = &name

	// fmt.Println("name (value) :", name)
	// fmt.Println("name (address) :", &name)
	// fmt.Println("nameAddress (value) :", *nameAddress)
	// fmt.Println("nameAddress (address) :", nameAddress)

	// name = "Doe"

	// fmt.Println("name (value) :", name)
	// fmt.Println("name (address) :", &name)
	// fmt.Println("nameAddress (value) :", *nameAddress)
	// fmt.Println("nameAddress (address) :", nameAddress)
	// var a1, a2, a3, a4, a5, a6, min, max int
	// fmt.Scan(&a1)
	// fmt.Scan(&a2)
	// fmt.Scan(&a3)
	// fmt.Scan(&a4)
	// fmt.Scan(&a5)
	// fmt.Scan(&a6)

	// min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	// fmt.Println("Nilai Minimum: ", min)
	// fmt.Println("Nilai Maksimum: ", max)

	// var a = Student{}

	// for i := 0; i < 3; i++ {
	// 	var name string
	// 	fmt.Println("Input " + string(rune(i)) + " Student's Name : ")
	// 	fmt.Scan(&name)
	// 	a.name = append(a.name, name)
	// 	var score int
	// 	fmt.Print("Input " + name + " Score : ")
	// 	fmt.Scan(&score)
	// 	a.score = append(a.score, score)
	// }

	// fmt.Println("\n\nAverage Score Students is", a.Average())
	// scoreMax, nameMax := a.Max()
	// fmt.Println("Max Score Student is : "+nameMax+"(", scoreMax, ")")
	// scoreMin, nameMin := a.Min()
	// fmt.Println("Min Score Student is : "+nameMin+"(", scoreMin, ")")

	var menu int
	var a = students{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name" + a.name + " is " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of Student's Name" + a.name + " is " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}

}

func swap(a, b *int) {
	a = b
	b = a
}

func getMinMax(numbers ...*int) (min int, max int) {

	for idx, item := range numbers {
		if idx == 0 {
			min = *item
		}
		if *item < min {
			min = *item
		}
		if *item > max {
			max = *item
		}
	}
	return min, max
}

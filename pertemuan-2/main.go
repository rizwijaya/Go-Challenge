package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		for k := j; k < 5; k++ {
	// 			fmt.Print(j, " ")
	// 		}
	// 	}
	// }

	// fmt.Println()
	fruit := []string{"apple", "banana", "grape", "orange", "mango"}
	for i, v := range fruit {
		println(i, v)
	}
	fmt.Println(fruit[10])
}

// package main

// import "fmt"

// func main() {
// 	char := []rune{'\u0421', '\n', '\u0410', '\n', '\u0428', '\n', '\u0410', '\n', '\u0420', '\n', '\u0412', '\n', '\u041E', '\n'}
// 	var result_1, result_2, russian []string
// 	var lima bool = true
// 	i := 0
// 	for i = 0; i < len(char); i++ {
// 		if i <= 4 {
// 			result_1 = append(result_1, fmt.Sprintf("Nilai i = %d", i))
// 		}
// 		if i%2 == 0 && lima {
// 			russian = append(russian, fmt.Sprintf("character %U '%c' starts at byte position %d", char[i], char[i], i))
// 		}

// 		if i+1 == len(char) {
// 			lima = false
// 			i = 6
// 		}
// 		if i > 5 && !lima {
// 			russian = append(russian, fmt.Sprintf("Nilai i = %d", i))
// 		}
// 		if i >= 10 && !lima {
// 			break
// 		}
// 		result_2 = append(result_2, fmt.Sprintf("Nilai j = %d", i))
// 	}
// 	result := append(result_1, result_2...)
// 	result = append(result, russian...)

// 	for _, v := range result {
// 		fmt.Println(v)
// 	}
// }


package main

import "fmt"

func main() {
	var result, nums, russian []string
	char := []rune{'\u0421', '\n', '\u0410', '\n', '\u0428', '\n', '\u0410', '\n', '\u0420', '\n', '\u0412', '\n', '\u041E', '\n'}
	for i := 0; i < len(char); i++ {
		if i%2 == 0 {
			russian = append(russian, fmt.Sprintf("character %U '%c' starts at byte position %d", char[i], char[i], i))
		}
		nums = append(nums, fmt.Sprint(i))
	}

	result = append(result, nums[0:5]...)
	result = append(result, nums[0:5]...)
	result = append(result, russian...)
	result = append(result, nums[6:11]...)
	//fmt.Println(result)
	for i, v := range result {
		if i < 5 {
			fmt.Println("Nilai i = ", v)
		} else if i >= 10 && i < 17 {
			fmt.Println(v)
		} else {
			fmt.Println("Nilai j = ", v)
		}
	}
}

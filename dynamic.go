package main

import "fmt"

var fizzMap = map[int]string{
	3: "Fizz",
	5: "Buzz",
}

func dynamic(max int) {
	for i := 1; i <= max; i++ {
		str := getWord(i)

		if str == "" {
			fmt.Println(i)
		} else {
			fmt.Println(str)
		}
	}
}

func getWord(i int) (str string) {
	for divisor, word := range fizzMap {
		str += divisibilityChecker(i, divisor, word)
	}
	return str
}

func divisibilityChecker(i, divisor int, output string) string {
	if i%divisor == 0 {
		return output
	}
	return ""
}

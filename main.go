package main

import (
	"fmt"
	"time"
)

func main() {
	const max = 1000

	m := funcTimer(masterString, max)
	d := funcTimer(dynamicMasterString, max)
	h := funcTimer(hardcoded, max)

	fmt.Println()
	fmt.Println(m)
	fmt.Println(d)
	fmt.Println(h)
}

func funcTimer(f func(int), max int) int64 {
	start := time.Now().UnixNano()
	f(max)
	end := time.Now().UnixNano()
	return totalTime(start, end)
}

func hardcoded(max int) {
	for i := 1; i <= max; i++ {
		if i%15 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		}
		fmt.Println(i)
	}
}

func masterString(max int) {
	for i := 1; i <= max; i++ {
		str := ""
		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		fmt.Print(str + " ")
		fmt.Println(i)
	}
}

func dynamicMasterString(max int) {
	divisibilityChecker := func(i, divisor int, output string) string {
		if i%divisor == 0 {
			return output
		}
		return ""
	}

	fizzMap := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	fizzMapLooper := func(i int) (str string) {
		for divisor, word := range fizzMap {
			str += divisibilityChecker(i, divisor, word)
		}
		return str
	}

	for i := 1; i <= max; i++ {
		str := fizzMapLooper(i)
		if str == "" {
			fmt.Println(i)
		} else {
			fmt.Println(str)
		}
	}
}

func totalTime(start int64, end int64) int64 {
	return end - start
}

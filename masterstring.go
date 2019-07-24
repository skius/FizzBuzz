package main

import "fmt"

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

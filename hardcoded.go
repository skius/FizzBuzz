package main

import "fmt"

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

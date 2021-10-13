package main

import (
	"fmt"
	"logic/leapyear"
	"logic/nearf"
	"logic/palindrome"
	"logic/reverse"
)

func main() {
	var p string = "Radar"
	palindrome.Palind(p)

	var awal int = 1900
	var akhir int = 2020
	leapyear.Leap(awal, akhir)

	var text string = "I am A Great Human"
	reverse.Revs(text)

	var arr = []int{15, 1, 6}
	nearf.Fibbo(arr)

	const limit int = 15
	var answer [limit + 1]string

	for i := 1; i <= limit; i++ {
		if i%5 == 0 && i%3 == 0 {
			answer[i] = "FizzBuzz"
		} else if i%5 == 0 {
			answer[i] = "Buzz"
		} else if i%3 == 0 {
			answer[i] = "Fizz"
		} else {
			answer[i] = fmt.Sprint(i)
		}

	}
	fmt.Println(answer)
}

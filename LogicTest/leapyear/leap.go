package leapyear

import "fmt"

func Leap(begin int, end int) {

	for i := begin; i <= end; i++ {
		if i%4 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println("")
}

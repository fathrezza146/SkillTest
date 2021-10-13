package palindrome

import (
	"fmt"
	"strings"
)

func Palind(A string) {
	var text string = strings.ToLower(A)
	var reversed string

	for i := len(A) - 1; i >= 0; i-- {
		reversed += string(text[i])
	}

	for i := 0; i < len(A); i++ {
		if i == len(A)-1 {
			fmt.Println("Palindrome")
			break
		}
		if text[i] == reversed[i] {
			continue
		} else {
			fmt.Println("Not Palindrome")
			break
		}
	}
}

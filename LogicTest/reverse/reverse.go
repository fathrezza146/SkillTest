package reverse

import (
	"fmt"
	"strings"
	"unicode"
)

func Revs(A string) {
	var temp string
	var rev string
	var counter int
	var result string

	for i := 0; i < len(A); i++ {
		if string(A[i]) != " " {
			temp += string(A[i])
			fmt.Println(temp)
		} else if string(A[i]) == " " {
			for j := i - counter - 1; j >= 0; j-- {
				rev += string(temp[j])
				if unicode.IsUpper(rune(temp[0])) {
					rev = strings.ToLower(rev)
					rev = strings.Title(rev)
				}
			}
			result += rev + " "
			counter = i + 1
			temp = ""
			rev = ""
		}
		if i == len(A)-1 {
			for j := i - counter; j >= 0; j-- {
				rev += string(temp[j])
			}
			if unicode.IsUpper(rune(temp[0])) {
				rev = strings.ToLower(rev)
				rev = strings.Title(rev)
			}
			result += rev + " "
		}

	}
	fmt.Print(result)
}

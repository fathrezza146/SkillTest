package nearf

import "fmt"

func Fibbo(arr []int) {
	var fibbonaci int
	var comparator int = 1
	var total int
	var lowcomparator int
	var up int
	var down int

	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	fmt.Println(total)
	for {
		lowcomparator = fibbonaci
		fibbonaci += comparator
		comparator = fibbonaci - comparator
		fmt.Print(fibbonaci, " ")
		if fibbonaci > total {
			up = fibbonaci - total
			down = total - lowcomparator
			if up < down {
				fmt.Println(up)
				break
			} else {
				fmt.Println(down)
				break
			}
		}
	}
}

package levelone

import (
	"fmt"
	"strconv"
)

func LargestSlideWindowProd(wlen int, number string) int {
	prod := 1
	zeros := 0
	maxProd := -1
	mpStart := 0
	mpEnd := 0
	winStart := 0
	winEnd := wlen - 1

	for l, n := range number {
		if l <= wlen-1 {
			num, _ := strconv.Atoi(string(n))
			if num == 0 {
				zeros++
			}
			if num != 0 {
				prod = prod * num
			}
			//fmt.Printf("Product %d number  %d l : %d\n", prod, num, l)
		} else {
			winStart = l - wlen
			winEnd = l
			sElement, _ := strconv.Atoi(string(number[winStart]))
			eElement, _ := strconv.Atoi(string(number[winEnd]))
			if sElement != 0 {
				prod = prod / sElement
				//fmt.Printf("removing element %d, prod : %d\n", sElement, prod)
			} else {
				zeros--
			}
			if eElement != 0 {
				prod = prod * eElement
				//fmt.Printf("adding element %d, prod : %d\n", eElement, prod)
			} else {
				zeros++
			}
			//fmt.Printf("Product %d number  %d l : %d\n", prod, eElement, l)
			if maxProd < prod && zeros == 0 {
				maxProd = prod
				mpStart = winStart
				mpEnd = winEnd
			}
		}
	}
	fmt.Printf("Max Product for window size %d is %d\n", wlen, maxProd)
	fmt.Printf("Max Product window is %s \n", string(number[mpStart+1:mpEnd+1]))
	return maxProd
}

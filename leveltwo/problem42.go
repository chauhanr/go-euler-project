package leveltwo

import "fmt"

var letterCount = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

func initializeTriangleSeries() map[int]int {
	tr := map[int]int{}

	for i := 1; i <= 20; i++ {
		term := i * (i + 1) / 2
		tr[term] = i
	}

	return tr
}

func FindTriangleTerms(terms []string) int {
	triangleTermCount := 0
	tset := initializeTriangleSeries()
	fmt.Printf("triangle Set %v\n", tset)
	for _, t := range terms {
		c := GetTermCount(t)
		if _, ok := tset[c]; ok {
			triangleTermCount++
		}
	}
	fmt.Printf("Triangle Term Count %d\n", triangleTermCount)
	return triangleTermCount
}

func GetTermCount(str string) int {
	trCount := 0
	for _, s := range str {
		l := letterCount[string(s)]
		trCount += l
	}
	//fmt.Printf("Term Count %s is %d\n", str, trCount)
	return trCount
}

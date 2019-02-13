package levelone

import "fmt"

func NumberToWords(n int) int {
	count := 0
	if n > 1000 {
		return -1
	}
	for i := 1; i <= n; i++ {
		_, c := Stringify(i)
		count += c
	}
	fmt.Printf("Count for %d is %d\n", n, count)
	return count
}

func charCount(n int, str string) int {
	count := 0
	for _, c := range str {
		if string(c) != " " {
			count++
		}
	}
	//fmt.Printf("word count for %d is %d\n", n, count)
	return count
}

func Stringify(n int) (string, int) {
	num := ""
	andbool := true
	if n%1000 == 0 {
		return "one thousand", charCount(n, "one thousand")
	}
	n = n % 1000
	if n >= 100 {
		p := n / 100
		v := StringifyTens(p)
		num = v + " hundred"
		andbool = false
	}
	n = n % 100
	if n < 100 && n >= 20 {
		if n%10 == 0 {
			v := StringifyTensMul(n)
			if num != "" {
				if !andbool {
					num = num + " and " + v
				} else {
					num = num + " " + v
				}
			} else {
				num += v
			}
		} else {
			h := (n - (n % 10))
			v := StringifyTensMul(h)
			if v != "" {
				if num != "" {
					num = num + " and " + v
					andbool = true
				} else {
					num += v
				}
			}
			v = StringifyTens(n % 10)
			if v != "" {
				if num != "" {
					if !andbool {
						num = num + " and" + v
					} else {
						num = num + " " + v
					}
				} else {
					num += v
				}
			}
		}
	} else {
		v := StringifyTens(n)
		if v != "" {
			if num != "" {
				if andbool {
					num = num + " " + v
				} else {
					num = num + " and " + v
				}
			} else {
				num += v
			}
		}
	}
	c := charCount(n, num)
	//fmt.Printf("Number: %s\n", num)
	return num, c

}

func StringifyTens(n int) string {
	if n < 20 {
		switch n {
		case 1:
			return "one"
		case 2:
			return "two"
		case 3:
			return "three"
		case 4:
			return "four"
		case 5:
			return "five"
		case 6:
			return "six"
		case 7:
			return "seven"
		case 8:
			return "eight"
		case 9:
			return "nine"
		case 10:
			return "ten"
		case 11:
			return "eleven"
		case 12:
			return "twelve"
		case 13:
			return "thirteen"
		case 14:
			return "fourteen"
		case 15:
			return "fifteen"
		case 16:
			return "sixteen"
		case 17:
			return "seventeen"
		case 18:
			return "eighteen"
		case 19:
			return "nineteen"
		}
	}
	return ""

}

func StringifyTensMul(n int) string {
	if n < 100 && n%10 == 0 {
		switch n {
		case 10:
			return "ten"
		case 20:
			return "twenty"
		case 30:
			return "thirty"
		case 40:
			return "forty"
		case 50:
			return "fifty"
		case 60:
			return "sixty"
		case 70:
			return "seventy"
		case 80:
			return "eighty"
		case 90:
			return "ninety"
		}
	}
	return ""
}

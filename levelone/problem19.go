package levelone

import (
	"fmt"
	"strconv"
)

func CountDaysInMonth(date int, days int) (int, int) {
	ds := strconv.Itoa(date) + ","
	count := 1
	for date+7 <= days {
		count++
		date += 7
		ds = ds + strconv.Itoa(date) + ","
	}
	nd := 0
	if date < days {
		l := days - date
		nd = 7 - l
	} else if date == days {
		nd = 7
	}
	//fmt.Printf("dates : %s\n", ds)
	return count, nd
}

var numDays = map[int]int{
	1:  31,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func CalculateCenturySundays(sd, year1, year2 int) int {
	// first calculate the first sunday in 1901
	total := 0
	year := year1 - 1
	_, nd := CountWeekDayInYear(sd, year)
	year++
	for year <= year2 {
		yc, nmd := CountWeekDayInYear(nd, year)
		nd = nmd
		year++
		total += yc
	}
	fmt.Printf("Count for weekday between %d and %d is %d\n", year1, year2, total)
	return total
}

func CountWeekDayInYear(date int, year int) (ycount int, nd int) {
	dm := 0
	ycount = 0
	month := 1
	nd = date
	for month <= 12 {
		if month != 2 {
			dm = numDays[month]

		} else {
			if isLeapYear(year) {
				dm = 29
			} else {
				dm = 28
			}
		}
		//fmt.Printf("----Month --- %d\n", month)
		if nd == 1 {
			ycount += 1
		}
		_, nmd := CountDaysInMonth(nd, dm)
		nd = nmd
		month++
		//fmt.Printf("Year %d Month %d, has %d sundays\n", year, month-1, c)
	}
	return ycount, nd
}

func isLeapYear(year int) bool {
	if year%100 == 0 {
		val := year / 100
		if val%4 == 0 {
			return true
		} else {
			return false
		}
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}

}

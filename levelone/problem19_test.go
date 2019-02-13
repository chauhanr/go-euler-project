package levelone

import "testing"

func TestWeekDayCountCentury(t *testing.T) {
	totalSundays := 171
	sun := CalculateCenturySundays(6, 1901, 2000)
	if sun != totalSundays {
		t.Errorf("Expected Sundays to be %d but got %d", totalSundays, sun)
	}
}

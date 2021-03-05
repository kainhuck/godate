package godate

import "fmt"

var (
	MonthArray         = [12]bool{true, false, true, false, true, false, true, true, false, true, false, true}
	WeekArray          = [7]string{"Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}
	LeapMonthDayArray  = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	ALeapMonthDayArray = [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

type Date struct {
	Year  int
	Month int
	Day   int
}

// this is th datetime when i create this project
var BirthDay = &Date{
	Year:  2021,
	Month: 3,
	Day:   5,
}

func NewDate(year int, month int, day int) *Date {
	Assert(month > 0 && month < 13, "The number of months is out of the range")
	Assert(day > 0 && day < 32, "The number of days is out of the range")
	return &Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

// return by string with format string or not
func (d *Date) String(format ...string) string {
	var formatString = "%d/%d/%d"
	if len(format) > 0 {
		formatString = format[0]
	}

	return fmt.Sprintf(formatString, d.Year, d.Month, d.Day)
}

// add one day
func (d *Date) Increase() {
	d.AddDay(1)
}

// sub one day
func (d *Date) Reduce() {
	d.SubDay(1)
}

func (d *Date) Copy(x *Date) {
	d.Year = x.Year
	d.Month = x.Month
	d.Day = x.Day
}

// add x days
func (d *Date) AddDay(x int) {
	Assert(x >= 0, "x must > 0")
	days := d.DaysOfYear() + x
	d.Month, d.Day = 1, 1
	for days > d.Days() {
		days -= d.Days()
		d.Year++
	}

	d.Copy(TurnDaysToDate(days, d.Year))
}

// sub one day
func (d *Date) SubDate(x *Date) int {
	if d.Year == x.Year {
		return d.DaysOfYear() - x.DaysOfYear()
	} else if d.Year > x.Year {
		total := 0

		for d.Year > x.Year {
			total += d.Days()
			d.Year--
		}

		return d.DaysOfYear() - x.DaysOfYear() + total
	} else {
		return -x.SubDate(d)
	}
}

// sub x days
func (d *Date) SubDay(x int) {
	Assert(x >= 0, "x must > 0")
	days := d.DaysOfYear() - x - 1
	d.Month, d.Day = 1, 1

	for days <= 0 {
		d.Year--
		days += d.Days()
	}

	d.AddDay(days)
}

// return What day is it today
func (d *Date) Week() string {
	// 2021/3/5 -> Friday
	days := d.SubDate(BirthDay)
	if days < 0 {
		days = -days
	}

	fmt.Println(days, days%7)
	return WeekArray[days%7]
}

// return if this year is leap year
func (d *Date) IsLeap() bool {
	return !(d.Year%400 == 0 || (d.Year%4 == 0 && d.Year%100 != 0))
}

// return the days of this year
func (d *Date) Days() int {
	if d.IsLeap() {
		return 365
	}
	return 366
}

// return the days of this year
func (d *Date) DaysOfYear() int {
	total := 0
	for i := 1; i < d.Month; i++ {
		if isBigMonth(i) {
			total += 31
		} else {
			if i == 2 && d.IsLeap() {
				total += 28
			} else if i == 2 && !d.IsLeap() {
				total += 29
			} else {
				total += 30
			}
		}
	}
	total += d.Day
	return total
}

// return if this is a right day
func (d *Date) Check() bool {
	if d.IsBigMonth() {
		return true
	}
	if d.Month != 2 {
		return d.Day < 31
	}
	if d.IsLeap() {
		return d.Day < 29
	}
	return d.Day < 30
}

// return if this month ids big month
func (d *Date) IsBigMonth() bool {
	return isBigMonth(d.Month)
}

func isBigMonth(m int) bool {
	return MonthArray[m-1]
}

//
func TurnDaysToDate(days int, year int) *Date {
	d := &Date{
		Year:  year,
		Month: 1,
		Day:   1,
	}
	monthDay := ALeapMonthDayArray
	if d.IsLeap() {
		Assert(days <= 365 && days > 0, "leap year is 365 days")
		monthDay = LeapMonthDayArray
	} else {
		Assert(days <= 366 && days > 0, "a leap year is 366 days")
	}

	for i := 1; i <= 12; i++ {
		if days > monthDay[i-1] {
			d.Month++
			days -= monthDay[i-1]
		} else {
			break
		}
	}

	d.Day = days

	return d
}

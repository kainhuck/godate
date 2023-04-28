package godate

import (
	"fmt"
	"time"
)

var (
	monthArray         = [12]bool{true, false, true, false, true, false, true, true, false, true, false, true}
	weekArray          = [7]string{Friday, Saturday, Sunday, Monday, Tuesday, Wednesday, Thursday}
	leapMonthDayArray  = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	aLeapMonthDayArray = [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

const (
	Monday    = "Monday"
	Tuesday   = "Tuesday"
	Wednesday = "Wednesday"
	Thursday  = "Thursday"
	Friday    = "Friday"
	Saturday  = "Saturday"
	Sunday    = "Sunday"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

// BirthDay this is th datetime when i create this project
var BirthDay = &Date{
	Year:  2021,
	Month: 3,
	Day:   5,
}

func NewDate(year int, month int, day int) (*Date, error) {
	if month <= 0 || month >= 13 {
		return nil, fmt.Errorf("the number of months is out of the range")
	}
	if day <= 0 || day >= 32 {
		return nil, fmt.Errorf("the number of days is out of the range")
	}

	return &Date{
		Year:  year,
		Month: month,
		Day:   day,
	}, nil
}

func NewDateFromStr(dateStr string, layout ...string) (*Date, error) {
	var (
		year  int
		month int
		day   int
	)

	var formatString = "%d-%02d-%02d"
	if len(layout) > 0 {
		formatString = layout[0]
	}

	if _, err := fmt.Sscanf(dateStr, formatString, &year, &month, &day); err != nil {
		return nil, err
	}

	return NewDate(year, month, day)
}

// String return by string with format string or not
func (d *Date) String(layout ...string) string {
	var formatString = "%d-%02d-%02d"
	if len(layout) > 0 {
		formatString = layout[0]
	}

	return fmt.Sprintf(formatString, d.Year, d.Month, d.Day)
}

// Increase add one day
func (d *Date) Increase() {
	d.AddDay(1)
}

// Decrease sub one day
func (d *Date) Decrease() {
	d.SubDay(1)
}

func (d *Date) copy(x *Date) {
	d.Year = x.Year
	d.Month = x.Month
	d.Day = x.Day
}

func (d *Date) clone() *Date {
	d1, _ := NewDate(d.Year, d.Month, d.Day)

	return d1
}

func (d *Date) Equal(x *Date) bool {
	return d.SubDate(x) == 0
}

func (d *Date) EarlyThan(x *Date) bool {
	return d.SubDate(x) < 0
}

func (d *Date) LaterThan(x *Date) bool {
	return d.SubDate(x) > 0
}

func (d *Date) EarlyThanOrEqual(x *Date) bool {
	return d.EarlyThan(x) && d.Equal(x)
}

func (d *Date) LaterThanOrEqual(x *Date) bool {
	return d.LaterThan(x) && d.Equal(x)
}

// AddDay add x days
func (d *Date) AddDay(x int) {
	assert(x >= 0, "x must > 0")
	days := d.DaysOfYear() + x
	d.Month, d.Day = 1, 1
	for days > d.Days() {
		days -= d.Days()
		d.Year++
	}

	d.copy(turnDaysToDate(days, d.Year))
}

// SubDate sub one day
func (d *Date) SubDate(x *Date) int {
	dd := d.clone()
	if dd.Year == x.Year {
		return dd.DaysOfYear() - x.DaysOfYear()
	} else if dd.Year > x.Year {
		total := 0

		for dd.Year > x.Year {
			total += dd.Days()
			dd.Year--
		}

		return dd.DaysOfYear() - x.DaysOfYear() + total
	} else {
		return -x.SubDate(d)
	}
}

// SubDay sub x days
func (d *Date) SubDay(x int) {
	assert(x >= 0, "x must > 0")
	days := d.DaysOfYear() - x - 1
	d.Month, d.Day = 1, 1

	for days <= 0 {
		d.Year--
		days += d.Days()
	}

	d.AddDay(days)
}

func (d *Date) AddWeek(x int) {
	d.AddDay(x * 7)
}

func (d *Date) SubWeek(x int) {
	d.SubDay(x * 7)
}

// Week return what day is it today
func (d *Date) Week() string {
	// 2021/3/5 -> Friday
	days := d.SubDate(BirthDay)

	result := days % 7
	if result < 0 {
		result += 7
	}

	return weekArray[result]
}

// IsLeap return true if this year is leap year
func (d *Date) IsLeap() bool {
	return d.Year%400 == 0 || (d.Year%4 == 0 && d.Year%100 != 0)
}

// Days return the days of this year
func (d *Date) Days() int {
	if d.IsLeap() {
		return 365
	}
	return 366
}

// DaysOfYear return the days of this year
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

func (d *Date) WeeksOfYear() int {
	days := d.DaysOfYear()
	if days%7 != 0 {
		return days/7 + 1
	}
	return days / 7
}

// Check return if this is a right day
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

// IsBigMonth return if this month ids big month
func (d *Date) IsBigMonth() bool {
	return isBigMonth(d.Month)
}

func isBigMonth(m int) bool {
	return monthArray[m-1]
}

func turnDaysToDate(days int, year int) *Date {
	d := &Date{
		Year:  year,
		Month: 1,
		Day:   1,
	}
	monthDay := aLeapMonthDayArray
	if d.IsLeap() {
		assert(days <= 365 && days > 0, "leap year is 365 days")
		monthDay = leapMonthDayArray
	} else {
		assert(days <= 366 && days > 0, "a leap year is 366 days")
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

func (d *Date) Accurate(x int) {
	assert(x > 0 && x <= 366, "x in wrong range")
	if d.IsLeap() && x == 366 {
		panic("x in wrong range")
	}
	d.copy(turnDaysToDate(x, d.Year))
}

func Today() *Date {
	t := time.Now()
	year, month, day := t.Date()
	d, _ := NewDate(year, int(month), day)

	return d
}

func (d *Date) IsToday() bool {
	return d.Equal(Today())
}

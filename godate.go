package godate

import (
	"fmt"
	"time"
)

var (
	weekArray = [7]string{Friday, Saturday, Sunday, Monday, Tuesday, Wednesday, Thursday}
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

// birthDay represents the date when I created this project.
var birthDay = &Date{
	Year:  2021,
	Month: 3,
	Day:   5,
}

// NewDate creates a Date instance from the specified year, month, and day. Returns an error if the date is invalid.
func NewDate(year int, month int, day int) (*Date, error) {
	if month < 1 || month > 12 || day < 1 || day > getDaysInMonth(year, month) {
		return nil, fmt.Errorf("invalid date")
	}

	return &Date{
		Year:  year,
		Month: month,
		Day:   day,
	}, nil
}

// NewDateFromStr parses a date string into a Date instance. The default format for the date string is "2006-01-02", but you can specify a custom format.
func NewDateFromStr(dateStr string, layout ...string) (*Date, error) {

	var formatString = "2006-01-02"
	if len(layout) > 0 {
		formatString = layout[0]
	}

	t, err := time.Parse(formatString, dateStr)
	if err != nil {
		return nil, err
	}

	return NewDate(t.Year(), int(t.Month()), t.Day())
}

// String converts a Date instance to a string representation. The default format is "2006-01-02", but you can specify a custom format.
func (d *Date) String(layout ...string) string {
	t := time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
	format := "2006-01-02"

	if len(layout) > 0 {
		format = layout[0]
	}

	return t.Format(format)
}

// Increase increments the current date by one day.
func (d *Date) Increase() {
	d.AddDays(1)
}

// Decrease decrements the current date by one day.
func (d *Date) Decrease() {
	d.SubDays(1)
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

// Equal checks if the current date is equal to x.
func (d *Date) Equal(x *Date) bool {
	return d.Year == x.Year && d.Month == x.Month && d.Day == x.Day
}

// EarlyThan checks if the current date is earlier than x.
func (d *Date) EarlyThan(x *Date) bool {
	if d.Year < x.Year {
		return true
	} else if d.Year != x.Year {
		return false
	}
	if d.Month < x.Month {
		return true
	} else if d.Month != x.Month {
		return false
	}

	return d.Day < x.Day
}

// LaterThan checks if the current date is later than x.
func (d *Date) LaterThan(x *Date) bool {
	return !d.EarlyThanOrEqual(x)
}

// EarlyThanOrEqual checks if the current date is not later than x.
func (d *Date) EarlyThanOrEqual(x *Date) bool {
	return d.EarlyThan(x) || d.Equal(x)
}

// LaterThanOrEqual checks if the current date is not earlier than x.
func (d *Date) LaterThanOrEqual(x *Date) bool {
	return d.LaterThan(x) || d.Equal(x)
}

// AddDays adds the specified number of days to the current date.
func (d *Date) AddDays(days int) {
	for i := 0; i < days; i++ {
		d.Day++
		if d.Day > d.DaysInMonth() {
			d.Day = 1
			d.Month++
			if d.Month > 12 {
				d.Month = 1
				d.Year++
			}
		}
	}

	for i := 0; i > days; i-- {
		d.Day--
		if d.Day < 1 {
			d.Month--
			if d.Month < 1 {
				d.Month = 12
				d.Year--
			}
			d.Day = d.DaysInMonth()
		}
	}
}

// SubDays subtracts the specified number of days from the current date.
func (d *Date) SubDays(days int) {
	d.AddDays(-days)
}

// DaysDifference calculates the number of days between the current date and x.
func (d *Date) DaysDifference(x *Date) int {
	d1 := d.clone()
	d2 := x.clone()

	if d1.Equal(d2) {
		return 0
	}

	if d1.LaterThan(d2) {
		return daysDifference(d1, d2)
	}

	return -daysDifference(d1, d2)
}

// Week returns the day of the week for the current date. It returns one of the constants from the weekArray.
func (d *Date) Week() string {
	// 2021/3/5 -> Friday
	days := d.DaysDifference(birthDay)

	result := days % 7
	if result < 0 {
		result += 7
	}

	return weekArray[result]
}

// IsLeapYear checks if the current year is a leap year.
func (d *Date) IsLeapYear() bool {
	return isLeapYear(d.Year)
}

// DaysInYear returns the number of days in the current year.
func (d *Date) DaysInYear() int {
	if d.IsLeapYear() {
		return 366
	}
	return 365
}

// DayOfYear returns the day of the year for the current date.
func (d *Date) DayOfYear() int {
	days := d.Day
	for month := 1; month < d.Month; month++ {
		days += getDaysInMonth(d.Year, month)
	}
	return days
}

// WeekOfYear returns the week number of the current date within the year.
func (d *Date) WeekOfYear() int {
	days := d.DayOfYear()
	if days%7 != 0 {
		return days/7 + 1
	}
	return days / 7
}

// DaysInMonth returns the number of days in the current month.
func (d *Date) DaysInMonth() int {
	return getDaysInMonth(d.Year, d.Month)
}

// IsValid checks if the current date is valid.
func (d *Date) IsValid() bool {
	if d.Year < 0 || d.Month < 1 || d.Month > 12 || d.Day < 1 {
		return false
	}

	daysInMonth := d.DaysInMonth()
	if d.Day > daysInMonth {
		return false
	}

	return true
}

// AddDaysOfYear converts the specified day of the year to a date. If the day exceeds the number of days in the year, it will roll over to the next year.
func AddDaysOfYear(days int, year int) *Date {
	date := Date{Year: year, Month: 1, Day: 1}

	for days > date.DaysInYear() {
		days -= date.DaysInYear()
		date.Year++
	}

	date.Month = 1
	for days > getDaysInMonth(date.Year, date.Month) {
		days -= getDaysInMonth(date.Year, date.Month)
		date.Month++
	}

	date.Day = days

	return &date
}

// Today returns the current date.
func Today() *Date {
	t := time.Now()
	year, month, day := t.Date()
	d, _ := NewDate(year, int(month), day)

	return d
}

// IsToday checks if the specified date is today.
func (d *Date) IsToday() bool {
	return d.Equal(Today())
}

// BeginOfThisYear returns the first day of the current year.
func (d *Date) BeginOfThisYear() *Date {
	d1, _ := NewDate(d.Year, 1, 1)

	return d1
}

// BeginOfThisMonth returns the first day of the current month.
func (d *Date) BeginOfThisMonth() *Date {
	d1, _ := NewDate(d.Year, d.Month, 1)

	return d1
}

// ToTimestamp converts the Date instance to a Unix timestamp (int64).
func (d *Date) ToTimestamp() int64 {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC).Unix()
}

// =========== helper ===========

func getDaysInMonth(year, month int) int {
	switch month {
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func daysDifference(d1, d2 *Date) int {
	if d1.Equal(d2) {
		return 0
	}

	// Ensure d1 is greater than d2
	if d1.EarlyThan(d2) {
		d1, d2 = d2, d1
	}

	diff := 0

	for d1.Year > d2.Year {
		diff += d2.DaysInYear()
		d2.Year++
	}

	return diff + d1.DayOfYear() - d2.DayOfYear()
}

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
	Assert(month > 0 && month < 13, "月数不在合理范围内")
	Assert(day > 0 && day < 32, "日数不在合理范围内")
	return &Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

// 以字符串的形式返回
func (d *Date) String(format ...string) string {
	var formatString = "%d/%d/%d"
	if len(format) > 0 {
		formatString = format[0]
	}

	return fmt.Sprintf(formatString, d.Year, d.Month, d.Day)
}

func (d *Date) Copy(x *Date) {
	d.Year = x.Year
	d.Month = x.Month
	d.Day = x.Day
}

// 加几天
func (d *Date) AddDay(x int) {
	Assert(x >= 0, "x 必须大于0")
	days := d.DaysOfYear() + x
	d.Month, d.Day = 1, 1
	for days > d.Days() {
		days -= d.Days()
		d.Year++
	}

	d.Copy(TurnDaysToDate(days, d.Year))
}

// 减具体一天 = 时间间隔
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

// 减几天
func (d *Date) SubDay(x int) {
	Assert(x >= 0, "x 必须大于0")
	days := d.DaysOfYear() - x - 1
	d.Month, d.Day = 1, 1

	for days <= 0 {
		d.Year--
		days += d.Days()
	}

	d.AddDay(days)
}

// 返回当天是星期几
func (d *Date) Week() string {
	// 2021/3/5 -> Friday
	days := d.SubDate(BirthDay)
	if days < 0 {
		days = -days
	}

	fmt.Println(days, days%7)
	return WeekArray[days%7]
}

// 返回今年是否是平年
func (d *Date) IsLeap() bool {
	return !(d.Year%400 == 0 || (d.Year%4 == 0 && d.Year%100 != 0))
}

// 返回今年有多少天
func (d *Date) Days() int {
	if d.IsLeap() {
		return 365
	}
	return 366
}

// 返回这一天是今年的第几天
func (d *Date) DaysOfYear() int {
	total := 0
	// 前面几个月的天数加上这个月的日数
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

// 判断这天是否合理
func (d *Date) Check() bool {
	// 大月肯定是31天
	if d.IsBigMonth() {
		return true
	}
	// 小月但不是2月肯定是30
	if d.Month != 2 {
		return d.Day < 31
	}
	// 2月平年28
	if d.IsLeap() {
		return d.Day < 29
	}
	// 2月闰年29
	return d.Day < 30
}

// 判断这个月是大月还是小月
func (d *Date) IsBigMonth() bool {
	return isBigMonth(d.Month)
}

func isBigMonth(m int) bool {
	return MonthArray[m-1]
}

// DaysOfYear的反函数
func TurnDaysToDate(days int, year int) *Date {
	d := &Date{
		Year:  year,
		Month: 1,
		Day:   1,
	}
	monthDay := ALeapMonthDayArray
	if d.IsLeap() {
		Assert(days <= 365 && days > 0, "平年365天")
		monthDay = LeapMonthDayArray
	} else {
		Assert(days <= 366 && days > 0, "闰年366天")
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

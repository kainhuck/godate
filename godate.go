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

// birthDay this is th datetime when i create this project
var birthDay = &Date{
	Year:  2021,
	Month: 3,
	Day:   5,
}

// NewDate
// 通过指定的年月日创建一个 Date 实例，日期范围有误会报错
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

// NewDateFromStr
// 解析日期字符串到一个 Date 实例，默认的字符串格式为 2006-01-02 可以手动指定字符串格式
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

// String
// 将 Date 实例转成字符串输出，默认格式为 2006-01-02 可手动指定格式
func (d *Date) String(layout ...string) string {
	t := time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
	format := "2006-01-02"

	if len(layout) > 0 {
		format = layout[0]
	}

	return t.Format(format)
}

// Increase
// 将当前对象天数+1
func (d *Date) Increase() {
	d.AddDays(1)
}

// Decrease
// 将当前对象天数-1
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

// Equal
// 判断当前日期是否和 x 相等
func (d *Date) Equal(x *Date) bool {
	return d.Year == x.Year && d.Month == x.Month && d.Day == x.Day
}

// EarlyThan
// 判断当前日期是否早于 x
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

// LaterThan
// 判断当前日期是否晚于 x
func (d *Date) LaterThan(x *Date) bool {
	return !d.EarlyThanOrEqual(x)
}

// EarlyThanOrEqual
// 判断当前日期是否不晚于 x
func (d *Date) EarlyThanOrEqual(x *Date) bool {
	return d.EarlyThan(x) || d.Equal(x)
}

// LaterThanOrEqual
// 判断当前日期是否不早于 x
func (d *Date) LaterThanOrEqual(x *Date) bool {
	return d.LaterThan(x) || d.Equal(x)
}

// AddDays
// 将档期日期加上 days 天
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

// SubDays
// 将当前日期减去 days 天
func (d *Date) SubDays(days int) {
	d.AddDays(-days)
}

// DaysDifference
// 计算当前日期和 x 日期相差的天数
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

// Week
// 返回今天是星期几
// 返回的是 weekArray 中的常量
func (d *Date) Week() string {
	// 2021/3/5 -> Friday
	days := d.DaysDifference(birthDay)

	result := days % 7
	if result < 0 {
		result += 7
	}

	return weekArray[result]
}

// IsLeapYear
// 判断今年是否是闰年
func (d *Date) IsLeapYear() bool {
	return isLeapYear(d.Year)
}

// DaysInYear
// 返回今年共有几天
func (d *Date) DaysInYear() int {
	if d.IsLeapYear() {
		return 366
	}
	return 365
}

// DayOfYear
// 返回今天是这一年中的第几天
func (d *Date) DayOfYear() int {
	days := d.Day
	for month := 1; month < d.Month; month++ {
		days += getDaysInMonth(d.Year, month)
	}
	return days
}

// WeekOfYear
// 返回这个星期是这年中的第几个星期
func (d *Date) WeekOfYear() int {
	days := d.DayOfYear()
	if days%7 != 0 {
		return days/7 + 1
	}
	return days / 7
}

// DaysInMonth
// 返回这个月共有几天
func (d *Date) DaysInMonth() int {
	return getDaysInMonth(d.Year, d.Month)
}

// IsValid
// 检查当前日期是否正确
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

// AddDaysOfYear
// 将某年的第几天转成日期，注意如果days超过当年天数，则会顺延到下一年
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

// Today
// 返回今天日期
func Today() *Date {
	t := time.Now()
	year, month, day := t.Date()
	d, _ := NewDate(year, int(month), day)

	return d
}

// IsToday
// 判断指定的日期是否是今天
func (d *Date) IsToday() bool {
	return d.Equal(Today())
}

// BeginOfThisYear
// 返回当前日期所在年份的第一天
func (d *Date) BeginOfThisYear() *Date {
	d1, _ := NewDate(d.Year, 1, 1)

	return d1
}

// BeginOfThisMonth
// 返回当前日期所在月份的第一天
func (d *Date) BeginOfThisMonth() *Date {
	d1, _ := NewDate(d.Year, d.Month, 1)

	return d1
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

	// 确保 d1 大于 d2
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

package godate

import (
	"reflect"
	"testing"
)

func TestAddDaysOfYear(t *testing.T) {
	type args struct {
		days int
		year int
	}
	tests := []struct {
		name string
		args args
		want *Date
	}{
		{
			"t1",
			args{days: 10, year: 2023},
			&Date{
				Year:  2023,
				Month: 1,
				Day:   10,
			},
		},
		{
			"t2",
			args{days: 40, year: 2023},
			&Date{
				Year:  2023,
				Month: 2,
				Day:   9,
			},
		},
		{
			"t3",
			args{days: 200, year: 2023},
			&Date{
				Year:  2023,
				Month: 7,
				Day:   19,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDaysOfYear(tt.args.days, tt.args.year); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDaysOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_AddDays(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	type args struct {
		days int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 1,
				Day:   1,
			},
			args{days: 200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			d.AddDays(tt.args.days)
		})
	}
}

func TestDate_BeginOfThisMonth(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Date
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 6,
				Day:   30,
			},
			&Date{
				Year:  2023,
				Month: 6,
				Day:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.BeginOfThisMonth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeginOfThisMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_BeginOfThisYear(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Date
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 6,
				Day:   30,
			},
			&Date{
				Year:  2023,
				Month: 1,
				Day:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.BeginOfThisYear(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeginOfThisYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_DaysDifference(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	type args struct {
		x *Date
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 6,
				Day:   30,
			},
			args{x: &Date{
				Year:  2019,
				Month: 12,
				Day:   1,
			}},
			1307,
		},
		{
			"t2",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2023,
				Month: 6,
				Day:   30,
			}},
			-1307,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.DaysDifference(tt.args.x); got != tt.want {
				t.Errorf("DaysDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_DaysInMonth(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 6,
				Day:   1,
			},
			30,
		},
		{
			"t2",
			fields{
				Year:  2020,
				Month: 2,
				Day:   1,
			},
			29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.DaysInMonth(); got != tt.want {
				t.Errorf("DaysInMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_DaysInYear(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 1,
				Day:   1,
			},
			365,
		},
		{
			"t2",
			fields{
				Year:  2020,
				Month: 1,
				Day:   1,
			},
			366,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.DaysInYear(); got != tt.want {
				t.Errorf("DaysInYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_DayOfYear(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 1,
				Day:   30,
			},
			30,
		},
		{
			"t2",
			fields{
				Year:  2023,
				Month: 6,
				Day:   30,
			},
			181,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.DayOfYear(); got != tt.want {
				t.Errorf("DaysOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Decrease(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 6,
				Day:   30,
			},
		},
		{
			"t2",
			fields{
				Year:  2023,
				Month: 1,
				Day:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			d.Decrease()
		})
	}
}

func TestDate_EarlyThan(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	type args struct {
		x *Date
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"t1",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2023,
				Month: 6,
				Day:   30,
			}},
			true,
		},
		{
			"t2",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2019,
				Month: 12,
				Day:   1,
			}},
			false,
		},
		{
			"t3",
			fields{
				Year:  2119,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2023,
				Month: 6,
				Day:   30,
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.EarlyThan(tt.args.x); got != tt.want {
				t.Errorf("EarlyThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_EarlyThanOrEqual(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	type args struct {
		x *Date
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"t1",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2023,
				Month: 6,
				Day:   30,
			}},
			true,
		},
		{
			"t2",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2019,
				Month: 12,
				Day:   1,
			}},
			true,
		},
		{
			"t3",
			fields{
				Year:  2119,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2023,
				Month: 6,
				Day:   30,
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.EarlyThanOrEqual(tt.args.x); got != tt.want {
				t.Errorf("EarlyThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Equal(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	type args struct {
		x *Date
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"t1",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2023,
				Month: 6,
				Day:   30,
			}},
			false,
		},
		{
			"t2",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2019,
				Month: 12,
				Day:   1,
			}},
			true,
		},
		{
			"t3",
			fields{
				Year:  2119,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2023,
				Month: 6,
				Day:   30,
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.Equal(tt.args.x); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Increase(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"t1",
			fields{
				Year:  2019,
				Month: 12,
				Day:   31,
			},
		},
		{
			"t2",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			d.Increase()
		})
	}
}

func TestDate_IsLeapYear(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"t1",
			fields{
				Year:  2019,
				Month: 1,
				Day:   1,
			},
			false,
		},
		{
			"t2",
			fields{
				Year:  2020,
				Month: 1,
				Day:   1,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.IsLeapYear(); got != tt.want {
				t.Errorf("IsLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_IsToday(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"t1",
			fields{
				Year:  Today().Year,
				Month: Today().Month,
				Day:   Today().Day,
			},
			true,
		},
		{
			"t2",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.IsToday(); got != tt.want {
				t.Errorf("IsToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_IsValid(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 2,
				Day:   29,
			},
			false,
		},
		{
			"t2",
			fields{
				Year:  2020,
				Month: 2,
				Day:   29,
			},
			true,
		},
		{
			"t3",
			fields{
				Year:  2023,
				Month: 4,
				Day:   31,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_LaterThan(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	type args struct {
		x *Date
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"t1",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2023,
				Month: 6,
				Day:   30,
			}},
			false,
		},
		{
			"t2",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2019,
				Month: 12,
				Day:   1,
			}},
			false,
		},
		{
			"t3",
			fields{
				Year:  2119,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2023,
				Month: 6,
				Day:   30,
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.LaterThan(tt.args.x); got != tt.want {
				t.Errorf("LaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_LaterThanOrEqual(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	type args struct {
		x *Date
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"t1",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2023,
				Month: 6,
				Day:   30,
			}},
			false,
		},
		{
			"t2",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2019,
				Month: 12,
				Day:   1,
			}},
			true,
		},
		{
			"t3",
			fields{
				Year:  2119,
				Month: 12,
				Day:   1,
			},
			args{x: &Date{
				Year:  2023,
				Month: 6,
				Day:   30,
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.LaterThanOrEqual(tt.args.x); got != tt.want {
				t.Errorf("LaterThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_String(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	type args struct {
		layout []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"t1",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{layout: nil},
			"2019-12-01",
		},
		{
			"t2",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			args{layout: []string{"01/02/2006"}},
			"12/01/2019",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.String(tt.args.layout...); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_SubDays(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	type args struct {
		days int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 1,
				Day:   7,
			},
			args{days: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			d.SubDays(tt.args.days)
		})
	}
}

func TestDate_Week(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 6,
				Day:   30,
			},
			Friday,
		},
		{
			"t2",
			fields{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			Sunday,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.Week(); got != tt.want {
				t.Errorf("Week() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_WeekOfYear(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 1,
				Day:   1,
			},
			1,
		},
		{
			"t2",
			fields{
				Year:  2023,
				Month: 6,
				Day:   30,
			},
			26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.WeekOfYear(); got != tt.want {
				t.Errorf("WeeksOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_clone(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Date
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 1,
				Day:   2,
			},
			&Date{
				Year:  2023,
				Month: 1,
				Day:   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_copy(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	type args struct {
		x *Date
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"t1",
			fields{
				Year:  2023,
				Month: 1,
				Day:   2,
			},
			args{x: &Date{
				Year:  2023,
				Month: 1,
				Day:   2,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			d.copy(tt.args.x)
		})
	}
}

func TestNewDate(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		args    args
		want    *Date
		wantErr bool
	}{
		{
			"t1",
			args{
				year:  2019,
				month: 12,
				day:   1,
			},
			&Date{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			false,
		},
		{
			"t1",
			args{
				year:  2019,
				month: 12,
				day:   110,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDate(tt.args.year, tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDateFromStr(t *testing.T) {
	type args struct {
		dateStr string
		layout  []string
	}
	tests := []struct {
		name    string
		args    args
		want    *Date
		wantErr bool
	}{
		{
			"t1",
			args{
				dateStr: "2019/12/01",
				layout:  []string{"2006/01/02"},
			},
			&Date{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			false,
		},
		{
			"t2",
			args{
				dateStr: "2019-12-01",
				layout:  nil,
			},
			&Date{
				Year:  2019,
				Month: 12,
				Day:   1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDateFromStr(tt.args.dateStr, tt.args.layout...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDateFromStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDateFromStr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToday(t *testing.T) {
	tests := []struct {
		name string
		want *Date
	}{
		{
			"t1",
			Today(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Today(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Today() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDaysInMonth(t *testing.T) {
	type args struct {
		year  int
		month int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"t1",
			args{
				year:  2019,
				month: 12,
			},
			31,
		},
		{
			"t2",
			args{
				year:  2019,
				month: 2,
			},
			28,
		},
		{
			"t3",
			args{
				year:  2020,
				month: 2,
			},
			29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDaysInMonth(tt.args.year, tt.args.month); got != tt.want {
				t.Errorf("getDaysInMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"t1",
			args{year: 1000},
			false,
		},
		{
			"t2",
			args{year: 2019},
			false,
		},
		{
			"t3",
			args{year: 2020},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLeapYear(tt.args.year); got != tt.want {
				t.Errorf("isLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

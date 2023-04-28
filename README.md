# godate
a friendly date operation of golang

## install
```
go get github.com/kainhuck/godate
```

# Usage
```golang
package main

import (
	"fmt"
	"github.com/kainhuck/godate"
)

func main() {
	d1,_ := godate.NewDate(2021, 10, 22)
	d2,_ := godate.NewDate(1999, 1, 23)
	d3,_ := godate.NewDateFromStr("2019-12-01")
    d4 := godate.Today()
	fmt.Println(d3.String())
    fmt.Println(d4.String("%d/%02d/%02d"))
    fmt.Println(d1.Equal(d2))
    fmt.Println(d1.EarlyThan(d2))
    fmt.Println(d1.LaterThan(d2))
    
    fmt.Println(d1.EarlyThanOrEqual(d2))
    fmt.Println(d1.LaterThanOrEqual(d2))
	
	d1.AddWeek(2)
    fmt.Println(d1.String())
	d1.SubWeek(2)
    fmt.Println(d1.String())

	if !d1.Check() || !d2.Check(){
		panic("wrong date")
	}

	d1.AddDay(200)
	fmt.Println(d1.String())

	d1.SubDay(123)
	fmt.Println(d1.String())

	days := d1.SubDate(d2)
	fmt.Println(days)

	fmt.Println(d1.Week())

	fmt.Println(d1.DaysOfYear())

	fmt.Println(d1.Days())

	fmt.Println(d1.IsLeap())

	fmt.Println(d1.IsBigMonth())

	d1.Increase()
	fmt.Println(d1)  // add one day

	d1.Decrease()
	fmt.Println(d1) // sub one day
}
```
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
	d1 := godate.NewDate(2021, 10, 22)
	d2 := godate.NewDate(1999, 1, 23)
	d3 := godate.NewDateByStr("2019-12-01")
    d4 := godate.Today()
	fmt.Println(d3.String())
    fmt.Println(d4.String("%d/%02d/%02d"))
    fmt.Println(d1.Equal(d2))

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
	
	fmt.Println(d1.Increase())  // add one day
	
	fmt.Println(d1.Reduce()) // sub one day
}
```
package godate

import (
	"fmt"
	"log"
	"testing"
)

func TestDate(t *testing.T) {
	//d := NewDate(2021, 1, 3)
	////d.SubDay(3000)
	//d.AddDay(2000)
	//fmt.Println(d.String())
	//d1 := NewDate(2023, 3, 6)
	//d2 := NewDate(2020, 1, 1)

	//d := d2.SubDate(d1)
	//fmt.Println(d)
	//fmt.Println(d1.Week())
	//d2.AddDay(456)
	//fmt.Println(d2.String())

	//d := NewDateByStr("2021-3-6")
	//fmt.Println(d.Week())
	d1, err := NewDateByStr("2019-12-1")
	if err != nil {
		log.Fatal(err)
	}
	//d2 := Today()
	//fmt.Println(d2.SubDate(d1))
	//
	//d1.AddDay(500)
	//fmt.Println(d1.String())
	//fmt.Println(d1.Week())
	c := 0
	i := 100
	for c < 10 {
		d1.AddDay(i)
		if d1.Week() == Saturday || d1.Week() == Sunday {
			fmt.Printf("我们在一起的第%d天是%s,这一天是%s\n", i, d1.String(), d1.Week())
			c++
		}
		d1.SubDay(i)
		i += 100
	}
}

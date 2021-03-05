package godate

import (
	"fmt"
	"testing"
)

func TestDate(t *testing.T) {
	//d := NewDate(2021, 1, 3)
	////d.SubDay(3000)
	//d.AddDay(2000)
	//fmt.Println(d.String())
	d1 := NewDate(2023, 3, 6)
	//d2 := NewDate(2020, 1, 1)

	//d := d2.SubDate(d1)
	//fmt.Println(d)
	fmt.Println(d1.Week())
	//d2.AddDay(456)
	//fmt.Println(d2.String())
}

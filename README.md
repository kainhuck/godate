# GoDate

GoDate is a Go library for date manipulation, providing functions for creating, comparing, and calculating dates.

## Usage

### Creating Date Instances

```go
package main

import (
	"fmt"
	"github.com/kainhuck/godate"
)

func main() {
	// Create a date instance by specifying the year, month, and day
	date1, err := godate.NewDate(2023, 6, 30)
	if err != nil {
		fmt.Println("Failed to create date:", err)
		return
	}
	fmt.Println("Date 1:", date1)

	// Create a date instance by parsing a date string
	date2, err := godate.NewDateFromStr("2023-06-30", "2006-01-02")
	if err != nil {
		fmt.Println("Failed to create date:", err)
		return
	}
	fmt.Println("Date 2:", date2)
}
```

### Date Operations

```go
package main

import (
	"fmt"
	"github.com/kainhuck/godate"
)

func main() {
	// Create a date instance
	date, _ := godate.NewDate(2023, 6, 30)

	// Increase the date by one day
	date.Increase()
	fmt.Println("Date after increasing by one day:", date)

	// Decrease the date by one day
	date.Decrease()
	fmt.Println("Date after decreasing by one day:", date)

	// Check if the date is equal to another date
	otherDate, _ := godate.NewDate(2023, 6, 30)
	fmt.Println("Is the date equal to another date:", date.Equal(otherDate))

	// Check if the date is earlier than another date
	anotherDate, _ := godate.NewDate(2023, 7, 1)
	fmt.Println("Is the date earlier than another date:", date.EarlyThan(anotherDate))

	// Check if the date is later than another date
	fmt.Println("Is the date later than another date:", date.LaterThan(anotherDate))

	// Check if the date is not later than another date
	fmt.Println("Is the date not later than another date:", date.EarlyThanOrEqual(anotherDate))

	// Check if the date is not earlier than another date
	fmt.Println("Is the date not earlier than another date:", date.LaterThanOrEqual(anotherDate))

	// Add a specified number of days
	date.AddDays(5)
	fmt.Println("Date after adding 5 days:", date)

	// Subtract a specified number of days
	date.SubDays(3)
	fmt.Println("Date after subtracting 3 days:", date)

	// Calculate the number of days difference between two dates
	diff := date.DaysDifference(anotherDate)
	fmt.Println("Date difference:", diff)
}
```

## Functionality

### Types

#### `type Date struct`

A structure representing a date, with year, month, and day fields.

#### `type Date`

Methods for the Date structure:

- `func NewDate(year int, month int, day int) (*Date, error)`

Creates a new Date instance with the specified year, month, and day.

- `func NewDateFromStr(dateStr string, layout ...string) (*Date, error)`

Parses a date string into a Date instance. The default date string format is "2006-01-02", but you can specify a custom format.

- `func (d *Date) String(layout ...string) string`

Converts the Date instance to a string representation. The default format is "2006

-01-02", but you can specify a custom format.

- `func (d *Date) Increase()`

Increases the date by one day.

- `func (d *Date) Decrease()`

Decreases the date by one day.

- `func (d *Date) Equal(other *Date) bool`

Checks if the date is equal to the other date.

- `func (d *Date) EarlyThan(other *Date) bool`

Checks if the date is earlier than the other date.

- `func (d *Date) LaterThan(other *Date) bool`

Checks if the date is later than the other date.

- `func (d *Date) EarlyThanOrEqual(other *Date) bool`

Checks if the date is not later than the other date.

- `func (d *Date) LaterThanOrEqual(other *Date) bool`

Checks if the date is not earlier than the other date.

- `func (d *Date) AddDays(days int)`

Adds the specified number of days to the date.

- `func (d *Date) SubDays(days int)`

Subtracts the specified number of days from the date.

- `func (d *Date) DaysDifference(other *Date) int`

Calculates the number of days difference between two dates.

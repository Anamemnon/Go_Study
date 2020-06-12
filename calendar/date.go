package calendar

import "errors"

var invalidYearErr = errors.New("invalid year")
var invalidMonthErr = errors.New("invalid month")
var invalidDayOfMonthErr = errors.New("invalid day of month")

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return invalidYearErr
	}

	d.year = year

	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return invalidMonthErr
	}

	d.month = month

	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return invalidDayOfMonthErr
	}

	d.day = day

	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}
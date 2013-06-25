package paymill

import (
	"net/url"
	"strconv"
	"strings"
)

type Resource interface {
	ResourceId() string
	ResourceName() string
	CreateValues() url.Values
	UpdateValues() url.Values
}

type Unit string

const (
	DayUnit   Unit = "DAY"
	WeekUnit  Unit = "WEEK"
	MonthUnit Unit = "MONTH"
	YearUnit  Unit = "YEAR"
)

func (u Unit) Valid() bool {
	switch u {
	case DayUnit, WeekUnit, MonthUnit, YearUnit:
		return true
	}
	return false
}

type Interval string

func (i Interval) Num() int {
	c := strings.SplitN(string(i), " ", 2)
	if len(c) == 2 {
		i, _ := strconv.Atoi(c[0])
		return i
	}
	return 0
}

func (i Interval) Unit() Unit {
	c := strings.SplitN(string(i), " ", 2)
	if len(c) == 2 {
		return Unit(c[1])
	}
	return ""
}

func (i Interval) Valid() bool {
	return i.Num() > 0 && i.Unit().Valid()
}

// Numeric is a type that unmarshals int or string numeric values.
// Unfortunately we need this type, because Paymill returns inconsistent results.
type Numeric string

func (n *Numeric) UnmarshalJSON(b []byte) error {
	var s string
	if len(b) > 0 && b[0] == '"' {
		s2, err := strconv.Unquote(string(b))
		if err != nil {
			return err
		}
		s = s2
	} else {
		s = string(b)
	}
	*n = Numeric(s)
	return nil
}

func (n *Numeric) Int64() int64 {
	i, _ := strconv.ParseInt(string(*n), 10, 64)
	return i
}

type Result struct {
	Data Resource
}

type ResultList struct {
	Data []Resource
}

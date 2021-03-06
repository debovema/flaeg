package flaeg

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

//TODO : add parsers on all types in https://golang.org/pkg/builtin/

// Parser is an interface that allows the contents of a flag.Getter to be set.
type Parser interface {
	flag.Getter
	SetValue(interface{})
}

// -- bool Value
type boolValue bool

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

func (b *boolValue) Get() interface{} { return bool(*b) }

func (b *boolValue) String() string { return fmt.Sprintf("%v", *b) }

func (b *boolValue) IsBoolFlag() bool { return true }

func (b *boolValue) SetValue(val interface{}) {
	*b = boolValue(val.(bool))
}

// optional interface to indicate boolean flags that can be
// supplied without "=value" text
type boolFlag interface {
	flag.Value
	IsBoolFlag() bool
}

// -- int Value
type intValue int

func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = intValue(v)
	return err
}

func (i *intValue) Get() interface{} { return int(*i) }

func (i *intValue) String() string { return fmt.Sprintf("%v", *i) }

func (i *intValue) SetValue(val interface{}) {
	*i = intValue(val.(int))
}

// -- int64 Value
type int64Value int64

func (i *int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = int64Value(v)
	return err
}

func (i *int64Value) Get() interface{} { return int64(*i) }

func (i *int64Value) String() string { return fmt.Sprintf("%v", *i) }

func (i *int64Value) SetValue(val interface{}) {
	*i = int64Value(val.(int64))
}

// -- uint Value
type uintValue uint

func (i *uintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uintValue(v)
	return err
}

func (i *uintValue) Get() interface{} { return uint(*i) }

func (i *uintValue) String() string { return fmt.Sprintf("%v", *i) }

func (i *uintValue) SetValue(val interface{}) {
	*i = uintValue(val.(uint))
}

// -- uint64 Value
type uint64Value uint64

func (i *uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uint64Value(v)
	return err
}

func (i *uint64Value) Get() interface{} { return uint64(*i) }

func (i *uint64Value) String() string { return fmt.Sprintf("%v", *i) }

func (i *uint64Value) SetValue(val interface{}) {
	*i = uint64Value(val.(uint64))
}

// -- string Value
type stringValue string

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

func (s *stringValue) Get() interface{} { return string(*s) }

func (s *stringValue) String() string { return fmt.Sprintf("%s", *s) }

func (s *stringValue) SetValue(val interface{}) {
	*s = stringValue(val.(string))
}

// -- float64 Value
type float64Value float64

func (f *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = float64Value(v)
	return err
}

func (f *float64Value) Get() interface{} { return float64(*f) }

func (f *float64Value) String() string { return fmt.Sprintf("%v", *f) }

func (f *float64Value) SetValue(val interface{}) {
	*f = float64Value(val.(float64))
}

// -- time.Duration Value
type durationValue time.Duration

func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = durationValue(v)
	return err
}

func (d *durationValue) Get() interface{} { return time.Duration(*d) }

func (d *durationValue) String() string { return (*time.Duration)(d).String() }

func (d *durationValue) SetValue(val interface{}) {
	*d = durationValue(val.(time.Duration))
}

// -- time.Time Value
type timeValue time.Time

func (t *timeValue) Set(s string) error {
	v, err := time.Parse(time.RFC3339, s)
	*t = timeValue(v)
	return err
}

func (t *timeValue) Get() interface{} { return time.Time(*t) }

func (t *timeValue) String() string { return (*time.Time)(t).String() }

func (t *timeValue) SetValue(val interface{}) {
	*t = timeValue(val.(time.Time))
}

package aladin

import (
	"strconv"
	"strings"
	"time"
)

var (
	_ Value = (*value)(nil)
	_ Value = (*stringValue)(nil)
)

type Value interface {
	String(string) string
	StringSlice([]string) []string
	Int64(int64) int64
	Uint64(uint64) uint64
	Int(int) int
	Uint(uint) uint
	Float64(float64) float64
	Bool(bool) bool
	Duration(time.Duration) time.Duration
}

type value struct{}

func defValue() Value { return new(value) }

// def 是默认值，即当Store为空时，能够获取到默认值，而不是返回error
func (value) String(def string) string                 { return def }
func (value) StringSlice(def []string) []string        { return def }
func (value) IntSlice(def []int) []int                 { return def }
func (value) Int64(def int64) int64                    { return def }
func (value) Uint64(def uint64) uint64                 { return def }
func (value) Int(def int) int                          { return def }
func (value) Uint(def uint) uint                       { return def }
func (value) Float64(def float64) float64              { return def }
func (value) Bool(def bool) bool                       { return def }
func (value) Duration(def time.Duration) time.Duration { return def }

type stringValue string

func (v stringValue) String(def string) string {
	if len(v) == 0 {
		return def
	}
	return string(v)
}

func (v stringValue) StringSlice(def []string) []string {
	if len(v) == 0 {
		return def
	}
	return strings.Split(string(v), ",")
}

func (v stringValue) Int64(def int64) int64 {
	if len(v) == 0 {
		return def
	}
	i, err := strconv.ParseInt(string(v), 10, 64)
	if err != nil {
		return def
	}
	return i
}

func (v stringValue) Uint64(def uint64) uint64 {
	if len(v) == 0 {
		return def
	}
	i, err := strconv.ParseUint(string(v), 10, 64)
	if err != nil {
		return def
	}
	return i
}

func (v stringValue) Int(def int) int {
	if len(v) == 0 {
		return def
	}
	return int(v.Int64(int64(def)))
}

func (v stringValue) Uint(def uint) uint {
	if len(v) == 0 {
		return def
	}
	return uint(v.Uint64(uint64(def)))
}

func (v stringValue) Float64(def float64) float64 {
	if len(v) == 0 {
		return def
	}
	f, err := strconv.ParseFloat(string(v), 64)
	if err != nil {
		return def
	}
	return f
}

func (v stringValue) Bool(def bool) bool {
	if len(v) == 0 {
		return def
	}
	b, err := strconv.ParseBool(string(v))
	if err != nil {
		return def
	}
	return b
}

func (v stringValue) Duration(def time.Duration) time.Duration {
	if len(v) == 0 {
		return def
	}
	d, err := time.ParseDuration(string(v))
	if err != nil {
		return def
	}
	return d
}

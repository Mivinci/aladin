package aladin

import "time"

type Value interface {
	String(string) string
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
func (value) Int64(def int64) int64                    { return def }
func (value) Uint64(def uint64) uint64                 { return def }
func (value) Int(def int) int                          { return def }
func (value) Uint(def uint) uint                       { return def }
func (value) Float64(def float64) float64              { return def }
func (value) Bool(def bool) bool                       { return def }
func (value) Duration(def time.Duration) time.Duration { return def }

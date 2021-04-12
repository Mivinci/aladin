package aladin

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/valyala/fastjson"
)

type jsonParser struct {
	p fastjson.Parser
}

func NewJSONParser() Parser {
	return new(jsonParser)
}

func (p *jsonParser) Parse(snap *Snapshot) (Store, error) {
	v, err := p.p.ParseBytes(snap.Data)
	return &jsonStore{v: v}, err
}

type jsonStore struct {
	v *fastjson.Value
}

func (s *jsonStore) Get(path string) Value {
	branch := strings.Split(path, ".")
	v := s.v.Get(branch...)
	if v != nil {
		return &jsonValue{v}
	}
	return defValue()
}

func (s *jsonStore) Set(path string, v interface{}) {
	branch := strings.Split(path, ".")
	trail, leaf := splitPath(branch)
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	s.v.Get(trail...).Set(leaf, fastjson.MustParseBytes(b))
}

func (s *jsonStore) Del(path string) {
	branch := strings.Split(path, ".")
	trail, leaf := splitPath(branch)
	s.v.Get(trail...).Del(leaf)
}

func (s *jsonStore) Scan(v interface{}) error {
	o, err := s.v.Object()
	if err != nil {
		return err
	}
	return json.Unmarshal(o.MarshalTo(nil), v)
}

func splitPath(path []string) ([]string, string) {
	return path[:len(path)-1], path[len(path)-1]
}

type jsonValue struct {
	v *fastjson.Value
}

func (v *jsonValue) String(def string) string {
	b, err := v.v.StringBytes()
	if err != nil {
		return def
	}
	return string(b)
}

func (v *jsonValue) StringSlice(def []string) []string {
	vals, err := v.v.Array()
	if err != nil {
		return def
	}
	a := make([]string, len(vals))
	for i, val := range vals {
		a[i] = val.String()
	}
	return a
}

func (v *jsonValue) Int64(def int64) int64 {
	i, err := v.v.Int64()
	if err != nil {
		return def
	}
	return i
}

func (v *jsonValue) Uint64(def uint64) uint64 {
	u, err := v.v.Uint64()
	if err != nil {
		return def
	}
	return u
}

func (v *jsonValue) Int(def int) int {
	return int(v.Int64(int64(def)))
}

func (v *jsonValue) Uint(def uint) uint {
	return uint(v.Uint64(uint64(def)))
}

func (v *jsonValue) Float64(def float64) float64 {
	f, err := v.v.Float64()
	if err != nil {
		return def
	}
	return f
}

func (v *jsonValue) Bool(def bool) bool {
	b, err := v.v.Bool()
	if err != nil {
		return def
	}
	return b
}

func (v *jsonValue) Duration(def time.Duration) time.Duration {
	switch v.v.Type() {
	case fastjson.TypeNumber:
		return time.Duration(v.Int64(v.v.GetInt64()))
	case fastjson.TypeString:
		d, err := time.ParseDuration(v.v.String())
		if err != nil {
			return def
		}
		return d
	default:
		return def
	}
}

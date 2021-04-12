package aladin

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

const testData = `
name: xjj
age: 18
bio:
  favorites:
    - wfs
    - sing
`

func createTestFile(data []byte) (*os.File, error) {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("aladin.%d", time.Now().Unix()))
	fp, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	_, err = fp.Write(data)
	if err != nil {
		return nil, err
	}
	return fp, nil
}

func ExampleConfig() {
	file, err := createTestFile([]byte(testData))
	if err != nil {
		return
	}

	path := file.Name()
	defer func() {
		file.Close()
		os.Remove(path)
	}()

	conf, err := New(
		WithSource(NewFileSource(path)),
		WithParser(NewYAMLParser()),
	)
	if err != nil {
		return
	}

	fmt.Println(conf.Get("name").String(""))
	fmt.Println(conf.Get("age").Int(0))
	fmt.Println(conf.Get("bio.favorites").StringSlice(nil))

	conf.Close()

	// Output:
	// xjj
	// 18
	// ["wfs" "sing"]
}

func TestSourceEnv(t *testing.T) {
	testCases := []struct {
		k, v   string
		expect interface{}
	}{
		{"NAME", "xjj", "xjj"},
		{"AGE", "18", 18},
		{"BIO_FAVORITES", "wfs,sing", []string{"wfs", "sing"}},
	}

	c, err := New(
		WithSource(NewEnvSource()),
		WithParser(NewEnvParser()),
	)
	if err != nil {
		return
	}
	defer c.Close()

	for _, tc := range testCases {
		os.Setenv(tc.k, tc.v)
	}

	v0 := c.Get("name").String("")
	e0 := testCases[0].expect.(string)
	if v0 != e0 {
		t.Logf("expect: %s, got %s", e0, v0)
		t.FailNow()
	}

	v1 := c.Get("age").Int(0)
	e1 := testCases[1].expect.(int)
	if v1 != e1 {
		t.Logf("expect: %d, got %d", e1, v1)
		t.FailNow()
	}

	v2 := c.Get("bio.favorites").StringSlice(nil)
	e2 := testCases[2].expect.([]string)
	if !reflect.DeepEqual(v2, e2) {
		t.Logf("expect: %v, got %v", e2, v2)
		t.FailNow()
	}
}

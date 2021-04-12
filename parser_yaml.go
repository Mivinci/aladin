package aladin

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

var _ Parser = (*yamlParser)(nil)

type yamlParser struct{}

func NewYAMLParser() Parser {
	return new(yamlParser)
}

func (yamlParser) Parse(snap *Snapshot) (Store, error) {
	data, err := yaml2json(snap.Data)
	if err != nil {
		return nil, err
	}
	snap.Data = data
	return NewJSONParser().Parse(snap)
}

func yaml2json(in []byte) (out []byte, err error) {
	var v interface{}
	err = yaml.Unmarshal(in, &v)
	if err != nil {
		return
	}
	return json.Marshal(v)
}

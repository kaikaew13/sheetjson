package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type UnmarshalledJSON []map[string]interface{}

func Setup(f *os.File) (*OrderedMap, error) {
	uJSON, err := unmarshal(f)
	if err != nil {
		return nil, err
	}

	om := newOrderedMap(uJSON)
	return om, err
}

func unmarshal(f *os.File) (UnmarshalledJSON, error) {
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	uJSON := UnmarshalledJSON{}
	if err = json.Unmarshal(bytes, &uJSON); err != nil {
		return nil, err
	}

	return uJSON, nil
}

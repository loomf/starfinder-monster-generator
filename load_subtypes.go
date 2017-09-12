package main

import (
	"encoding/json"
	"io/ioutil"
)

func LoadSubtypes(filename string) ([]Subtype, error) {
	type file struct {
		Types []Subtype
	}

	var fileSubtypes file

	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &fileSubtypes)
	if err != nil {
		return nil, err
	}

	return fileSubtypes.Types, nil
}

package main

import (
	"encoding/json"
	"io/ioutil"
)

func LoadTypes(filename string) ([]Type, error) {
	type file struct {
		Types []Type
	}

	var fileTypes file

	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &fileTypes)
	if err != nil {
		return nil, err
	}

	return fileTypes.Types, nil
}

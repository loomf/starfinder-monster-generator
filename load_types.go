package main

import (
	"encoding/json"
	"io/ioutil"
)

func LoadTypes(filename string) (map[string]Type, error) {
	type file struct {
		Types map[string]Type
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

	for name, creatureType := range fileTypes.Types {
		creatureType.Name = name
		fileTypes.Types[name] = creatureType
	}

	return fileTypes.Types, nil
}

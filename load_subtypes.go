package main

import (
	"encoding/json"
	"io/ioutil"
)

func LoadSubtypes(filename string) (map[string]Subtype, error) {
	type file struct {
		Subtypes map[string]Subtype
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

	for name, subtype := range fileSubtypes.Subtypes {
        subtype.Name = name
        fileSubtypes.Subtypes[name] = subtype
	}

	return fileSubtypes.Subtypes, nil
}

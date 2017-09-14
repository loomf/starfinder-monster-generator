package main

import (
	"encoding/json"
	"io/ioutil"
)

func LoadSkills(filename string) ([]string, error) {
	type file struct {
		Skills []string
	}

	var fileSkills file

	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &fileSkills)
	if err != nil {
		return nil, err
	}

	return fileSkills.Skills, nil
}

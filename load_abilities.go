package main

import (
	"encoding/json"
	"io/ioutil"
)

func LoadAbilities(filename string) ([]Ability, error) {
	// TODO: differentiate between free and non-free abilities
	type file struct {
		Abilities []Ability
	}

	var fileAbilities file

	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &fileAbilities)
	if err != nil {
		return nil, err
	}

	return fileAbilities.Abilities, nil
}

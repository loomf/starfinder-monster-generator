package main

import (
	"encoding/json"
	"io/ioutil"
)

func LoadArrays(filename string) (map[string](map[string]Array), error) {
	arrayMap := make(map[string](map[string]Array))

	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &arrayMap)
	if err != nil {
		return nil, err
	}

	for name, list := range arrayMap {
		for cr, array := range list {
			array.Name = name
			array.CR = cr
			arrayMap[name][cr] = array
		}
	}

	return arrayMap, nil
}

package main

import (
	"os"
	"encoding/json"

	// "github.com/davecgh/go-spew/spew"
)

func main() {
	arrays, err := LoadArrays("arrays.json", "attack_arrays.json")
	if err != nil {
		panic(err)
	}
	types, err := LoadTypes("types.json")
	if err != nil {
		panic(err)
	}
	subtypes, err := LoadSubtypes("subtypes.json")
	if err != nil {
		panic(err)
	}
	skills, err := LoadSkills("skills.json")
	if err != nil {
		panic(err)
	}
	abilities, err := LoadAbilities("abilities.json")
	if err != nil {
		panic(err)
	}
	// spew.Dump(arrays)
	builder := CreatureBuilder{}
	builder.GetArrayType(arrays)
	builder.GetType(types)
	builder.GetSubtype(subtypes)
	// spew.Dump(builder)
	creature := builder.Build(skills, abilities)
	//spew.Dump(creature)
	encoder := json.NewEncoder(os.Stderr)
	err = encoder.Encode(creature)
	if err != nil {
		panic(err)
	}
}

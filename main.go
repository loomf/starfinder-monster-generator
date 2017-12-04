package main

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
	// "github.com/davecgh/go-spew/spew"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetOutput(os.Stderr)
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
	creature := Creature{}
    err = creature.Complete(arrays, types, subtypes, skills, abilities)
	fmt.Printf("%s\n", creature)
	fmt.Println(err)
	// builder := CreatureBuilder{}
	// builder.GetArrayType(arrays)
	// builder.GetType(types)
	// builder.GetSubtype(subtypes)
	// spew.Dump(builder)
	// creature := builder.Build(skills, abilities)
	//spew.Dump(creature)
	encoder := json.NewEncoder(os.Stdout)
    encoder.SetIndent("", "  ")
	err = encoder.Encode(creature)
	if err != nil {
		panic(err)
	}

    statBlock, err := creature.GenerateStatBlock()
	if err != nil {
		panic(err)
	}

    err = encoder.Encode(statBlock)
	if err != nil {
		panic(err)
	}
}

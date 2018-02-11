package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
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
	creature := Creature{
		TypeSpec: TypeSpec{},
	}
	err = creature.Complete(arrays, types, subtypes, skills)
	//fmt.Printf("%s\n", creature)
	log.Println(err)
	// builder := CreatureBuilder{}
	// builder.GetArrayType(arrays)
	// builder.GetType(types)
	// builder.GetSubtype(subtypes)
	// spew.Dump(builder)
	// creature := builder.Build(skills, abilities)
	//spew.Dump(creature)
	debugEncoder := json.NewEncoder(os.Stderr)
	debugEncoder.SetIndent("", "  ")
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err = debugEncoder.Encode(creature)
	if err != nil {
		panic(err)
	}

	statBlock, err := creature.GenerateStatBlock(abilities)
	if err != nil {
		panic(err)
	}

	err = encoder.Encode(statBlock)
	if err != nil {
		panic(err)
	}
}

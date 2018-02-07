package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
)

func Keys(hashMap interface{}) ([]string, error) {
	v := reflect.ValueOf(hashMap)
	if v.Kind() != reflect.Map {
		return nil, fmt.Errorf("given value %s was of type %T, not map: cannot get random key", hashMap, hashMap)
	}

	keys := v.MapKeys()
	stringKeys := make([]string, len(keys))
	for i, key := range keys {
		if key.Kind() != reflect.String {
			return nil, fmt.Errorf("key %s was %s, not string", key, key.Type())
		}
		stringKeys[i] = key.String()
	}

	return stringKeys, nil
}

func RandomKey(hashMap interface{}) (string, error) {
	keys, err := Keys(hashMap)
	if err != nil {
		return "", err
	}

	if len(keys) == 0 {
		return "", fmt.Errorf("map contained no keys")
	}

	return keys[rand.Intn(len(keys))], nil
}

func (this *Creature) Complete(arrays map[string]map[string]Array, types map[string]Type, subtypes map[string]Subtype, skills []string, abilities []Ability) error {
	var err error
	err = this.CompleteArrayType(arrays)
	if err != nil {
		return err
	}
	err = this.CompleteCR()
	if err != nil {
		return err
	}
	err = this.CompleteType(types, subtypes)
	if err != nil {
		return err
	}
	err = this.CompleteSubtype(types, subtypes)
	if err != nil {
		return err
	}
	err = this.CompleteModifierAssignments()
	if err != nil {
		return err
	}
	return nil
}

func (this *Creature) CompleteArrayType(validArrays map[string]map[string]Array) error {
	if this.ArrayType == "" {
		randomArrayType, err := RandomKey(validArrays)
		if err != nil {
			return err
		}
		log.Printf("Assinging array type %q\n", randomArrayType)
		this.ArrayType = randomArrayType
	}
	arrayTypeArrays, validArrayType := validArrays[this.ArrayType]
	if !validArrayType {
		return fmt.Errorf("Unknown array type %q\n", this.ArrayType)
	}

	this.CRArrays = arrayTypeArrays
	return nil
}

func (this *Creature) CompleteCR() error {
	if this.CR == "" {
		randomCR, err := RandomKey(this.CRArrays)
		if err != nil {
			panic(err)
		}
		log.Printf("Assinging CR %q\n", randomCR)
		this.CR = randomCR
	}
	array, validCR := this.CRArrays[this.CR]
	if !validCR {
		return fmt.Errorf("Unknown CR %q\n", this.CR)
	}
	this.Array = &array
	return nil
}

func (this *Creature) CompleteType(validTypes map[string]Type, validSubtypes map[string]Subtype) error {
	if this.Type == nil {
		var err error
		var validTypeChoices []string
		if this.Subtype == nil {
			validTypeChoices, err = Keys(validTypes)
			if err != nil {
				panic(err)
			}
		} else {
			var typeChoices []string
			if len(this.Subtype.ValidTypes) == 0 {
				typeChoices, err = Keys(validTypes)
				if err != nil {
					panic(err)
				}
			} else {
				typeChoices = this.Subtype.ValidTypes
			}
			validTypeChoices = make([]string, 0)
			for _, choice := range typeChoices {
				ctype := validTypes[choice]
				typeValidForSubtype := len(ctype.ValidSubtypes) == 0
				for _, typeCompatibleSubtype := range ctype.ValidSubtypes {
					if this.Subtype.Name == typeCompatibleSubtype {
						typeValidForSubtype = true
						break
					}
				}
				if typeValidForSubtype {
					validTypeChoices = append(validTypeChoices, choice)
				}
			}
		}
		randomValidType := validTypes[validTypeChoices[rand.Intn(len(validTypeChoices))]]
		log.Printf("Assinging creature type %q\n", randomValidType.Name)
		this.Type = &randomValidType
	}

	return nil
}

func (this *Creature) CompleteSubtype(validTypes map[string]Type, validSubtypes map[string]Subtype) error {
	if this.Subtype == nil {
		var err error
		var validSubtypeChoices []string
		if this.Type == nil {
			validSubtypeChoices, err = Keys(validSubtypes)
			if err != nil {
				panic(err)
			}
		} else {
			var typeChoices []string
			if len(this.Type.ValidSubtypes) == 0 {
				typeChoices, err = Keys(validSubtypes)
				if err != nil {
					panic(err)
				}
			} else {
				typeChoices = this.Type.ValidSubtypes
			}
			validSubtypeChoices = make([]string, 0)
			for _, choice := range typeChoices {
				ctype := validSubtypes[choice]
				typeValidForType := len(ctype.ValidTypes) == 0
				for _, typeCompatibleType := range ctype.ValidTypes {
					if this.Type.Name == typeCompatibleType {
						typeValidForType = true
						break
					}
				}
				if typeValidForType {
					validSubtypeChoices = append(validSubtypeChoices, choice)
				}
			}
		}
		randomValidSubtype := validSubtypes[validSubtypeChoices[rand.Intn(len(validSubtypeChoices))]]
		log.Printf("Assinging creature type %q\n", randomValidSubtype.Name)
		this.Subtype = &randomValidSubtype
	}

	return nil
}

func (this *Creature) CompleteModifierAssignments() error {
	if this.ModifierAssignments == [6]int{0, 0, 0, 0, 0, 0} {

		shuffle := func(list []int) {
			if len(list) == 0 {
				return
			}
			saved := list[0]
			i := 0
			perm := rand.Perm(len(list))
			for {
				if perm[i] == 0 {
					list[i] = saved
					break
				} else {
					list[i] = list[perm[i]]
					i = perm[i]
				}
			}
		}

		switch this.ArrayType {
		case "Combatant":
			// best 2 stats go to two of STR, DEX, CON
			this.ModifierAssignments = [6]int{0, 1, 2, 3, 4, 5}
			shuffle(this.ModifierAssignments[:3])
			// last 4 stats are anything
			shuffle(this.ModifierAssignments[2:])
		case "Spellcaster":
			// best stat is one of INT, WIS, CHA
			this.ModifierAssignments = [6]int{3, 4, 5, 0, 1, 2}
			shuffle(this.ModifierAssignments[:3])
			// last 5 stats are anything
			shuffle(this.ModifierAssignments[1:])
		case "Expert":
			// fully random
			this.ModifierAssignments = [6]int{0, 1, 2, 3, 4, 5}
			shuffle(this.ModifierAssignments[:])
		}
	}

	fmt.Println("assignments: %s\n", this.ModifierAssignments)

	return nil
}

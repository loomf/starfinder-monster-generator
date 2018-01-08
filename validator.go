package main

import (
	"fmt"
	"reflect"
	"math/rand"
	"log"
	"time"

	"github.com/go-playground/validator"
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

func NewValidator(arrays map[string]map[string]Array, types map[string]Type, subtypes map[string]Subtype, skills []string, abilities []Ability) *validator.Validate {
	rand.Seed(time.Now().UnixNano())
	validate := validator.New()
	validate.RegisterStructValidation(validateArraySpec(arrays), ArraySpec{})
	validate.RegisterStructValidation(validateTypeSpec(types, subtypes), TypeSpec{})
	return validate
}


func validateArraySpec(validArrays map[string]map[string]Array) validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		arraySpec := sl.Current().Interface().(ArraySpec)
		if arraySpec.ArrayType == "" {
			randomArrayType, err := RandomKey(validArrays)
			if err != nil {
				log.Println(err)
				sl.ReportError(arraySpec.ArrayType, "ArrayType", "arraytype", "arrayconfig", "")
				return
			}
			log.Printf("Assinging array type %q\n", randomArrayType)
			arraySpec.ArrayType = randomArrayType
		}
		arrayTypeArrays, validArrayType := validArrays[arraySpec.ArrayType]
		if !validArrayType {
			log.Printf("Unknown array type %q\n", arraySpec.ArrayType)
			sl.ReportError(arraySpec.ArrayType, "ArrayType", "arraytype", "arrayconfig", "")
			return
		}
		if arraySpec.CR == "" {
			randomCR, err := RandomKey(arrayTypeArrays)
			if err != nil {
				log.Println(err)
				sl.ReportError(arraySpec.CR, "CR", "cr", "arrayconfig", "")
				return
			}
			log.Printf("Assinging CR %q\n", randomCR)
			arraySpec.CR = randomCR
		}
		array, validCR := arrayTypeArrays[arraySpec.CR]
		if !validCR {
			log.Printf("Unknown CR %q\n", arraySpec.CR)
			sl.ReportError(arraySpec.CR, "CR", "cr", "arrayconfig", "")
			return
		}
        arraySpec.Array = &array
		if arraySpec != sl.Current().Interface().(ArraySpec) {
			if sl.Current().CanSet() {
				sl.Current().Set(reflect.ValueOf(arraySpec))
			} else {
				log.Println("Cannot set ArraySpec: perhaps you should trying validating a pointer type?")
				sl.ReportError(arraySpec, "ArraySpec", "settable", "arrayconfig", "")
			}
		}
	}
}

func validateTypeSpec(validTypes map[string]Type, validSubtypes map[string]Subtype) validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		typeSpec := sl.Current().Interface().(TypeSpec)

		// case both empty
			// choose random type
			// choose random valid subtype for type
		// case empty subtype 
			// choose random valid subtype for type

        var err error
		if typeSpec.Subtype == nil {
			if typeSpec.Type == nil {
				randomKey, err := RandomKey(validTypes)
                if err != nil {
                    panic(err)
                    return
                }
				randomType := validTypes[randomKey]
                log.Printf("Assinging creature type %q\n", randomType.Name)
				typeSpec.Type = &randomType
			}
            var subtypeChoices []string
            if len(typeSpec.Type.ValidSubtypes) == 0 {
                subtypeChoices, err = Keys(validSubtypes)
                if err != nil {
                    panic(err)
                    return
                }
            } else {
                subtypeChoices = typeSpec.Type.ValidSubtypes
            }
            validSubtypeChoices := make([]string, 0)
            for _, choice := range subtypeChoices {
                subtype := validSubtypes[choice]
                subtypeValidForType := len(subtype.ValidTypes) == 0
                for _, subtypeCompatibleType := range subtype.ValidTypes {
                    if typeSpec.Type.Name == subtypeCompatibleType {
                        subtypeValidForType = true
                        break
                    }
                }
                if subtypeValidForType {
                    validSubtypeChoices = append(validSubtypeChoices, choice)
                }
            }
            randomValidSubtype := validSubtypes[validSubtypeChoices[rand.Intn(len(validSubtypeChoices))]]
            log.Printf("Assinging creature subtype %q\n", randomValidSubtype.Name)
            typeSpec.Subtype = &randomValidSubtype
		}

		// case empty type
			// choose random valid type for subtype

        if typeSpec.Type == nil {
            var typeChoices []string
            if len(typeSpec.Subtype.ValidTypes) == 0 {
                typeChoices, err = Keys(validTypes)
                if err != nil {
                    panic(err)
                    return
                }
            } else {
                typeChoices = typeSpec.Subtype.ValidTypes
            }
            validTypeChoices := make([]string, 0)
            for _, choice := range typeChoices {
                ctype := validTypes[choice]
                typeValidForSubtype := len(ctype.ValidSubtypes) == 0
                for _, typeCompatibleSubtype := range ctype.ValidSubtypes {
                    if typeSpec.Subtype.Name == typeCompatibleSubtype {
                        typeValidForSubtype = true
                        break
                    }
                }
                if typeValidForSubtype {
                    validTypeChoices = append(validTypeChoices, choice)
                }
            }
            randomValidType := validTypes[validTypeChoices[rand.Intn(len(validTypeChoices))]]
            log.Printf("Assinging creature type %q\n", randomValidType.Name)
            typeSpec.Type = &randomValidType
        }

        // check type is valid
        typeValid := len(typeSpec.Subtype.ValidTypes) == 0
        for _, validTypeName := range typeSpec.Subtype.ValidTypes {
            if validTypeName == typeSpec.Type.Name {
                typeValid = true
            }
        }

        if !typeValid {
            log.Printf("Creature with subtype %q cannot have type %q\n", typeSpec.Subtype.Name, typeSpec.Type.Name)
            sl.ReportError(typeSpec.Type, "TypeSpec.Type", "validForSubtype", "typespec", "")
        }

        // check subtype is valid
        subtypeValid := len(typeSpec.Type.ValidSubtypes) == 0
        for _, validSubtypeName := range typeSpec.Type.ValidSubtypes {
            if validSubtypeName == typeSpec.Subtype.Name {
                subtypeValid = true
            }
        }

        if !subtypeValid {
            log.Printf("Creature with type %q cannot have subtype %q\n", typeSpec.Type.Name, typeSpec.Subtype.Name)
            sl.ReportError(typeSpec.Subtype, "TypeSpec.Subtype", "validForType", "typespec", "")
        }

        if !typeValid || !subtypeValid {
            return
        }

		if typeSpec != sl.Current().Interface().(TypeSpec) {
			if sl.Current().CanSet() {
				sl.Current().Set(reflect.ValueOf(typeSpec))
			} else {
				log.Println("Cannot set TypeSpec: perhaps you should trying validating a pointer type?")
				sl.ReportError(typeSpec, "TypeSpec", "settable", "typespec", "")
			}
		}
	}
}

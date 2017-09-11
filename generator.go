package main

import (
    "strconv"
    "math/rand"

    "github.com/c-bata/go-prompt"
)

type Generator struct {
}

func NoComplete(prompt.Document) []prompt.Suggest {
    return []prompt.Suggest{}
}

func GetOneOf(prefix string, choiceList []string) string {
	for {
		switch choice := prompt.Input(prefix, NoComplete); choice {
		case "":
			return choiceList[rand.Intn(len(choiceList))]
		default:
            for _, elem := range choiceList {
                if elem == choice {
                    return choice
                }
            }
		}
	}
}

func GetCR() float64 {
	var cr float64
	var err error
	for {
		switch crStr := prompt.Input("CR: ", NoComplete); crStr {
		case "":
			switch crInt := 1 + rand.Intn(27); crInt {
			case 26:
				cr = 1 / 3.0
			case 27:
				cr = 1 / 2.0
			default:
				cr = float64(crInt)
			}
		case "1/3":
			cr = 1 / 3.0
		case "1/2":
			cr = 1 / 2.0
		default:
            var crInt int
			crInt, err = strconv.Atoi(crStr)
            cr = float64(crInt)
		}
		if err == nil {
			return cr
		}
	}
}

func GetArrayType() string {
	// TODO: generate from map of array types
    arrayTypes := []string{"combatant", "expert", "spellcaster"}
    return GetOneOf("Array Type: ", arrayTypes)
}

func (this *Generator) GenerateArrayConfig() ArrayConfig {
	return ArrayConfig{
		CR:        GetCR(),
		ArrayType: GetArrayType(),
	}
}

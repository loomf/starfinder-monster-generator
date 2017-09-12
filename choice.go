package main

import (
    "fmt"
	"math/rand"

	"github.com/c-bata/go-prompt"
)

func NoComplete(prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func GetOneOf(prefix string, choiceList []string) string {
	for {
        fmt.Println(choiceList)
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

package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	builder := CreatureBuilder{}
	generator := Generator{}
	arrayConfig := generator.GenerateArrayConfig()
	arrayConfig.Apply(&builder)
	spew.Dump(builder)
}

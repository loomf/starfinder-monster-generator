package main

import (
)

type CreatureBuilder struct {
	Array
	Type
	Subtype
}

func (this *CreatureBuilder) GetArrayType(arrayMap map[string](map[string]Array)) {
	arrayTypes := make([]string, 0, len(arrayMap))
	for k := range arrayMap {
		arrayTypes = append(arrayTypes, k)
	}
	arrayType := GetOneOf("Array Type: ", arrayTypes)

    crMap := arrayMap[arrayType]
    crs := make([]string, 0, len(crMap))
    for k := range crMap {
        crs = append(crs, k)
    }
    cr := GetOneOf("CR: ", crs)
    this.Array = crMap[cr]
}

func (this *CreatureBuilder) GetType(types []Type) {
	typeNames := make([]string, len(types))
	typeMap := make(map[string]Type, len(types))
	for i, v := range types {
		typeNames[i] = v.Name
		typeMap[v.Name] = v
	}
	this.Type = typeMap[GetOneOf("Type: ", typeNames)]
}

func (this *CreatureBuilder) GetSubtype(subtypes []Subtype) {
	subtypeNames := make([]string, len(subtypes))
	subtypeMap := make(map[string]Subtype, len(subtypes))
	for i, v := range subtypes {
		subtypeNames[i] = v.Name
		subtypeMap[v.Name] = v
	}
	this.Subtype = subtypeMap[GetOneOf("Subtype: ", subtypeNames)]
}

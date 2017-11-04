package main

type CreatureBuilder struct {
	ArrayConfig
	TypeConfig
	SubtypeConfig

	Creature
}

type Config interface {
	Apply(*CreatureBuilder)
}

type ArrayConfig struct {
	CR int
	ArrayType string //TODO: enum this up
}

func (config ArrayConfig) Apply(builder *CreatureBuilder) {
	builder.ArrayConfig = config
	array := LookupArray(config)
	// TODO: copy StatsArray into corresponding spots in builder.Creature
}

func LookupArray(config ArrayConfig) StatsArray {
	// TODO: use CR and ArrayType
}

type TypeConfig struct {
	Name string
	Adjustments
	Abilities map[Ability]struct{}
}

func (config TypeConfig) Apply(builder *CreatureBuilder) {
	builder.TypeConfig = config
	// TODO: apply adjustment modifiers to creature
}

type SubtypeConfig struct {
	Name string
	Abilities map[Ability]struct{}
	Skills string //TODO: figure out how to handle the options here
	Speed map[string]bool
}

func (config SubtypeConfig) Apply(builder *CreatureBuilder) {
	builder.SubtypeConfig = config
	config.Adjustments.Adjust(&builder.Array)
}

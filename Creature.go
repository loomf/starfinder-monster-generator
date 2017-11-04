package main

type Creature struct {
	Array StatsArray
	Type CreatureType
	Abilities map[SpecialAbility]bool
	Class CreatureClass
	Skills map[string]int
	Spells map[Spell]int

	CR double
	XP int
	Size string
	Initiative int
	Senses map[string]bool
	EAC, KAC int
	HP int
	Fort, Reflex, Will int
	DefensiveAbilities map[string]bool
	DR string
	Immunities map[string]bool
	Speed map[string]int
	Melee string
	Ranged string
	OffensiveAbilities map[string]bool
	STR, DEX, CON, INT, WIS, CHA int
	Languages map[string]bool
}

type StatsArray struct {
}

type CreatureType struct {
	Subtype CreatureSubtype
}

type CreatureSubType struct {
}

type SpecialAbility struct {
}

type CreatureClass struct {
}

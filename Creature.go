package main

type CreatureBuilder struct {
	CR int
	ArrayType string //TODO: enum this up
	Type CreatureType
	Abilities map[SpecialAbility]bool
	Class CreatureClass

}

func (this* CreatureBuilder) getArray() StatsArray {
}
`
type Creature struct {
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
	Resistances map[string]int
	Speed map[string]int
	Melee map[Attack]bool
	Ranged map[Attack]bool
	OffensiveAbilities map[string]bool
	STR, DEX, CON, INT, WIS, CHA int
	Languages map[string]bool
	AbilityDC, BaseSpellDC int
}

type StatsArray struct {
	CR double
	EAC, KAC int
	Fort, Reflex, Will int
	HP int
	AbilityDC, BaseSpellDC int
	AbilityScoreBonuses [3]int
	MasterSkillBonus, MasterSkills int
	GoodSkillBonus, GoodSkills int
}

type CreatureType struct {
	Subtype CreatureSubtype

	Name string
	Senses map[string]bool
	Abilities map[SpecialAbility]bool
	Immunities map[string]bool
	Fort, Reflex, Will, RandomSave int
	AttackBonus int
	STR, DEX, CON, INT, WIS, CHA int

}

type CreatureSubtype struct {
	Name string
	Senses map[string]bool
	Immunities map[string]bool
	OffensiveAbilities map[OffensiveAbility]bool
	DefensiveAbilities map[DefensiveAbility]bool
	OtherAbilities map[OtherAbility]bool
	Weaknesses map[Weakness]bool
	Skills string //TODO: figure out how to handle the options here
	Speed map[string]bool
	Resistances map[string]int
}

type SpecialAbility interface {
	getDescription()
}

type OffensiveAbility struct{
}

type DefensiveAbility struct{

}

type OtheAbility struct {
}

type CreatureClass struct {
}

type Attack struct {
	Dice, DamageDie, DamageBonus int
	AttackBonus int
	DamageType string
}

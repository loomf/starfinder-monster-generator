package main

type AbilityType int

const (
	SENSE      AbilityType = iota
	IMMUNITY   AbilityType = iota
	RESISTANCE AbilityType = iota
	WEAKNESS   AbilityType = iota
	ATTACK     AbilityType = iota
	SPECIAL    AbilityType = iota
	OFFENSE    AbilityType = iota
	DEFENSE    AbilityType = iota
	OTHER      AbilityType = iota
)

type Adjustments struct {
	AttackBonus        int
	Fort, Reflex, Will int
}

type Ability struct {
	Name        string
	AbilityType string
	Format      string
}

type StatsArray struct {
	CR                             float64
	EAC, KAC                       int
	Fort, Reflex, Will             int
	HP                             int
	AbilityDC, BaseSpellDC         int
	AbilityScoreBonuses            [3]int
	MasterSkillBonus, MasterSkills int
	GoodSkillBonus, GoodSkills     int
}

type Creature struct {
	Skills                       map[string]int
    //Spells                       map[Spell]int
	CR                           float64
	XP                           int
	Size                         string
	Initiative                   int
	Senses                       map[string]bool
	EAC, KAC                     int
	HP                           int
	Fort, Reflex, Will           int
	DefensiveAbilities           map[string]bool
	DR                           string
	Immunities                   map[string]bool
	Resistances                  map[string]int
	Speed                        map[string]int
	Melee                        map[Attack]bool
	Ranged                       map[Attack]bool
	OffensiveAbilities           map[string]bool
	STR, DEX, CON, INT, WIS, CHA int
	Languages                    map[string]bool
	AbilityDC, BaseSpellDC       int
}
type Attack struct {
	Dice, DamageDie, DamageBonus int
	AttackBonus                  int
	DamageType                   string
}

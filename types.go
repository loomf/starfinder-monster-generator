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

//	{	"name": "Aberration",
//		"abilities": ["Darkvision"],
//		"adjustments": {
//			"intelligence": [-4,-5],
//			"attackBonus": 0,
//			"fort": 0,
//			"reflex": 0,
//			"will": 2
//		}
//	},

type Type struct {
	Name        string
	Adjustments Adjustments
	Abilities   []string
}

type Adjustments struct {
	AttackBonus        int
	Fort, Reflex, Will int
}

//	{	"name": "Aquatic",
//		"abilities": ["Water Breathing", "Amphibious"],
//		"skills": [{"Athletics":"master"}],
//		"speed": ["Swim"]
//	},

type Subtype struct {
	Name      string
	Abilities []string
	Skills    []map[string]string
	Speed     []string
}

type Ability struct {
	Name        string
	AbilityType string
	Format      string
}

//  "Combatant": {
//     "1": {
//        "EAC": 11,
//        "KAC": 13,
//        "FORT": 3,
//        "REF": 3,
//        "WILL": 1,
//        "HP": 20,
//        "ABILITY DC": 10,
//        "BASE SPELL DC": 9,
//        "ABILITY SCORE BONUSES": [4, 2, 1],
//        "SPECIAL ABILITIES": 1,
//        "MASTER SKILL BONUS": 10,
//        "MASTER SKILLS": 1,
//        "GOOD SKILL BONUS": 5,
//        "GOOD SKILLS": 2
//     },

type Array struct {
	Name                string
	CR                  string
	EAC, KAC            int
	Fort, Reflex, Will  int
	HP                  int
	AbilityDC           int    `json:"ABILITY DC"`
	BaseSpellDC         int    `json:"BASE SPELL DC"`
	AbilityScoreBonuses [3]int `json:"ABILITY SCORE BONUSES"`
	MasterSkillBonus    int    `json:"MASTER SKILL BONUS"`
	MasterSkills        int    `json:"MASTER SKILLS"`
	GoodSkillBonus      int    `json:"GOOD SKILL BONUS"`
	GoodSkills          int    `json:"GOOD SKILLS"`
}

type Creature struct {
	Skills map[string]int
	//Spells                       map[Spell]int
	CR                           string
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

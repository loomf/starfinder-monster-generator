package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

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

type Type struct {
	Name          string
	Adjustments   Adjustments
	Abilities     []string
	ValidSubtypes []string `json:"-"`
}

type Adjustments struct {
	AttackBonus        int
	Fort, Reflex, Will int
}

type Subtype struct {
	Name       string
	Abilities  []string
	Skills     map[string]string
	Speed      []string
	ValidTypes []string `json:"-"`
}

type Ability struct {
	Name        string
	Type        string
	Format      string
	Description string
}

type Array struct {
	Name                string
	CR                  string
	EAC, KAC            int
	Fort, Reflex, Will  int
	HP                  int
	AbilityDC           int    `json:"ABILITY DC"`
	BaseSpellDC         int    `json:"BASE SPELL DC"`
	AbilityScoreBonuses [3]int `json:"ABILITY SCORE BONUSES"`
	SpecialAbilities    int    `json:"SPECIAL ABILITIES"`
	MasterSkillBonus    int    `json:"MASTER SKILL BONUS"`
	MasterSkills        int    `json:"MASTER SKILLS"`
	GoodSkillBonus      int    `json:"GOOD SKILL BONUS"`
	GoodSkills          int    `json:"GOOD SKILLS"`

	AttackArray
}

type AttackArray struct {
	High, Low                 int
	Energy, Kinetic, Standard Dice
}

type Dice struct {
	Num, Size int
}

func (this *Dice) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	badPieces := strings.Split(unquoted, "+")
	unquoted = badPieces[0]
	pieces := strings.Split(unquoted, "d")
	if len(pieces) != 2 {
		return fmt.Errorf("Invalid Dice spec %s, dice must be specified in the form XdY", unquoted)
	}
	num, err := strconv.Atoi(pieces[0])
	if err != nil {
		return fmt.Errorf("Invalid Dice spec %s, dice must be specified in the form XdY", unquoted)
	}
	size, err := strconv.Atoi(pieces[1])
	if err != nil {
		return fmt.Errorf("Invalid Dice spec %s, dice must be specified in the form XdY", unquoted)
	}

	this.Num = num
	this.Size = size

	return nil
}

type StatBlock struct {
	CR                 string
	XP                 int
	Size               string
	Initiative         int
	EAC, KAC           int
	HP                 int
	Fort, Reflex, Will int
	AbilityDC, BaseSpellDC       int
	STR, DEX, CON, INT, WIS, CHA int
	Languages                    map[string]struct{}
	Skills                       map[string]int
	Melee                        []Attack
	Ranged                       []Attack
	Speed                        map[string]int
    Abilities                    map[string]map[string]Ability
	DR                           string
	//Spells                       map[Spell]int
}

func (statBlock *StatBlock) AddAbilities(abilityNames []string, abilities map[string]Ability) error {
    if statBlock.Abilities == nil {
        statBlock.Abilities = make(map[string]map[string]Ability)
    }
    for _, abilityName := range abilityNames {
        ability, realAbilityName := abilities[abilityName]
        if !realAbilityName {
            return fmt.Errorf("Unknown ability: %s\n", realAbilityName)
        }
        if _, ok := statBlock.Abilities[ability.Type]; !ok {
            statBlock.Abilities[ability.Type] = make(map[string]Ability)
        }
        statBlock.Abilities[ability.Type][ability.Name] = ability
        log.Printf("abilities: %s\n", statBlock.Abilities)
    }
    return nil
}

type Attack struct {
	DamageDice  Dice
	AttackBonus int
	DamageType  string
}

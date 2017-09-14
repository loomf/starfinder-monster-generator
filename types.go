package main

import (
	"fmt"
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
	Name        string
	Adjustments Adjustments
	Abilities   []string
}

type Adjustments struct {
	AttackBonus        int
	Fort, Reflex, Will int
}

type Subtype struct {
	Name      string
	Abilities []string
	Skills    []map[string]string
	Speed     []string
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
}

type Creature struct {
	CR                 string
	XP                 int
	Size               string
	Initiative         int
	EAC, KAC           int
	HP                 int
	Fort, Reflex, Will int
	DR                 string
	Skills             map[string]int
	//Spells                       map[Spell]int
	Speed                        map[string]int
	Senses                       map[string]struct{}
	Immunities                   map[string]struct{}
	Resistances                  map[string]struct{}
	Weaknesses                   map[string]struct{}
	Specials                     map[string]struct{}
	Melee                        map[Attack]struct{}
	Ranged                       map[Attack]struct{}
	OffensiveAbilities           map[string]struct{}
	DefensiveAbilities           map[string]struct{}
	SpecialAbilities             map[string]struct{}
	OtherAbilities               map[string]struct{}
	STR, DEX, CON, INT, WIS, CHA int
	Languages                    map[string]bool
	AbilityDC, BaseSpellDC       int
}

func (this *CreatureBuilder) Build(skills []string, abilities []Ability) Creature {
	creature := Creature{
		CR:          this.Array.CR,
		EAC:         this.Array.EAC,
		KAC:         this.Array.KAC,
		Fort:        this.Array.Fort,
		Reflex:      this.Array.Reflex,
		Will:        this.Array.Will,
		HP:          this.Array.HP,
		AbilityDC:   this.Array.AbilityDC,
		BaseSpellDC: this.Array.BaseSpellDC,
	}

	// apply adjustments from type
	creature.Fort += this.Type.Adjustments.Fort
	creature.Reflex += this.Type.Adjustments.Reflex
	creature.Will += this.Type.Adjustments.Will

	modifiers := this.Array.AbilityScoreBonuses[:]
	// 3 not as good ability scores
	modifiers = append(modifiers, int((float64(modifiers[2])*0.95 - 0.25)))
	modifiers = append(modifiers, int((float64(modifiers[3])*0.85 - 0.75)))
	modifiers = append(modifiers, int((float64(modifiers[4])*0.6 - 2.0)))

	// determine creature's ability scores
	switch this.Array.Name {
	case "Combatant":
		creature.AssignAbilityScores(modifiers, []string{"STR", "DEX", "CON"}, 2)
	case "Expert":
		creature.AssignAbilityScores(modifiers, []string{"STR", "DEX", "CON", "INT", "WIS", "CHA"}, 6)
	case "Spellcaster":
		creature.AssignAbilityScores(modifiers, []string{"INT", "WIS", "CHA"}, 1)
	default:
		panic("Unknown array: " + this.Array.Name)
	}

	// determine creatures skills
	creature.AssignSkills(skills, this.MasterSkills, this.MasterSkillBonus, this.GoodSkills, this.GoodSkillBonus)

	// determine creature abilities
	creature.AssignAbilities(abilities, this.SpecialAbilities)

	return creature
}

func (this *Creature) AssignAbilities(abilities []Ability, numAbilities int) {
	this.Senses = make(map[string]struct{})
	this.Immunities = make(map[string]struct{})
	this.Resistances = make(map[string]struct{})
	this.OffensiveAbilities = make(map[string]struct{})
	this.DefensiveAbilities = make(map[string]struct{})
	this.SpecialAbilities = make(map[string]struct{})
	this.OtherAbilities = make(map[string]struct{})
	abilMap := make(map[string]Ability)
	for _, ability := range abilities {
		abilMap[ability.Name] = ability
	}

	for i := 0; i < numAbilities; i++ {
		abilList := make([]string, 0, len(abilMap))
		for ability := range abilMap {
			abilList = append(abilList, ability)
		}
		ability := abilMap[GetOneOf("Choose an ability: ", abilList)]
		switch ability.Type {
		case "SENSE":
			this.Senses[ability.Name] = struct{}{}
		case "IMMUNITY":
			this.Immunities[ability.Name] = struct{}{}
		case "RESIST":
			this.Resistances[ability.Name] = struct{}{}
		case "OFFENSE":
			this.OffensiveAbilities[ability.Name] = struct{}{}
		case "DEFENSE":
			this.DefensiveAbilities[ability.Name] = struct{}{}
		case "SPECIAL":
			this.SpecialAbilities[ability.Name] = struct{}{}
		case "OTHER":
			this.OtherAbilities[ability.Name] = struct{}{}
		default:
			panic(fmt.Sprintf("Unknown ability type: %s for ability %s", ability.Type, ability.Name))
		}
		delete(abilMap, ability.Name)
	}
}

func (this *Creature) AssignSkills(skills []string, masterSkills, masterBonus, goodSkills, goodBonus int) {
	this.Skills = make(map[string]int)
	skillMap := make(map[string]struct{})
	for _, skill := range skills {
		skillMap[skill] = struct{}{}
	}

	assignSkills := func(desc string, numSkills, skillBonus int) {
		for i := 0; i < numSkills; i++ {
			skillList := make([]string, 0, len(skillMap))
			for skill := range skillMap {
				skillList = append(skillList, skill)
			}
			skill := GetOneOf("Choose a "+desc+" skill: ", skillList)
			this.Skills[skill] = skillBonus
			delete(skillMap, skill)
		}
	}
	assignSkills("master", masterSkills, masterBonus)
	if _, ok := this.Skills["Perception"]; ok {
		// Perception was chosen as a master skill
		// grant an extra good skill
		goodSkills++
	} else {
		// Perception was not chosen as a master skill
		// make it a good skill
		this.Skills["Perception"] = goodBonus
		delete(skillMap, "Perception")
	}
	assignSkills("good", goodSkills, goodBonus)
}

func (this *Creature) AssignAbilityScores(scores []int, primaryChoices []string, numPrimaries int) {
	secondaryAbilMap := map[string]struct{}{
		"STR": {},
		"DEX": {},
		"CON": {},
		"INT": {},
		"WIS": {},
		"CHA": {},
	}
	abilMap := make(map[string]struct{})
	for _, primary := range primaryChoices {
		abilMap[primary] = struct{}{}
		delete(secondaryAbilMap, primary)
	}
	for i, score := range scores {
		abilList := make([]string, 0, len(abilMap))
		for ability := range abilMap {
			abilList = append(abilList, ability)
		}
		ability := GetOneOf(fmt.Sprintf("Ability for modifier (%d) : ", score), abilList)
		switch ability {
		case "STR":
			this.STR = score
		case "DEX":
			this.DEX = score
		case "CON":
			this.CON = score
		case "INT":
			this.INT = score
		case "WIS":
			this.WIS = score
		case "CHA":
			this.CHA = score
		}
		delete(abilMap, ability)
		if i == numPrimaries-1 {
			for ability := range secondaryAbilMap {
				abilMap[ability] = struct{}{}
			}
		}
	}
}

type Attack struct {
	Dice, DamageDie, DamageBonus int
	AttackBonus                  int
	DamageType                   string
}

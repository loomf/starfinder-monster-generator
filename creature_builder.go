package main

import (
//"fmt"
)

type Creature struct {
	ArraySpec
	TypeSpec
	ModifierAssignments [6]int `validate:"modiferAssignments"`
	GoodSkills          []string
	MasterSkills        []string
}

type ArraySpec struct {
	ArrayType string
	CR        string
	CRArrays  map[string]Array `json:"-"`
	Array     *Array           `json:"-"`
}

type TypeSpec struct {
	Type    *Type
	Subtype *Subtype
}

func (this *Creature) GenerateStatBlock() (StatBlock, error) {
	var statBlock StatBlock

	// Step 1: Array
	//
	// 1a: CR
	statBlock.CR = this.Array.CR
	// 1b: EAC, KAC, Saving Throw Bonuses
	statBlock.KAC = this.Array.KAC
	statBlock.EAC = this.Array.EAC
	statBlock.Fort = this.Array.Fort
	statBlock.Reflex = this.Array.Reflex
	statBlock.Will = this.Array.Will
	// 1c: Hit Points
	statBlock.HP = this.Array.HP
	// 1d: Ability and Spell DCs
	statBlock.AbilityDC = this.Array.AbilityDC
	statBlock.BaseSpellDC = this.Array.BaseSpellDC
	// 1e: Ability Score Modifiers
	modifiers := this.Array.AbilityScoreBonuses[:]
	modifiers = append(modifiers, int((float64(modifiers[2])*0.95 - 0.25)))
	modifiers = append(modifiers, int((float64(modifiers[3])*0.85 - 0.75)))
	modifiers = append(modifiers, int((float64(modifiers[4])*0.6 - 2.0)))
	statBlock.STR = modifiers[this.ModifierAssignments[0]]
	statBlock.DEX = modifiers[this.ModifierAssignments[1]]
	statBlock.CON = modifiers[this.ModifierAssignments[2]]
	statBlock.INT = modifiers[this.ModifierAssignments[3]]
	statBlock.WIS = modifiers[this.ModifierAssignments[4]]
	statBlock.CHA = modifiers[this.ModifierAssignments[5]]
	// 1f: Special Abilities
	// See Step 6
	// 1g: Skills
	statBlock.Skills = make(map[string]int)
	for _, skill := range this.GoodSkills {
		statBlock.Skills[skill] = this.Array.GoodSkillBonus
	}
	for _, skill := range this.MasterSkills {
		statBlock.Skills[skill] = this.Array.MasterSkillBonus
	}
	// 1h: Attack Bonuses
	// 1i: Ranged Damange
	// 1j: Melee Damage
	// Other Statistics
	// 1k: Initiative
	// 1l: Speed
	// 1l: Feats
	// 1m: Languages

	// Step 2: Creature Type Graft

	return statBlock, nil
}

/*

func (this *Creature) AssignAttacks(attackArray AttackArray, bonus int) {
	this.Melee = make([]Attack, 0)
	this.Ranged = make([]Attack, 0)
	attackMap := map[string]struct{}{
		"Melee":  {},
		"Ranged": {},
	}

	assignAttack := func(attackName string, attackBonus int) {
		attackList := make([]string, 0, len(attackMap))
		for attack := range attackMap {
			attackList = append(attackList, attack)
		}
		attackType := GetOneOf(attackName+" attack: ", attackList)
		switch attackType {
		case "Melee":
			attack := Attack{
				AttackBonus: attackBonus + bonus,
				DamageDice:  attackArray.Standard,
				DamageType:  "Kinetic",
			}
			this.Melee = append(this.Melee, attack)
		case "Ranged":
			var damageDice Dice
			damageType := GetOneOf("Damage type: ", []string{"Kinetic", "Energy"})
			switch damageType {
			case "Kinetic":
				damageDice = attackArray.Kinetic
			case "Energy":
				damageDice = attackArray.Energy
			}
			attack := Attack{
				AttackBonus: attackBonus + bonus,
				DamageDice:  damageDice,
				DamageType:  damageType,
			}
			this.Ranged = append(this.Ranged, attack)
		}
		delete(attackMap, attackType)
	}
	assignAttack("Primary", attackArray.High)
	attackMap["None"] = struct{}{}
	assignAttack("Secondary", attackArray.Low)
}

func (this *Creature) AssignAbilities(abilities []Ability, extraAbilities []string, numAbilities int) {
	this.Senses = make(map[string]struct{})
	this.Immunities = make(map[string]struct{})
	this.Resistances = make(map[string]struct{})
	this.Weaknesses = make(map[string]struct{})
	this.OffensiveAbilities = make(map[string]struct{})
	this.DefensiveAbilities = make(map[string]struct{})
	this.SpecialAbilities = make(map[string]struct{})
	this.OtherAbilities = make(map[string]struct{})
	abilMap := make(map[string]Ability)
	for _, ability := range abilities {
		abilMap[ability.Name] = ability
	}

	assignAbility := func(ability Ability) {
		switch ability.Type {
		case "SENSE":
			this.Senses[ability.Name] = struct{}{}
		case "IMMUNITY":
			this.Immunities[ability.Name] = struct{}{}
		case "RESIST":
			this.Resistances[ability.Name] = struct{}{}
		case "WEAKNESS":
			this.Weaknesses[ability.Name] = struct{}{}
		case "OFFENSE":
			this.OffensiveAbilities[ability.Name] = struct{}{}
		case "DEFENSE":
			this.DefensiveAbilities[ability.Name] = struct{}{}
		case "SPECIAL":
			this.SpecialAbilities[ability.Name] = struct{}{}
		case "OTHER":
			this.OtherAbilities[ability.Name] = struct{}{}
		case "ATTACK":
			this.OffensiveAbilities[ability.Name] = struct{}{}
		case "":
			//Ability is unfinished
		default:
			panic(fmt.Sprintf("Unknown ability type: %s for ability %s", ability.Type, ability.Name))
		}
		delete(abilMap, ability.Name)
	}

	for _, ability := range extraAbilities {
		assignAbility(abilMap[ability])
	}

	for i := 0; i < numAbilities; i++ {
		abilList := make([]string, 0, len(abilMap))
		for ability := range abilMap {
			abilList = append(abilList, ability)
		}
		ability := abilMap[GetOneOf("Choose an ability: ", abilList)]
		assignAbility(ability)
	}
}
*/

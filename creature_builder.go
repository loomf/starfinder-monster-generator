package main

import (
	"fmt"
)

type Creature struct {
	ArraySpec
	TypeSpec
	ModifierAssignments [6]int `validate:"modiferAssignments"`
	GoodSkills          []string
	MasterSkills        []string
	Attacks             []string
	Speed               map[string]int
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

func (this *Creature) GenerateStatBlock(abilities map[string]Ability) (StatBlock, error) {
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
	for i, attack := range this.Attacks {
		var bonus int
		if i == 0 {
			bonus = this.Array.AttackArray.High
		} else {
			bonus = this.Array.AttackArray.Low
		}
		switch attack {
		case "Melee":
			statBlock.Melee = append(statBlock.Melee,
				Attack{
					AttackBonus: bonus,
					DamageDice:  this.Array.AttackArray.Standard,
					DamageType:  "Kinetic",
				},
			)
		case "Ranged-Kinetic":
			statBlock.Ranged = append(statBlock.Ranged,
				Attack{
					AttackBonus: bonus,
					DamageDice:  this.Array.AttackArray.Kinetic,
					DamageType:  "Kinetic",
				},
			)
		case "Ranged-Energy":
			statBlock.Ranged = append(statBlock.Ranged,
				Attack{
					AttackBonus: bonus,
					DamageDice:  this.Array.AttackArray.Energy,
					DamageType:  "Energy",
				},
			)
		default:
			panic(fmt.Sprintf("unknown attack %q", attack))
		}
	}
	// Other Statistics
	// 1k: Initiative
	statBlock.Initiative = statBlock.DEX
	// 1l: Speed
	statBlock.Speed = this.Speed
	// 1l: Feats
	// nothing to do
	// 1m: Languages
	// TODO

	// Step 2: Creature Type Graft
	statBlock.Type = this.Type.Name
	// 2a: Traits
	statBlock.AddAbilities(this.Type.Abilities, abilities)
	// 2b: Adjustments
	for _, attack := range statBlock.Melee {
		attack.AttackBonus += this.Type.Adjustments.AttackBonus
	}
	for _, attack := range statBlock.Ranged {
		attack.AttackBonus += this.Type.Adjustments.AttackBonus
	}
	statBlock.Fort += this.Type.Adjustments.Fort
	statBlock.Reflex += this.Type.Adjustments.Reflex
	statBlock.Will += this.Type.Adjustments.Will

	// Step 3: Creature Subtype Graft
	statBlock.Subtype = this.Subtype.Name
	statBlock.AddAbilities(this.Subtype.Abilities, abilities)
	for skillName, level := range this.Subtype.Skills {
		switch level {
		case "Good":
			statBlock.Skills[skillName] = this.Array.GoodSkillBonus
		case "Master":
			statBlock.Skills[skillName] = this.Array.MasterSkillBonus
		default:
			panic(fmt.Errorf("Unknown skill %q from subtype %q\n", skillName, this.Subtype.Name))
		}
	}

	return statBlock, nil
}

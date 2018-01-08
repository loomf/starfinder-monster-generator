package main

import (
	//"fmt"
)

type Creature struct {
	ArraySpec
	TypeSpec
}

type ArraySpec struct {
	ArrayType string
	CR string
    Array *Array
}

type TypeSpec struct {
	Type *Type
	Subtype *Subtype
}

func (this *Creature) GenerateStatBlock() (StatBlock, error) {
    var statBlock StatBlock

    // Step 1: Array
    statBlock.CR = this.Array.CR
    statBlock.KAC = this.Array.KAC
    statBlock.EAC = this.Array.EAC
    statBlock.Fort = this.Array.Fort
    statBlock.Reflex = this.Array.Reflex
    statBlock.Will = this.Array.Will
    statBlock.HP = this.Array.HP
    statBlock.AbilityDC = this.Array.AbilityDC
    statBlock.BaseSpellDC = this.Array.BaseSpellDC

    modifiers := GenerateModifiers(this.Array.Name, this.Array.AbilityScoreBonuses)
    statBlock.STR = modifiers[0]
    statBlock.DEX = modifiers[1]
    statBlock.CON = modifiers[2]
    statBlock.INT = modifiers[3]
    statBlock.WIS = modifiers[4]
    statBlock.CHA = modifiers[5]

    return statBlock, nil
}

func GenerateModifiers(arrayType string, modifiers []int) [6]int {
	modifiers = append(modifiers, int((float64(modifiers[2])*0.95 - 0.25)))
	modifiers = append(modifiers, int((float64(modifiers[3])*0.85 - 0.75)))
	modifiers = append(modifiers, int((float64(modifiers[4])*0.6 - 2.0)))

    var assignedModifiers [6]int
    var preferrendIndices map[int]struct{}
    var numPreferredModifiers int
	switch this.Array.Name {
	case "Combatant":
        preferredIndices = map[int]struct{
            0: struct{}{},
            1: struct{}{},
            2: struct{}{},
        }
        numPreferredModifiers = 2
	case "Expert":
        preferredIndices = map[int]struct{
        }
        numPreferredModifiers = 0
	case "Spellcaster":
        preferredIndices = map[int]struct{
            3: struct{}{},
            4: struct{}{},
            5: struct{}{},
        }
        numPreferredModifiers = 1
	default:
		panic("Unknown array: " + this.Array.Name)

        for i := 0; i < numPreferredModifiers; i++ {

        }
    }
}

/*

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

	// determine creature attacks
	creature.AssignAttacks(this.Array.AttackArray, this.Type.Adjustments.AttackBonus)

	// determine creature abilities
	creature.AssignAbilities(abilities, this.Subtype.Abilities, this.SpecialAbilities)

	// determine creatures skills
	creature.AssignSkills(skills, this.Subtype.Skills, this.MasterSkills, this.MasterSkillBonus, this.GoodSkills, this.GoodSkillBonus)

	return creature
}

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

func (this *Creature) AssignSkills(skills []string, extraSkills map[string]string, masterSkills, masterBonus, goodSkills, goodBonus int) {
	this.Skills = make(map[string]int)
	skillMap := make(map[string]struct{})
	for _, skill := range skills {
		skillMap[skill] = struct{}{}
	}

	for skill, quality := range extraSkills {
		switch quality {
		case "master":
			this.Skills[skill] = masterBonus
		case "good":
			this.Skills[skill] = goodBonus
		}
		delete(skillMap, skill)
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
*/

package party

import (
	"math/rand"
	"time"
	"the_brink/characters"

	"github.com/fatih/color"
  )

// TargetMember returns an attackble player from the party
func (party *Party) TargetMember() *characters.Character {

	// for all members of the party
	randomIndex := rand.Intn(len(party.Members))
	for i := 0; i < len(party.Members); i++ {
		
		
		randomPartyMember := party.Members[randomIndex]
		if randomPartyMember.Stats.Health > 0 {
			return randomPartyMember
		}
	}
	return nil
}


// Battle other party
func (party *Party) Battle(otherParty *Party) {

	round := 1
	for party.GetHealth() > 0 && otherParty.GetHealth() > 0 {
		time.Sleep(150 * time.Millisecond)
		color.Green("\n\nRound %d", round)

		for _, member := range party.Members {
			// skip turn if dead
			if member.Stats.Health <= 0 {
				continue
			}
			// find target
			target := otherParty.TargetMember()
			if target == nil {
				continue
			}

			// member action
			member.Act(target)
			
		}
		
		for _, member := range otherParty.Members {

			// skip turn if dead
			if member.Stats.Health <= 0 {
				continue
			}
			// find target
			target := party.TargetMember()
			if target == nil {
				continue
			}

			// member action
			member.Act(target)
		}
		round++
	}

	// If the party doesn't die, add XP
	if party.GetHealth() > 0 {
		for _, otherMember := range otherParty.Members {
			for _, selfMember := range party.Members {
				selfMember.Stats.XP += otherMember.Stats.XP
			}
		}
		color.Cyan("Player Party wins the battle\n")
		return
	}

	color.Cyan("Enemy Party wins the battle\n")
}
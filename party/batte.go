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
	for i := 0; i < len(party.Members); i++ {
		
		randomIndex := rand.Intn(len(party.Members))
		randomPartyMember := party.Members[ randomIndex]
		if randomPartyMember.Stats.Health > 0 {
			return randomPartyMember
		}
	}
	return nil
}


// Battle other party
func (selfParty *Party) Battle(otherParty *Party) {

	round := 1
	for selfParty.GetHealth() > 0 && otherParty.GetHealth() > 0 {
		time.Sleep(150 * time.Millisecond)
		color.Green("\n\nRound %d", round)

		for _, member := range selfParty.Members {
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
			target := selfParty.TargetMember()
			if target == nil {
				continue
			}

			// member action
			member.Act(target)
		}
		round++
	}

	if selfParty.GetHealth() > 0 {
		for _, otherMember := range otherParty.Members {
			for _, selfMember := range selfParty.Members {
				selfMember.Stats.XP += otherMember.Stats.XP
			}
		}
		color.Cyan("Player Party wins the battle\n")
		return
	}

	color.Cyan("Enemy Party wins the battle\n")
}
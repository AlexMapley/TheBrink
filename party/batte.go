package party

import (
	"math/rand"
	"the_brink/characters"
	"time"

	"github.com/fatih/color"
)

// ReturnMember returns an attackble player from the party
func (party *Party) ReturnMember() *characters.Character {

	// Max target attempts: 10
	for i := 0; i < 10; i++ {
		randomIndex := rand.Intn(len(party.Members))
		if party.Members[randomIndex].Stats.Health > 0 {
			return party.Members[randomIndex]
		}
	}
	return nil
}

// Battle other party
func (party *Party) Battle(otherParty *Party) {

	round := 1
	for party.GetHealth() > 0 && otherParty.GetHealth() > 0 {
		time.Sleep(100 * time.Millisecond)
		color.Green("\n\nRound %d", round)

		for _, member := range party.Members {
			// skip turn if dead
			if member.Stats.Health <= 0 {
				continue
			}
			// find target
			target := otherParty.ReturnMember()
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
			target := party.ReturnMember()
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

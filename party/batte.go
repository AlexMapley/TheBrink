package party

import (
	"fmt"
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
		fmt.Printf("random index: %v\n", randomIndex)

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

		fmt.Println("break 1")
		for _, member := range selfParty.Members {
			fmt.Println("break 2")
			// skip turn if dead
			if member.Stats.Health <= 0 {
				continue
			}
			// find target
			target := otherParty.TargetMember()
			fmt.Printf("Target: %s\n", target.Stats.Name)
			if target == nil {
				continue
			}

			// member action
			fmt.Println("break 3")
			member.Act(target)
			
		}
		
		for _, member := range otherParty.Members {
			fmt.Println("break 4")

			// skip turn if dead
			if member.Stats.Health <= 0 {
				continue
			}
			// find target
			target := selfParty.TargetMember()
			fmt.Printf("Target: %s\n", target.Stats.Name)
			if target == nil {
				continue
			}

			// member action
			fmt.Println("break 5")
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
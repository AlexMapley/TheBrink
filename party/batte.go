package party

import (
	"math/rand"
	"time"

	"the_brink/characters"
  )

// TargetMember returns an attackble player from the party
func (party *Party) TargetMember() *characters.Character {

	rand.Seed(time.Now().Unix())

	for {
		randomPartyMember := party.Members[rand.Intn(len(party.Members))]
		if randomPartyMember.Health > 0 {
			return randomPartyMember
		}
	}
}
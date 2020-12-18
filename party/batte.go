package party

import (
	"math/rand"
	"time"
  )

// TargetMember returns an attackble player from the party
func (party *Party) TargetMember() *Character {

	rand.Seed(time.Now().Unix())

	for {
		randomPartyMember := Party.Members[rand.Intn(len(Party.Members))]
		if randomPartyMember.Health > 0 {
			return randomPartyMember
		}
	}
}
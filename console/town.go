package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 6)
	actions[0] = "Character"
	actions[1] = "Map"
	actions[2] = "Patrol the town"
	actions[3] = "Rest"
	actions[4] = "Become Paladin"
	actions[5] = "Become NightBlade"

	console.Actions = actions

	return console
}

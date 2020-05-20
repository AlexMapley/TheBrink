package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 7)
	actions[0] = "Stats"
	actions[1] = "Skills"
	actions[2] = "Inventory"
	actions[3] = "Map"
	actions[4] = "Patrol the town"
	actions[5] = "Rest"
	actions[6] = "Level Up"
	actions[6] = "Become Paladin"

	console.Actions = actions

	return console
}

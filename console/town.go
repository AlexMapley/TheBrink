package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 6)
	actions[0] = "Stats"
	actions[1] = "Skills"
	actions[2] = "Inventory"
	actions[3] = "Patrol the town"
	actions[4] = "Rest"
	actions[5] = "Level Up"

	console.Actions = actions

	return console
}

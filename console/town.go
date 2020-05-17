package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 5)
	actions[0] = "Stats"
	actions[1] = "Inventory"
	actions[2] = "Patrol the town"
	actions[3] = "Rest"
	actions[4] = "Level Up"

	console.Actions = actions

	return console
}

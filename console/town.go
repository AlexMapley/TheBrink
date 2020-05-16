package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 5)
	actions[0] = "End the day"
	actions[1] = "Stats"
	actions[2] = "Inventory"
	actions[3] = "Patrol the town"
	actions[4] = "Rest"

	console.Actions = actions

	return console
}

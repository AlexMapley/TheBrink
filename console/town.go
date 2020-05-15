package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 4)
	actions[0] = "End the day"
	actions[1] = "Patrol the town"
	actions[2] = "Stats"
	actions[3] = "Rest"

	console.Actions = actions

	return console
}

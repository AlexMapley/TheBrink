package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 3)
	actions[0] = "End the day"
	actions[1] = "Patrol the town"
	actions[2] = "Rest"

	console.Actions = actions

	return console
}

package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make(map[int]string)
	actions[1] = "End the day"
	actions[2] = "Patrol the town"
	actions[3] = "Rest"

	console.Actions = actions

	return console
}

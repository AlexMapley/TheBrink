package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 6)
	actions[0] = "Character"
	actions[1] = "Map"
	actions[2] = "Patrole the town"
	actions[3] = "Rest"
	actions[4] = "Level Up"
	actions[5] = "x | q | esc to exit"

	console.Actions = actions

	return console
}

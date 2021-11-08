package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 6)
	actions[0] = "Character"
	actions[1] = "Map"
	actions[2] = "Defend the Town"
	actions[3] = "Level Up"
	actions[4] = "Time Warp"
	actions[5] = "x | q | esc to exit"

	console.Actions = actions

	return console
}

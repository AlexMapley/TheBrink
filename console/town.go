package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 7)
	actions[0] = "Character"
	actions[1] = "Map"
	actions[2] = "Defend the Town"
	actions[3] = "Rest"
	actions[4] = "Level Up"
	actions[5] = "Time Warp"
	actions[6] = "x | q | esc to exit"

	console.Actions = actions

	return console
}

package console

var trim string = "^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-"

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 5)
	actions[0] = "Arrow | WASD Keys for Movement"
	actions[1] = "Character"
	actions[2] = "Defend the Town"
	actions[3] = "Time Warp"
	actions[4] = "x | q | esc to exit"

	console.Actions = actions
	return console
}
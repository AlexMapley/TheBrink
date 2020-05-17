package console

func NewClassConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 3)
	actions[0] = "Rogue"
	actions[1] = "Warrior"
	actions[2] = "Wizard"

	console.Actions = actions

	return console
}

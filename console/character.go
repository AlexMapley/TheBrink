package console


func NewCharacterConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 3)
	actions[0] = "Stats"
	actions[1] = "Inventory"

	console.Actions = actions

	return console
}

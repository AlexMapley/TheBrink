package main

func NewBattleConsole() Console {
	console := Console{}

	// set default options
	actions := make(map[int]string)
	actions[1] = "End the day"
	console.Actions = actions

	return console
}

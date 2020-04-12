package main

func NewBattleConsole() Console {
	console := Console{}

	// set default options
	actions := make(map[int]string)
	actions[1] = "Attack"
	console.Actions = actions

	return console
}

package town

type BattleConsole struct {
	Console Console
}

func NewBattleConsole() BattleConsole {
	console := BattleConsole{}

	// set default options
	actions := make(map[int]string)
	actions[1] = "Attack"
	console.Actions = actions

	return console
}

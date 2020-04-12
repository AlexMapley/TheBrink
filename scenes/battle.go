package town

type BattleConsole struct {
	Console Console
}

func NewBattleConsole() BattleConsole {
	console := BattleConsole{}

	// set default options
	options := make(map[int]string)
	options[1] = "Attack"
	console.Options = options

	return console
}

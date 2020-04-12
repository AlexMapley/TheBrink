package town

type TownConsole struct {
	Console Console
}

func NewTownSoncole() TownConsole {
	console := TownConsole{}

	// set default options
	options := make(map[int]string)
	options[1] = "Attack"
	console.Options = options

	return console
}

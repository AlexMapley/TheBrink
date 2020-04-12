package town

type TownConsole struct {
	Console Console
}

func NewTownConsole() TownConsole {
	console := TownConsole{}

	// set default options
	options := make(map[int]string)
	options[1] = "End the day"
	console.Options = options

	return console
}

package main

func main() {
	cfg := config{
		commandRegistry: getCommands(),
	}
	startRepl(&cfg)
}

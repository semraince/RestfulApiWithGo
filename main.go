package main

import (
	"GoDatabaseAssignment/app"
	"GoDatabaseAssignment/config"
	"GoDatabaseAssignment/inmemory"
)

func main() {
	config.Load()
	inmemory.Init()
	inmemory.Init()
	app.Init()

}

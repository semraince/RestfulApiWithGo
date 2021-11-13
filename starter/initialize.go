package starter

import (
	"GoDatabaseAssignment/config"
	"GoDatabaseAssignment/database"
	"GoDatabaseAssignment/inmemory"
)

func Init() {
	config.Load()
	inmemory.Init()
	database.Init()
}

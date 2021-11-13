package main

import (
	"GoDatabaseAssignment/app"
	"GoDatabaseAssignment/config"
	"GoDatabaseAssignment/inmemory"
	"log"
)

func main() {
	config.Load()
	log.Println("Config sanity check: OK!")
	inmemory.Init()
	log.Println("DB sanity check: OK!")
	inmemory.Init()
	b := inmemory.KeyValue{ // b == Student{"Bob", 0}
		Key:   "Bob",
		Value: "Semra",
	}
	inmemory.Write(&b)
	inmemory.Read("Bob")
	c := inmemory.KeyValue{ // b == Student{"Bob", 0}
		Key:   "Bob",
		Value: "abc",
	}
	inmemory.Write(&c)
	inmemory.Read("Bob")
	app.Init()

}

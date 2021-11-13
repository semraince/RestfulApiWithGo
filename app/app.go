package app

import (
	"GoDatabaseAssignment/config"
	"GoDatabaseAssignment/pairs"
	"GoDatabaseAssignment/records"
	"log"
	"net/http"
	"os"
)

func Init() {
	addr := config.Get().Servers.AppServer
	envPort := os.Getenv("PORT")
	if envPort != "" {
		addr = ":" + envPort
	}
	log.Printf("App Address : %s", addr)

	http.HandleFunc("/records", records.Records)
	http.HandleFunc("/config", pairs.KeyValuePairs)
	http.ListenAndServe(addr, nil)

}

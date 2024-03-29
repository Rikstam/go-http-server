package main

import (
	"log"
	"net/http"

	poker "github.com/Rikstam/go-http-server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	defer close()
	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5001", server); err != nil {
		log.Fatalf("could not listen on port 5001 %v", err)
	}
}

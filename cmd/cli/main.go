package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/Rikstam/go-http-server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}

package main

import (
	"log"
	"github.com/uttamgupta2712/database"
	"github.com/uttamgupta2712/router"
)

func main() {
	database.Setup()                    //establish the dtatbase connection
	engine := routers.Router()          //get the customized engine
	err := engine.Run("127.0.0.1:8080") // start the engine
	if err != nil {
		log.Fatal(err)
	}
}
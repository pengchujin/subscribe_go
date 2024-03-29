package main

import (
	"github.com/pengchujin/subscribe_go/routes"
	"github.com/pengchujin/subscribe_go/database"
	"github.com/pengchujin/subscribe_go/config"
	"fmt"
)

func main() {
	if err:= config.Load("config/config.yaml"); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("err open databases")
		return
	}
	defer db.Close()
	router := routes.InitRouter()
	router.Run(config.Get().Addr)
}
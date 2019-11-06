package main

import (
	"github.com/pengchujin/subscribe/routes"
	"github.com/pengchujin/subscribe/database"
	"github.com/pengchujin/subscribe/config"
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
	router.RUN(config.Get().Addr)
}
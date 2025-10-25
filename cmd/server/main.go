package main

//Github configured

import (
	"basics/internal/config"
	"basics/internal/routes"
	"basics/internal/utils"
	"fmt"
	"log"
)

func main() {
	utils.LoadEnvConfig()

	config.ConnectDB()

	//“Okay, remember this — I’ll run it later.”
	defer config.DB.Close()

	r := routes.SetupRouter()

	fmt.Println("server starts")

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Something went worng during run the Port", err)
	}
}

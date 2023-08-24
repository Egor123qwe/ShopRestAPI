package main

import (
	"ShopRestAPI/internal/server/ServerApi"
	"log"
)

func main() {

	config := ServerApi.NewConfig()

	if err := ServerApi.Start(config); err != nil {
		log.Fatal(err)
	}
}

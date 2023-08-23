package main

import (
	"ShopRestAPI/internal/server"
	"ShopRestAPI/internal/server/ServerApi"
	"log"
)

func main() {

	config := server.NewConfig()
	s := ServerApi.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

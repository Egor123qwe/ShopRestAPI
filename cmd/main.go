package main

import (
	"ShopRestAPI/internal/ServerAPI"
	"log"
)

func main() {

	config := ServerAPI.NewConfig()
	s := ServerAPI.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

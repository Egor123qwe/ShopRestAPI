package main

import (
	"ShopRestAPI/internal/serverAPI"
	"log"
)

func main() {

	config := serverAPI.NewConfig()
	s := serverAPI.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

package ServerAPI

import (
	"ShopRestAPI/internal/Storage"
	"ShopRestAPI/internal/Storage/models"
	"fmt"
	"log"
)

type ServerApi struct {
	config *Config
	store  *Storage.Store
}

func New(config *Config) *ServerApi {
	return &ServerApi{
		config: config,
		store:  Storage.New(config.store),
	}
}

func (s *ServerApi) Start() error {

	if err := s.configureStore(); err != nil {
		log.Fatal(err)
	}

	//s.configureRoutes()

	//Тест
	s.store.Product().GetPropertyList("color")

	var filter = &models.ProductFilter{
		Term:     "",
		MinPrice: 0,
		MaxPrice: 0,
		Print:    []string{},
		Types:    []string{},
		Style:    []string{},
		Season:   []string{},
		Country:  []string{},
		Color:    []string{},
		Size:     []string{},
	}
	var products, _ = s.store.Product().ProductsSearch(10, 0, filter)
	fmt.Println(products)

	return nil
}

func (s *ServerApi) configureStore() error {
	if err := s.store.Open(); err != nil {
		return err
	}
	return nil
}

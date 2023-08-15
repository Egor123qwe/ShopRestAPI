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

	//Тест
	props := []models.Property{
		{
			Color:  "Black",
			Photos: "src/chlen",
			Size:   "M",
			Amount: 2,
		},
	}

	product := &models.Product{
		Name:        "super SweetShot",
		Price:       49.99,
		Description: "Nice SweetShot!",
		Print:       "Penis",
		Types:       "SweetShot",
		Style:       "gangsta",
		Season:      "summer",
		Country:     "USA",
		Properties:  &props,
	}

	if err := s.store.Product().Create(product); err != nil {
		log.Fatal(err)
	}
	//s.store.Product().Delete(22)
	product, _ = s.store.Product().Get(1)
	fmt.Println(product)
	fmt.Println(product.Properties)
	return nil
}

func (s *ServerApi) configureStore() error {
	if err := s.store.Open(); err != nil {
		return err
	}
	return nil
}

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
	m := &models.Product{
		ID:          1,
		Name:        "dick",
		Amount:      777,
		Description: "трусы топ",
		TypeId:      1,
		PhotosId:    2,
		Price:       3.3,
	}

	s.store.Product().Create(m)
	fmt.Println(s.store.Product().Find(1))

	return nil
}

func (s *ServerApi) configureStore() error {
	if err := s.store.Open(); err != nil {
		return err
	}
	return nil
}

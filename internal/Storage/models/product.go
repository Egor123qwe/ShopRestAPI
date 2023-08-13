package models

type Product struct {
	ID          int
	Name        string
	Amount      int
	Description string
	TypeId      int
	PhotosId    int
	Price       float32
}

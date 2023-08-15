package models

type Product struct {
	Id          int
	Name        string
	Price       float32
	Description string
	Print       string
	Types       string
	Style       string
	Season      string
	Country     string
	Properties  *[]Property
}

type Property struct {
	Id     int
	Color  string
	Photos string
	Size   string
	Amount int
}

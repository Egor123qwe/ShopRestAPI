package models

type Product struct {
	Id          int
	Name        string
	Price       float32
	Description string
	Photos      string
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

type ProductFilter struct {
	Term     string
	MinPrice float32
	MaxPrice float32
	Print    []string
	Types    []string
	Style    []string
	Season   []string
	Country  []string
	Color    []string
	Size     []string
}

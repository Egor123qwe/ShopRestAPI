package products

type Product struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Price       float32    `json:"price"`
	Description string     `json:"description"`
	Photos      string     `json:"photos"`
	Print       string     `json:"print"`
	Types       string     `json:"types"`
	Style       string     `json:"style"`
	Season      string     `json:"season"`
	Country     string     `json:"country"`
	Instances   []Instance `json:"properties"`
}

type Instance struct {
	Id        int    `json:"id"`
	ProductId int    `json:"productId"`
	Color     string `json:"color"`
	Photos    string `json:"photos"`
	Size      string `json:"size"`
	Amount    int    `json:"amount"`
}

type Filter struct {
	Page     int      `json:"page"`
	Count    int      `json:"count"`
	Term     string   `json:"term"`
	MinPrice float32  `json:"minPrice"`
	MaxPrice float32  `json:"maxPrice"`
	Print    []string `json:"print"`
	Types    []string `json:"types"`
	Style    []string `json:"style"`
	Season   []string `json:"season"`
	Country  []string `json:"country"`
	Color    []string `json:"color"`
	Size     []string `json:"size"`
}

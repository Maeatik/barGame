package models

type Drinks struct {
	Item 	 string  `json:"item" db:"item"`
	Cost 	 float64 `json:"cost" db:"cost"`
	Promille float64 `json:"promille" db:"promille"`
}

type DrinksUpdateInput struct {
	Cost 	 float64 `json:"cost" db:"cost"`
	Promille float64 `json:"promille" db:"promille"`
}

type DrinkId struct {
	Id string	`json:"id" db:"id"`
}
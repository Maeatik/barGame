package models

type User struct {
	Id			int		`json:"-" db:"id"`
	Name 		string 	`json:"name" binding:"required"`
	Password 	string 	`json:"password" binding:"required"`
	Money 		float64 	`json:"money"`
	Promille	float64 `json:"promille"`
	Status		bool 	`json:"status"`
	Role 		bool 	`json:"role"`
}

type Me struct {
	Money 		float64 	`json:"money" db:"money"`
	Promille 	float64		`json:"promille" db:"promille"`
	Status 		string		`json:"status" db:"status"`
}
type CheckStatus struct {
	Status string `json:"status" db:"status"`
}
type CheckRole struct {
	Role bool `json:"role" db:"role"`
}

type PromilleDecrease struct {
	Id 			int 	`json:"id" db:"id"`
	Promille 	float64 `json:"promille" db:"promille"`
	Status 		string	`json:"status" db:"status"`
	Role 		bool 	`json:"role" db:"role"`
}
package models

// Customer model
type Customer struct {
	PartnerId  string   `json:"partner_id"`
	Email       string   `json:"email"`
	FirstName    string   `json:"firstname"`
	LastName string   `json:"lastname"`
	Address string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
}
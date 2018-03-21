package models

// Customer used for customer data
type Customer struct {
	PartnerId   string `json:"partner_id"`
	Id          string `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
}

type CustomerError struct {
	Error string `json:"error"`
}

type CustomerFilter struct {
	PartnerId string `json:"partner_id"`
	Email     string `json:"email"`
}


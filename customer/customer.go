package customer

import "time"

// Represent the customer
// At the moment it can't be update
type Customer struct {
	ID       ID        `json:"id,string"`
	Email    Email     `json:"email,string"`
	Password Password  `json:"password,string"`
	Created  time.Time `json:"created"`
}

// New returns a new customer created for the first time
func New(id ID, email Email, password Password) *Customer {
	return &Customer{
		ID:       id,
		Email:    email,
		Password: password,
		Created:  time.Now(),
	}
}

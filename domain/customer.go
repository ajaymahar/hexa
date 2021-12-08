package domain

// Customer struct represent the domain
type Customer struct {
	ID     string `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Email  string `json:"email" xml:"email"`
	Status string `json:"status" xml:"status"`
}

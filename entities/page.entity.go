package entities

type Page struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name" validate:"required"`
	URL    string `json:"url" validate:"required,url"`
	Status bool   `json:"status,omitempty"`
}

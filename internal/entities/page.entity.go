package entities

type Page struct {
	ID     int32  `json:"id,omitempty"`
	Name   string `json:"name" validate:"required"`
	URL    string `json:"url" validate:"required,url"`
	Status string `json:"status,omitempty"`
}

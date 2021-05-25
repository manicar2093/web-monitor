package entities

type Phrase struct {
	ID     string `json:"id,omitempty"`
	Phrase string `json:"phrase" validate:"required"`
}

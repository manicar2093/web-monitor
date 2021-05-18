package entities

type Phrase struct {
	ID     int32  `json:"id,omitempty"`
	Phrase string `json:"phrase" validate:"required"`
}

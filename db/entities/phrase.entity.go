package entities

const PhraseCreationTableSQL = `
CREATE TABLE IF NOT EXISTS "phrases" (
	id TEXT NOT NULL,
	phrase TEXT NOT NULL,
	PRIMARY KEY("id")
);
`

type Phrase struct {
	ID     string `json:"id,omitempty"`
	Phrase string `json:"phrase" validate:"required"`
}

func (c Phrase) Table() string {
	return "phrases"
}

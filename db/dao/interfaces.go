package dao

import "github.com/manicar2093/web-monitor/db/entities"

type PageDao interface {
	GetAllPages() ([]entities.Page, error)
	Delete(pageID string) error
	Save(page entities.Page) error
	Update(data *entities.Page) error
	FindPageByURL(url string) (entities.Page, error)
}

type PhraseDao interface {
	GetAllPhrases() ([]entities.Phrase, error)
	Delete(phraseID string) error
	Save(phrase entities.Phrase) error
	Update(phrase *entities.Phrase) error
}

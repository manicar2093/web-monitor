package services

import (
	"github.com/manicar2093/web-monitor/internal/entities"
)

type PhraseService interface {
	GetAllPhrases() ([]entities.Phrase, error)
	DeletePhrase(phraseID int32) error
	AddPhrase(phrase entities.Phrase) error
	EditPhrase(data entities.Phrase) error
}

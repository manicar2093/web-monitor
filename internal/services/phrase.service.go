package services

import (
	"github.com/google/uuid"
	"github.com/manicar2093/web-monitor/internal/dao"
	"github.com/manicar2093/web-monitor/internal/entities"
)

type PhraseService interface {
	GetAllPhrases() ([]entities.Phrase, error)
	DeletePhrase(phraseID string) error
	AddPhrase(phrase entities.Phrase) error
	EditPhrase(data entities.Phrase) error
}

type PhraseServiceImpl struct {
	phraseDao dao.PhraseDao
}

func NewPhraseService(phraseDao dao.PhraseDao) PhraseService {
	return &PhraseServiceImpl{phraseDao: phraseDao}
}

func (p PhraseServiceImpl) GetAllPhrases() ([]entities.Phrase, error) {
	return p.phraseDao.GetAllPhrases()
}

func (p PhraseServiceImpl) DeletePhrase(phraseID string) error {
	return p.phraseDao.Delete(phraseID)
}

func (p PhraseServiceImpl) AddPhrase(phrase entities.Phrase) error {
	phrase.ID = uuid.NewString()
	return p.phraseDao.Save(phrase)
}

func (p PhraseServiceImpl) EditPhrase(data entities.Phrase) error {
	return p.phraseDao.Update(&data)
}

package services

import (
	"github.com/google/uuid"
	"github.com/manicar2093/web-monitor/db/dao"
	"github.com/manicar2093/web-monitor/db/entities"
)

type PhraseService interface {
	AddPhrase(phrase entities.Phrase) error
}

type PhraseServiceImpl struct {
	phraseDao dao.PhraseDao
}

func NewPhraseService(phraseDao dao.PhraseDao) PhraseService {
	return &PhraseServiceImpl{phraseDao: phraseDao}
}

func (p PhraseServiceImpl) AddPhrase(phrase entities.Phrase) error {
	phrase.ID = uuid.NewString()
	return p.phraseDao.Save(phrase)
}

package dao

import (
	"encoding/json"

	"github.com/manicar2093/web-monitor/connections"
	"github.com/manicar2093/web-monitor/entities"
)

type PhraseDao interface {
	GetAllPhrases() ([]entities.Phrase, error)
	Delete(phraseID string) error
	Save(phrase entities.Phrase) error
	Update(phrase *entities.Phrase) error
}

type PhraseDaoImpl struct {
	fdb *connections.FileDatabase
}

func NewPhraseDao(fdb *connections.FileDatabase) PhraseDao {
	return &PhraseDaoImpl{fdb: fdb}
}

func (p PhraseDaoImpl) GetAllPhrases() ([]entities.Phrase, error) {
	var res []entities.Phrase

	err := p.fdb.ReadData(func(data string) error {
		return json.Unmarshal([]byte(data), &res)
	})

	return res, err
}

func (p PhraseDaoImpl) Delete(phraseID string) error {
	var phrases []entities.Phrase

	return p.fdb.SaveData(func(data string) (string, error) {

		err := json.Unmarshal([]byte(data), &phrases)
		if err != nil {
			return "", err
		}

		var newDB []entities.Phrase
		var rewrite bool
		for _, v := range phrases {
			if v.ID != phraseID {
				newDB = append(newDB, v)
			} else {
				rewrite = true
			}
		}
		if len(phrases) == 1 && rewrite {
			return string("[]"), err
		}
		d, err := json.Marshal(&newDB)
		return string(d), err
	})
}

func (p PhraseDaoImpl) Save(phrase entities.Phrase) error {
	var phrases []entities.Phrase
	return p.fdb.SaveData(func(data string) (string, error) {
		if data != "" {
			err := json.Unmarshal([]byte(data), &phrases)
			if err != nil {
				return "", err
			}
		}
		phrases = append(phrases, phrase)
		d, err := json.Marshal(&phrases)
		if err != nil {
			return "", err
		}
		return string(d), nil
	})
}

func (p PhraseDaoImpl) Update(phrase *entities.Phrase) error {
	var phrases []entities.Phrase

	return p.fdb.SaveData(func(data string) (string, error) {
		err := json.Unmarshal([]byte(data), &phrases)
		if err != nil {
			return "", err
		}

		found := false
		for i, v := range phrases {
			if v.ID == phrase.ID {
				phrases[i] = *phrase
				found = true
				break
			}
		}

		if !found {
			return "", ErrNotFound
		}

		j, err := json.Marshal(&phrases)

		return string(j), err

	})
}

package dao

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/manicar2093/web-monitor/db/entities"
)

type PhraseDaoSqlite struct {
	repo rel.Repository
	ctx  context.Context
}

func NewPhraseDaoSqlite(repo rel.Repository) PhraseDao {
	return &PhraseDaoSqlite{repo: repo, ctx: context.Background()}
}

func (c *PhraseDaoSqlite) GetAllPhrases() ([]entities.Phrase, error) {
	var phrases []entities.Phrase
	if err := c.repo.FindAll(c.ctx, &phrases); err != nil {
		return phrases, err
	}
	return phrases, nil
}

func (c *PhraseDaoSqlite) Delete(phraseID string) error {
	phrase := entities.Phrase{ID: phraseID}
	return c.repo.Delete(c.ctx, &phrase)
}

func (c *PhraseDaoSqlite) Save(phrase entities.Phrase) error {
	return c.repo.Insert(c.ctx, &phrase)
}

func (c *PhraseDaoSqlite) Update(phrase *entities.Phrase) error {
	return c.repo.Update(c.ctx, phrase)
}

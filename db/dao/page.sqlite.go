package dao

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/web-monitor/db/entities"
)

type PageDaoSqlite struct {
	repo rel.Repository
	ctx  context.Context
}

func NewPageDaoSqlite(repo rel.Repository) PageDao {
	return &PageDaoSqlite{repo: repo, ctx: context.Background()}
}

func (c *PageDaoSqlite) GetAllPages() ([]entities.Page, error) {
	var pages []entities.Page
	err := c.repo.FindAll(c.ctx, &pages)
	if err != nil {
		return pages, err
	}
	return pages, nil
}

func (c *PageDaoSqlite) Delete(pageID string) error {
	page := entities.Page{ID: pageID}
	return c.repo.Delete(c.ctx, &page)
}

func (c *PageDaoSqlite) Save(page entities.Page) error {
	return c.repo.Insert(c.ctx, &page)
}

func (c *PageDaoSqlite) Update(data *entities.Page) error {
	return c.repo.Update(c.ctx, data)
}

func (c *PageDaoSqlite) FindPageByURL(url string) (entities.Page, error) {
	var page entities.Page
	if err := c.repo.Find(c.ctx, &page, where.Eq("url", url)); err != nil {
		return page, err
	}
	return page, nil
}

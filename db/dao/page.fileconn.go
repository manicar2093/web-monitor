package dao

import (
	"encoding/json"
	"errors"

	"github.com/manicar2093/web-monitor/db/connections"
	"github.com/manicar2093/web-monitor/db/entities"
)

var (
	ErrNotFound = errors.New("page not found")
)

type PageDaoImpl struct {
	fdb *connections.FileDatabase
}

func NewPageDao(fdb *connections.FileDatabase) PageDao {
	return &PageDaoImpl{fdb}
}

func (p PageDaoImpl) GetAllPages() ([]entities.Page, error) {
	var res []entities.Page

	err := p.fdb.ReadData(func(data string) error {
		return json.Unmarshal([]byte(data), &res)
	})

	return res, err

}

func (p PageDaoImpl) Delete(pageID string) error {

	var pages []entities.Page

	return p.fdb.SaveData(func(data string) (string, error) {

		err := json.Unmarshal([]byte(data), &pages)
		if err != nil {
			return "", err
		}

		var newDB []entities.Page
		var rewrite bool
		for _, v := range pages {
			if v.ID != pageID {
				newDB = append(newDB, v)
			} else {
				rewrite = true
			}
		}
		if len(pages) == 1 && rewrite {
			return string("[]"), err
		}
		d, err := json.Marshal(&newDB)
		return string(d), err
	})
}

func (p PageDaoImpl) Save(page entities.Page) error {

	var pages []entities.Page
	return p.fdb.SaveData(func(data string) (string, error) {
		if data != "" {
			err := json.Unmarshal([]byte(data), &pages)
			if err != nil {
				return "", err
			}
		}
		pages = append(pages, page)
		d, err := json.Marshal(&pages)
		if err != nil {
			return "", err
		}
		return string(d), nil
	})
}

func (p PageDaoImpl) Update(page *entities.Page) error {
	var pages []entities.Page

	return p.fdb.SaveData(func(data string) (string, error) {
		err := json.Unmarshal([]byte(data), &pages)
		if err != nil {
			return "", err
		}

		found := false
		for i, v := range pages {
			if v.ID == page.ID {
				pages[i] = *page
				found = true
				break
			}
		}

		if !found {
			return "", ErrNotFound
		}

		j, err := json.Marshal(&pages)

		return string(j), err

	})

}

func (p PageDaoImpl) FindPageByURL(url string) (entities.Page, error) {
	var pages []entities.Page
	var res entities.Page

	err := p.fdb.ReadData(func(data string) error {
		err := json.Unmarshal([]byte(data), &pages)
		if err != nil {
			return err
		}
		found := false
		for _, v := range pages {
			if v.URL == url {
				res = v
				found = true
				return nil
			}
		}

		if !found {
			return ErrNotFound
		}

		return nil

	})

	return res, err
}

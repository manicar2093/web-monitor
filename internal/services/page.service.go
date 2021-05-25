package services

import (
	"github.com/google/uuid"
	"github.com/manicar2093/web-monitor/internal/dao"
	"github.com/manicar2093/web-monitor/internal/entities"
)

type PageService interface {
	GetAllPages() ([]entities.Page, error)
	DeletePage(pageID string) error
	AddPage(page entities.Page) (entities.Page, error)
	EditPage(data entities.Page) error
	PageExists(url string) (bool, error)
}

type PageServiceImpl struct {
	pageDao dao.PageDao
}

func NewPageDao(pageDao dao.PageDao) PageService {
	return &PageServiceImpl{pageDao: pageDao}
}

func (p PageServiceImpl) GetAllPages() ([]entities.Page, error) {
	return p.pageDao.GetAllPages()
}

func (p PageServiceImpl) DeletePage(pageID string) error {
	return p.pageDao.Delete(pageID)
}

func (p PageServiceImpl) AddPage(page entities.Page) (entities.Page, error) {
	page.ID = uuid.NewString()
	err := p.pageDao.Save(page)
	return page, err
}

func (p PageServiceImpl) EditPage(data entities.Page) error {
	return p.pageDao.Update(&data)
}

func (p PageServiceImpl) PageExists(url string) (bool, error) {
	_, err := p.pageDao.FindPageByURL(url)
	if err != nil {
		if err == dao.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

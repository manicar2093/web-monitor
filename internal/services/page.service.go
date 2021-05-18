package services

import "github.com/manicar2093/web-monitor/internal/entities"

type PageService interface {
	GetAllPages() ([]entities.Page, error)
	DeletePage(pageID int32) error
	AddPage(page entities.Page) (entities.Page, error)
	EditPage(data entities.Page) error
	PageExists(url string) (bool, error)
}

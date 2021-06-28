package services

import (
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/manicar2093/web-monitor/dao"
	"github.com/manicar2093/web-monitor/entities"
	"github.com/manicar2093/web-monitor/models"
)

type Notification struct {
	PageID string         `json:"page_id,omitempty"`
	Error  string         `json:"error,omitempty"`
	Cause  string         `json:"cause,omitempty"`
	Page   *entities.Page `json:"page,inline,omitempty"`
}

type ValidatorService interface {
	// ValidatePages valida las paginas. Al haber error realiza el panic
	Start()
	ValidatePage(page *entities.Page, isInstant bool) (Notification, bool)
}

type ValidatorServiceImpl struct {
	Seconds   int
	pagesDao  dao.PageDao
	observers []models.Observer
	_cron     *gocron.Scheduler
	client    models.HTTPClient
}

func NewValidatorService(Seconds int, pagesDao dao.PageDao, client models.HTTPClient, observers ...models.Observer) ValidatorService {
	v := &ValidatorServiceImpl{
		Seconds,
		pagesDao,
		observers,
		gocron.NewScheduler(time.UTC),
		client,
	}
	v.Start()
	return v
}

// ValidatePages valida las paginas. Al haber error realiza el panic
func (v ValidatorServiceImpl) Start() {
	v._cron.Every(v.Seconds).Seconds().Tag("validator").Do(v.validateAllPages)
	go func() {
		v._cron.StartBlocking()
	}()
}

func (v ValidatorServiceImpl) validateAllPages() {
	log.Println("Comienza la validación")
	pages, err := v.pagesDao.GetAllPages()
	if err != nil {
		panic(err)
	}

	if len(pages) == 0 {
		log.Println("Sin paginas para validar. Termina proceso")
		return
	}

	for _, d := range pages {
		d.CreateMemento()
		n, changed := v.ValidatePage(&d, false)
		if changed {
			v.notifyAll(&n)
		}
	}

	log.Println("Termina la validación de paginas")
}

func (v ValidatorServiceImpl) ValidatePage(page *entities.Page, isInstant bool) (Notification, bool) {

	res, err := v.client.Get(page.URL)
	if err != nil {

		n := Notification{
			PageID: page.ID,
			Error:  err.Error(),
			Cause:  "client error. validate correct page registry",
		}

		return n, false
	}

	if isInstant {
		return v.instantValidation(page, res)
	}

	return v.normalValidation(page, res)

}

func (v ValidatorServiceImpl) instantValidation(page *entities.Page, res *http.Response) (Notification, bool) {
	page.AssignHTTPResValues(res)
	v.pagesDao.Update(page)
	return Notification{Page: page}, false
}

func (v ValidatorServiceImpl) normalValidation(page *entities.Page, res *http.Response) (Notification, bool) {

	if page.HasChange(res) {
		v.pagesDao.Update(page)
		return Notification{Page: page}, true
	}

	return Notification{}, false
}

func (v ValidatorServiceImpl) notifyAll(data interface{}) {
	for _, d := range v.observers {
		d.Notify(data)
	}
}

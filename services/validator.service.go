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

type ValidatorService interface {
	Start()
	// ValidatePages valida las paginas. Al haber error realiza el panic
	ValidatePage(page *entities.Page, isInstant bool) (models.Notification, bool)
}

type ValidatorServiceImpl struct {
	Seconds   int
	pagesDao  dao.PageDao
	observers []models.Observer
	_cron     *gocron.Scheduler
	client    models.HTTPClient
}

func NewValidatorService(pagesDao dao.PageDao, client models.HTTPClient) ValidatorService {
	v := &ValidatorServiceImpl{
		pagesDao: pagesDao,
		client:   client,
	}
	return v
}

func NewValidatorServiceAndStart(Seconds int, pagesDao dao.PageDao, client models.HTTPClient, observers ...models.Observer) ValidatorService {
	v := &ValidatorServiceImpl{
		Seconds:   Seconds,
		pagesDao:  pagesDao,
		_cron:     gocron.NewScheduler(time.UTC),
		observers: observers,
		client:    client,
	}
	v.Start()
	return v
}

func (v ValidatorServiceImpl) Start() {
	v._cron.Every(v.Seconds).Seconds().Tag("validator").Do(v.validateAllPages)
	go func() {
		v._cron.StartBlocking()
	}()
}

// ValidatePages valida las paginas. Al haber error realiza el panic
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

func (v ValidatorServiceImpl) ValidatePage(page *entities.Page, directValidation bool) (models.Notification, bool) {

	res, err := v.client.Get(page.URL)
	if err != nil {

		n := models.Notification{
			PageID: page.ID,
			Error:  err.Error(),
			Cause:  "client error. validate correct page registry",
		}

		return n, false
	}

	if directValidation {
		return v.instantValidation(page, res)
	}

	return v.cronValidation(page, res)

}

func (v ValidatorServiceImpl) instantValidation(page *entities.Page, res *http.Response) (models.Notification, bool) {
	page.AssignHTTPResValues(res)
	v.pagesDao.Update(page)
	return models.Notification{Page: page}, false
}

func (v ValidatorServiceImpl) cronValidation(page *entities.Page, res *http.Response) (models.Notification, bool) {

	if page.HasChange(res) {
		v.pagesDao.Update(page)
		return models.Notification{Page: page}, true
	}

	return models.Notification{}, false
}

func (v ValidatorServiceImpl) notifyAll(data interface{}) {
	for _, d := range v.observers {
		d.Notify(data)
	}
}

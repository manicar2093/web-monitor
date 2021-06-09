package services

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/manicar2093/web-monitor/dao"
	"github.com/manicar2093/web-monitor/entities"
	"github.com/manicar2093/web-monitor/models"
)

type Notification struct {
	PageID     string `json:"pageID"`
	Error      string `json:"error"`
	StatusCode int    `json:"status_code,omitempty"`
	Cause      string `json:"cause"`
}

type ValidatorService interface {
	// ValidatePages valida las paginas. Al haber error realiza el panic
	Start()
}

type ValidatorServiceImpl struct {
	Minutes   int
	pagesDao  dao.PageDao
	observers []models.Observer
	_cron     *gocron.Scheduler
	client    *http.Client
}

func NewValidatorService(minutes int, pagesDao dao.PageDao, observers ...models.Observer) ValidatorService {
	c := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	v := &ValidatorServiceImpl{minutes, pagesDao, observers, gocron.NewScheduler(time.UTC), c}
	v.Start()
	return v
}

// ValidatePages valida las paginas. Al haber error realiza el panic
func (v ValidatorServiceImpl) Start() {
	v._cron.Every(v.Minutes).Seconds().Tag("validator").Do(v.validateAllPages)
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
		v.ValidatePage(d)
	}

	log.Println("Termina la validación de paginas")
}

func (v ValidatorServiceImpl) ValidatePage(page entities.Page) {
	res, err := v.client.Get(page.URL)
	// log.Println("Response from", page.URL, ":", res)
	// TODO: distinguir cuando regresa un 429 Too Many Requests para excluirlo por un determinado tiempo
	if err != nil {
		v.notifyAll(
			&Notification{
				PageID: page.ID,
				Error:  err.Error(),
				Cause:  "Error on client :/",
			})
		page.Status = false
		v.pagesDao.Update(&page)
		return
	}

	if res.StatusCode != http.StatusOK {
		v.notifyAll(
			&Notification{
				PageID: page.ID,
				Error:  "Calling to URL wasn't success",
				Cause:  "No 200 status code",
			})
		page.Status = false
		v.pagesDao.Update(&page)
		return
	}
	// TODO: agregar distintivo cuando camba status de false a true

	page.Status = true
	v.pagesDao.Update(&page)
}

func (v ValidatorServiceImpl) notifyAll(data interface{}) {
	for _, d := range v.observers {
		d.Notify(data)
	}
}

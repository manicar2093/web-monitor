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
	Error      string `json:"error,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	Cause      string `json:"cause"`
	Recovered  bool   `json:"recovered"`
}

type ValidatorService interface {
	// ValidatePages valida las paginas. Al haber error realiza el panic
	Start()
	ValidatePage(page entities.Page) Notification
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
		n := v.ValidatePage(d)
		v.notifyAll(&n)
	}

	log.Println("Termina la validación de paginas")
}

func (v ValidatorServiceImpl) ValidatePage(page entities.Page) Notification {
	res, err := v.client.Get(page.URL)
	// log.Println("Response from", page.URL, ":", res)
	// TODO: distinguir cuando regresa un 429 Too Many Requests para excluirlo por un determinado tiempo
	if err != nil {

		page.Status = false
		v.pagesDao.Update(&page)
		return Notification{
			PageID:    page.ID,
			Error:     err.Error(),
			Cause:     "Error on client :/",
			Recovered: false,
		}
	}

	if res.StatusCode != http.StatusOK {

		page.Status = false
		v.pagesDao.Update(&page)
		return Notification{
			PageID:     page.ID,
			Error:      "Calling to URL wasn't success",
			Cause:      "No 200 status code",
			StatusCode: res.StatusCode,
			Recovered:  false,
		}
	}
	// TODO: agregar distintivo cuando camba status de false a true

	if !page.Status {
		page.Status = true
		v.pagesDao.Update(&page)
		return Notification{
			PageID:     page.ID,
			Cause:      "Good response",
			StatusCode: res.StatusCode,
			Recovered:  true,
		}
	}

	page.Status = true
	v.pagesDao.Update(&page)
	return Notification{}
}

func (v ValidatorServiceImpl) notifyAll(data interface{}) {
	for _, d := range v.observers {
		d.Notify(data)
	}
}

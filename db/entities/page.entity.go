package entities

import (
	"net/http"

	"github.com/manicar2093/web-monitor/utils"
)

type PageMemento struct {
	Status    string
	Code      int
	IsWorking bool
}

type Page struct {
	ID        string       `json:"id,omitempty"`
	Name      string       `json:"name" validate:"required"`
	URL       string       `json:"url" validate:"required,url"`
	Status    string       `json:"status"`
	Code      int          `json:"code"`
	Recovered bool         `json:"recovered"`
	IsWorking bool         `json:"is_working"`
	memento   *PageMemento `json:"-"`
}

// CreateMemento creates state for the page to validate if there were a change in its status
func (p *Page) CreateMemento() {
	p.memento = &PageMemento{
		Status:    p.Status,
		Code:      p.Code,
		IsWorking: p.IsWorking,
	}
}

// HasChange validates if the status and code changed of memento state
func (p *Page) HasChange(res *http.Response) bool {

	p.AssignHTTPResValues(res)

	if p.memento.Code != p.Code && p.memento.Status != p.Status {
		p.Recovered = p.memento.IsWorking != p.IsWorking && p.IsWorking
		return true
	}
	return false
}

func (p *Page) AssignHTTPResValues(res *http.Response) {
	p.Status = http.StatusText(res.StatusCode)
	p.Code = res.StatusCode
	p.IsWorking = utils.IsValidStatus(p.Code)
}

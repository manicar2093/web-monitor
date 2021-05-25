package controllers

import (
	"net/http"

	"github.com/manicar2093/web-monitor/internal/services"
)

type TemplateController interface {
	IndexPage(w http.ResponseWriter, r *http.Request)
}

type TemplateControllerImpl struct {
	templateService services.TemplateService
}

func NewTemplateController(templateService services.TemplateService) TemplateController {
	return &TemplateControllerImpl{templateService: templateService}
}

func (t TemplateControllerImpl) IndexPage(w http.ResponseWriter, r *http.Request) {

	err := t.templateService.Render(w, "index", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.templateService.Render(w, "500", nil)
		return
	}

}

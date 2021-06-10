package services

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

var tplFiles *template.Template

type TemplateService interface {
	Render(w http.ResponseWriter, templateName string, content interface{}) error
}

type TemplateServiceImpl struct{}

func NewTemplateService(files *embed.FS) TemplateService {
	tplFiles = template.Must(template.ParseFS(files, "templates/*.html"))
	return &TemplateServiceImpl{}
}

func (t TemplateServiceImpl) Render(w http.ResponseWriter, templateName string, content interface{}) error {
	err := tplFiles.ExecuteTemplate(w, fmt.Sprintf("%s.html", templateName), content)
	if err != nil {
		return err
	}
	return nil
}

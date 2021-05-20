package services

import (
	"embed"
	"net/http"
)

type TemplateService interface {
	Render(w http.ResponseWriter, templateName string, content interface{}) error
}

// go:embed ../../templates
var templates embed.FS

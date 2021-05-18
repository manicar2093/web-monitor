package controllers

import "net/http"

type TemplateController interface {
	ServeTemplateApp(w http.ResponseWriter, r *http.Request)
}

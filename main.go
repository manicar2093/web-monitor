package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manicar2093/web-monitor/controllers"
	"github.com/manicar2093/web-monitor/services"
)

//go:embed templates/*
var tpl embed.FS
var controller controllers.TemplateController
var templateService services.TemplateService

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.IndexPage)

	log.Fatal("Error al iniciar servidor: ", http.ListenAndServe(":7890", router))

}

func init() {
	templateService = services.NewTemplateService(&tpl)
	controller = controllers.NewTemplateController(templateService)
}

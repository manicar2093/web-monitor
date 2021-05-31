package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manicar2093/web-monitor/connections"
	"github.com/manicar2093/web-monitor/controllers"
	"github.com/manicar2093/web-monitor/dao"
	"github.com/manicar2093/web-monitor/services"
	"github.com/manicar2093/web-monitor/sse"
)

//go:embed templates/*
var tpl embed.FS
var phraseDao dao.PhraseDao
var pageDao dao.PageDao

var templateService services.TemplateService
var phraseService services.PhraseService
var pageService services.PageService
var validatorService services.ValidatorService

var controller controllers.TemplateController
var phraseController controllers.PhraseController
var pageController controllers.PageController
var sseValidatorController *sse.Broker

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.IndexPage).Methods(http.MethodGet)

	phraseRouter := router.PathPrefix("/phrases").Subrouter()
	phraseRouter.HandleFunc("/add", phraseController.AddPhrase).Methods(http.MethodPost)
	phraseRouter.HandleFunc("/delete", phraseController.DeletePhrase).Methods(http.MethodDelete)
	phraseRouter.HandleFunc("/update", phraseController.EditPhrase).Methods(http.MethodPut)
	phraseRouter.HandleFunc("/all", phraseController.GetAllPhrases).Methods(http.MethodGet)

	pageRouter := router.PathPrefix("/pages").Subrouter()
	pageRouter.HandleFunc("/add", pageController.AddPage).Methods(http.MethodPost)
	pageRouter.HandleFunc("/delete", pageController.DeletePage).Methods(http.MethodDelete)
	pageRouter.HandleFunc("/update", pageController.EditPage).Methods(http.MethodPut)
	pageRouter.HandleFunc("/all", pageController.GetAllPages).Methods(http.MethodGet)
	pageRouter.HandleFunc("/do-exists", pageController.PageExists).Methods(http.MethodGet)
	pageRouter.HandleFunc("/validate", pageController.ValidatePage).Methods(http.MethodGet)

	sseRouter := router.PathPrefix("/sse").Subrouter()
	sseRouter.Handle("/sse-validator", sseValidatorController).Methods(http.MethodGet)

	log.Fatal("Error al iniciar servidor: ", http.ListenAndServe(":7890", router))

}

func init() {
	phraseConnection := connections.NewFileDatabase("./phrases.json")
	pagesConnection := connections.NewFileDatabase("./pages.json")

	phraseDao = dao.NewPhraseDao(phraseConnection)
	pageDao = dao.NewPageDao(pagesConnection)

	templateService = services.NewTemplateService(&tpl)
	phraseService = services.NewPhraseService(phraseDao)
	pageService = services.NewPageService(pageDao)

	controller = controllers.NewTemplateController(templateService)
	phraseController = controllers.NewPhraseController(phraseDao, phraseService)
	pageController = controllers.NewPageController(pageDao, pageService)
	sseValidatorController = sse.NewBroker()

	validatorService = services.NewValidatorService(5, pageDao, sseValidatorController)
}

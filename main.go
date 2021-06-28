package main

import (
	"crypto/tls"
	"embed"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/manicar2093/web-monitor/config"
	"github.com/manicar2093/web-monitor/connections"
	"github.com/manicar2093/web-monitor/controllers"
	"github.com/manicar2093/web-monitor/dao"
	"github.com/manicar2093/web-monitor/models"
	"github.com/manicar2093/web-monitor/services"
	"github.com/manicar2093/web-monitor/sse"
)

//go:embed templates/*
var tpl embed.FS

//go:embed static/*
var static embed.FS
var phraseDao dao.PhraseDao
var pageDao dao.PageDao

var templateService services.TemplateService
var phraseService services.PhraseService
var pageService services.PageService
var validatorService services.ValidatorService

var client models.HTTPClient

var controller controllers.TemplateController
var phraseController controllers.PhraseController
var pageController controllers.PageController
var sseValidatorController *sse.Broker

var srv *http.Server

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.IndexPage).Methods(http.MethodGet)
	router.HandleFunc("/close", func(rw http.ResponseWriter, r *http.Request) {

		e := srv.Shutdown(r.Context())
		if e != nil {
			panic(e)
		}
	}).Methods(http.MethodPost)

	router.PathPrefix("/static/").Handler(http.FileServer(http.FS(static)))

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
	pageRouter.HandleFunc("/validate", pageController.ValidatePage).Methods(http.MethodPost)

	sseRouter := router.PathPrefix("/sse").Subrouter()
	sseRouter.Handle("/sse-validator", sseValidatorController).Methods(http.MethodGet)

	srv = &http.Server{
		Addr:         config.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	log.Fatal("Error al iniciar servidor: ", srv.ListenAndServe())

}

func init() {
	sseValidatorController = sse.NewBroker()

	phraseConnection := connections.NewFileDatabase(config.PhrasesFile)
	pagesConnection := connections.NewFileDatabase(config.PagesFile)

	phraseDao = dao.NewPhraseDao(phraseConnection)
	pageDao = dao.NewPageDao(pagesConnection)

	templateService = services.NewTemplateService(&tpl)
	phraseService = services.NewPhraseService(phraseDao)
	pageService = services.NewPageService(pageDao)
	client = &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	validatorService = services.NewValidatorServiceAndStart(config.SecondsToValidate, pageDao, client, sseValidatorController)

	controller = controllers.NewTemplateController(templateService)
	phraseController = controllers.NewPhraseController(phraseDao, phraseService)
	pageController = controllers.NewPageController(pageDao, pageService, validatorService)

}

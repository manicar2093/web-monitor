package main

import (
	"crypto/tls"
	"embed"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/manicar2093/web-monitor/config"
	"github.com/manicar2093/web-monitor/controllers"
	"github.com/manicar2093/web-monitor/db/connections"
	"github.com/manicar2093/web-monitor/db/dao"
	"github.com/manicar2093/web-monitor/models"
	"github.com/manicar2093/web-monitor/scripts"
	"github.com/manicar2093/web-monitor/services"
	"github.com/manicar2093/web-monitor/sse"
)

//go:generate cp -r ../../templates ./templates/
//go:embed templates/*
var tpl embed.FS

//go:generate cp -r ../../static ./static/
//go:embed static/*
var static embed.FS

var phraseDao dao.PhraseDao
var phraseDaoSqlite dao.PhraseDao
var pageDao dao.PageDao
var pageDaoSqlite dao.PageDao

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

	log.Printf("Server started on http://localhost%s port", config.Port)

	log.Fatal("Error al iniciar servidor: ", srv.ListenAndServe())

}

func init() {
	sseValidatorController = sse.NewBroker()

	phraseConnection := connections.NewFileDatabase(config.PhrasesFile)
	pagesConnection := connections.NewFileDatabase(config.PagesFileV2)
	repo := connections.NewSqliteConection(config.DatabaseName)

	phraseDao = dao.NewPhraseDao(phraseConnection)
	pageDao = dao.NewPageDao(pagesConnection)
	pageDaoSqlite = dao.NewPageDaoSqlite(repo)
	phraseDaoSqlite = dao.NewPhraseDaoSqlite(repo)

	scripts.DatabaseMigrationV1_3(pageDao)
	scripts.DatabaseMigrationToSqlite(&scripts.MigrationToSqlite{
		Conn:            repo,
		PageDaoFile:     pageDao,
		PageDaoSqlite:   pageDaoSqlite,
		PhraseDaoFile:   phraseDao,
		PhraseDaoSqlite: phraseDaoSqlite,
	})

	templateService = services.NewTemplateService(&tpl)
	phraseService = services.NewPhraseService(phraseDaoSqlite)
	pageService = services.NewPageService(pageDaoSqlite)
	client = &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	validatorService = services.NewValidatorServiceAndStart(config.SecondsToValidate, pageDaoSqlite, client, sseValidatorController)

	controller = controllers.NewTemplateController(templateService)
	phraseController = controllers.NewPhraseController(phraseDaoSqlite, phraseService)
	pageController = controllers.NewPageController(pageDaoSqlite, pageService, validatorService)

}

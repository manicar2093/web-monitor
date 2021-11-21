package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/manicar2093/web-monitor/db/dao"
	"github.com/manicar2093/web-monitor/db/entities"
	"github.com/manicar2093/web-monitor/services"
)

type PhraseRequest struct {
	PhraseID string `json:"id"`
	URL      string `json:"url"`
}

type PhraseController interface {
	// GetAllPages regresa todas las paginas registradas
	GetAllPhrases(w http.ResponseWriter, r *http.Request)
	//DeletePage se usa para eliminar una pagina
	DeletePhrase(w http.ResponseWriter, r *http.Request)
	// AddPage para agregar una nueva pagina
	AddPhrase(w http.ResponseWriter, r *http.Request)
	// EditPage para editar una pagina
	EditPhrase(w http.ResponseWriter, r *http.Request)
}

type PhraseControllerImpl struct {
	phraseDao     dao.PhraseDao
	phraseService services.PhraseService
}

func NewPhraseController(phraseDao dao.PhraseDao, phraseService services.PhraseService) PhraseController {
	return &PhraseControllerImpl{phraseDao, phraseService}
}

// GetAllPages regresa todas las paginas registradas
func (p PhraseControllerImpl) GetAllPhrases(w http.ResponseWriter, r *http.Request) {
	pages, err := p.phraseDao.GetAllPhrases()
	if err != nil {
		log.Printf("error al obtener todas las frases. Detalles: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(pages)
}

//DeletePage se usa para eliminar una pagina
func (p PhraseControllerImpl) DeletePhrase(w http.ResponseWriter, r *http.Request) {
	var phraseReq PhraseRequest
	err := json.NewDecoder(r.Body).Decode(&phraseReq)
	if err != nil {
		log.Printf("error al obtener ID de la frase a eliminar. Detalles: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = p.phraseDao.Delete(phraseReq.PhraseID)
	if err != nil {
		log.Printf("error al obtener todas las frases. Detalles: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// AddPage para agregar una nueva pagina
func (p PhraseControllerImpl) AddPhrase(w http.ResponseWriter, r *http.Request) {
	var phrase entities.Phrase
	err := json.NewDecoder(r.Body).Decode(&phrase)
	if err != nil {
		log.Printf("error al obtener data de la frase para guardar. Detalles: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = p.phraseService.AddPhrase(phrase); err != nil {
		log.Printf("error al guardar la frase. Detalles: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// EditPage para editar una pagina
func (p PhraseControllerImpl) EditPhrase(w http.ResponseWriter, r *http.Request) {
	var phrase entities.Phrase
	err := json.NewDecoder(r.Body).Decode(&phrase)
	if err != nil {
		log.Printf("error al obtener data de la frase a editar. Detalles: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = p.phraseDao.Update(&phrase); err != nil {
		log.Printf("error al guardar la frase. Detalles: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

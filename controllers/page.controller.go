package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/manicar2093/web-monitor/dao"
	"github.com/manicar2093/web-monitor/entities"
	"github.com/manicar2093/web-monitor/services"
)

type PageRequest struct {
	PageID string `json:"id"`
	URL    string `json:"url"`
}

type PageController interface {
	// GetAllPages regresa todas las paginas registradas
	GetAllPages(w http.ResponseWriter, r *http.Request)
	//DeletePage se usa para eliminar una pagina
	DeletePage(w http.ResponseWriter, r *http.Request)
	// AddPage para agregar una nueva pagina
	AddPage(w http.ResponseWriter, r *http.Request)
	// EditPage para editar una pagina
	EditPage(w http.ResponseWriter, r *http.Request)
	// ValidatePage para realizar la validación individual de una pagina
	ValidatePage(w http.ResponseWriter, r *http.Request)
	// PageExists valida si la pagina ya esta registrada
	PageExists(w http.ResponseWriter, r *http.Request)
}

type PageControllerImpl struct {
	pageDao     dao.PageDao
	pageService services.PageService
	validator   services.ValidatorService
}

func NewPageController(pageDao dao.PageDao, pageService services.PageService, validator services.ValidatorService) PageController {
	return &PageControllerImpl{pageDao, pageService, validator}
}

// GetAllPages regresa todas las paginas registradas
func (p PageControllerImpl) GetAllPages(w http.ResponseWriter, r *http.Request) {
	pages, err := p.pageDao.GetAllPages()
	if err != nil {
		log.Printf("error al obtener todas las paginas. Detalles: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(pages)
}

//DeletePage se usa para eliminar una pagina
func (p PageControllerImpl) DeletePage(w http.ResponseWriter, r *http.Request) {
	var pageReq PageRequest
	err := json.NewDecoder(r.Body).Decode(&pageReq)
	if err != nil {
		log.Printf("error al obtener id de la pagina. Detalles: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = p.pageDao.Delete(pageReq.PageID)
	if err != nil {
		log.Printf("error al obtener todas las paginas. Detalles: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// AddPage para agregar una nueva pagina
func (p PageControllerImpl) AddPage(w http.ResponseWriter, r *http.Request) {
	var page entities.Page
	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		log.Printf("error al obtener data de la pagina para guardar. Detalles: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if page, err = p.pageService.AddPage(page); err != nil {
		log.Printf("error al guardar la pagina. Detalles: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	n, _ := p.validator.ValidatePage(&page, true)

	json.NewEncoder(w).Encode(&n)

}

// EditPage para editar una pagina
func (p PageControllerImpl) EditPage(w http.ResponseWriter, r *http.Request) {
	var page entities.Page
	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		log.Printf("error al obtener data de la frase a editar. Detalles: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = p.pageDao.Update(&page); err != nil {
		log.Printf("error al guardar la frase. Detalles: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ValidatePage para realizar la validación individual de una pagina
func (p PageControllerImpl) ValidatePage(w http.ResponseWriter, r *http.Request) {
	var pageID PageRequest
	err := json.NewDecoder(r.Body).Decode(&pageID)
	if err != nil {
		log.Printf("error al obtener id de la pagina. Detalles: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	page, err := p.pageDao.FindPageByURL(pageID.URL)
	if err != nil {
		log.Printf("error al obtener id de la pagina. Detalles: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	n, _ := p.validator.ValidatePage(&page, true)

	json.NewEncoder(w).Encode(n)

}

// PageExists valida si la pagina ya esta registrada
func (p PageControllerImpl) PageExists(w http.ResponseWriter, r *http.Request) {
	var pageReq PageRequest
	err := json.NewDecoder(r.Body).Decode(&pageReq)
	if err != nil {
		log.Printf("error al obtener id de la pagina. Detalles: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	doExists, err := p.pageService.PageExists(pageReq.URL)
	if err != nil {
		log.Printf("error determinar si la pagina ya existe. Detalles: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{
		"exists": doExists,
	})

}

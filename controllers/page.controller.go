package controllers

import "net/http"

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

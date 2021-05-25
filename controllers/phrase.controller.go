package controllers

import "net/http"

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

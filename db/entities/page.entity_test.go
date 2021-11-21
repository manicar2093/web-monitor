package entities

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMemento(t *testing.T) {
	res := http.Response{
		StatusCode: http.StatusNotFound,
	}
	page := Page{
		ID:        "anID",
		Name:      "aName",
		URL:       "aURL",
		Status:    http.StatusText(http.StatusOK),
		Code:      http.StatusOK,
		Recovered: false,
		IsWorking: true,
	}

	page.CreateMemento()

	assert.True(t, page.HasChange(&res), "Should return true. There are changes from the original page state")

}

package services

import (
	"net/http"
	"testing"

	"github.com/manicar2093/web-monitor/entities"
	"github.com/manicar2093/web-monitor/mocks"
)

func TestValidatePage(t *testing.T) {
	pageDao := mocks.PageDao{}
	client := mocks.HTTPClient{}
	page := entities.Page{
		ID:        "anID",
		Name:      "name",
		URL:       "https://google.com",
		Status:    http.StatusText(http.StatusOK),
		Code:      http.StatusOK,
		Recovered: false,
		IsWorking: true,
	}

	client.On("Get", page.URL).Return(&http.Response{
		StatusCode: http.StatusNotFound,
	})
	pageDao.On("Update", &page).Return(nil)

	validator := NewValidatorService(&pageDao, &client)

	n, _ := validator.ValidatePage(&page, true)

	t.Log(n)
}

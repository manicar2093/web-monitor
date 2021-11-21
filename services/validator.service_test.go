package services

import (
	"net/http"
	"testing"

	"github.com/manicar2093/web-monitor/db/entities"
	"github.com/manicar2093/web-monitor/mocks"
	"github.com/stretchr/testify/assert"
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
	}, nil)
	pageDao.On("Update", &page).Return(nil)

	validator := NewValidatorService(&pageDao, &client)

	n, _ := validator.ValidatePage(&page, true)

	client.AssertExpectations(t)
	pageDao.AssertExpectations(t)

	assert.Equal(t, n.Page.IsWorking, false, "should not working")
	assert.Equal(t, n.Page.Recovered, false, "should not recovered")
	assert.Equal(t, n.Page.Code, http.StatusNotFound, "should return a 404 code")

}

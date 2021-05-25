package controllers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/manicar2093/web-monitor/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTemplateController(t *testing.T) {
	httpData := func() (*httptest.ResponseRecorder, *http.Request) {
		return httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil)
	}
	templateServiceMock := &mocks.TemplateService{}

	t.Run("IndexPage happy path", func(t *testing.T) {
		w, r := httpData()
		templateServiceMock.On("Render", w, "index", nil).Run(func(args mock.Arguments) {
			w := args[0].(http.ResponseWriter)
			fmt.Fprint(w, "DOCTYPE")
		}).Return(nil)

		controller := NewTemplateController(templateServiceMock)

		router := mux.NewRouter()
		router.HandleFunc("/", controller.IndexPage)
		router.ServeHTTP(w, r)

		templateServiceMock.AssertExpectations(t)

		require.Equal(t, http.StatusOK, w.Code, "incorrect status code")
		body, err := io.ReadAll(w.Body)
		require.NoError(t, err, "error inesperado")

		require.Contains(t, string(body), "DOCTYPE", "incorrect content")
		templateServiceMock.ExpectedCalls = nil
	})

	t.Run("IndexPage template Render fails", func(t *testing.T) {
		w, r := httpData()
		templateServiceMock.On("Render", w, "index", nil).Return(errors.New("an error occured"))

		templateServiceMock.On("Render", w, "500", nil).Run(func(args mock.Arguments) {
			w := args[0].(http.ResponseWriter)
			fmt.Fprint(w, "DOCTYPE")
		}).Return(nil)

		controller := NewTemplateController(templateServiceMock)

		router := mux.NewRouter()
		router.HandleFunc("/", controller.IndexPage)
		router.ServeHTTP(w, r)

		templateServiceMock.AssertExpectations(t)

		require.Equal(t, http.StatusInternalServerError, w.Code, "incorrect status code")
		body, err := io.ReadAll(w.Body)
		require.NoError(t, err, "error inesperado")

		require.Contains(t, string(body), "DOCTYPE", "incorrect content")
		templateServiceMock.ExpectedCalls = nil
	})
}

package services

import (
	"embed"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed templates/*.html
var tpl embed.FS

func TestTemplateService(t *testing.T) {
	tplService := NewTemplateService(&tpl)

	t.Run("Happy path", func(t *testing.T) {
		w := httptest.NewRecorder()
		err := tplService.Render(w, "index", nil)
		require.NoError(t, err, "unexpected error")
		b, err := io.ReadAll(w.Body)
		require.NoError(t, err, "unexpected error")
		require.Contains(t, string(b), "DOCTYPE", "content do not match")
	})

	t.Run("template do not exists", func(t *testing.T) {
		w := httptest.NewRecorder()
		err := tplService.Render(w, "do_not_exists", nil)
		require.Error(t, err, "should return an error")
		t.Log(err)
	})

}

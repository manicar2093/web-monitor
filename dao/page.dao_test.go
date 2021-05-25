package dao

import (
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/manicar2093/web-monitor/internal/connections"
	"github.com/manicar2093/web-monitor/internal/entities"
	"github.com/stretchr/testify/require"
)

func TestPageDao(t *testing.T) {
	dbFile := "./.testPageDao.json"
	fdb := connections.NewFileDatabase(dbFile)
	dao := NewPageDao(fdb)

	d := []entities.Page{
		{
			ID:     uuid.New().String(),
			Name:   "testing",
			URL:    "aurl",
			Status: true,
		},
		{
			ID:     uuid.New().String(),
			Name:   "testing",
			URL:    "aurl",
			Status: true,
		},
	}

	t.Run("Validar que se puedan guardar registros", func(t *testing.T) {

		for i, v := range d {
			err := dao.Save(v)
			require.NoError(t, err, "error inesperado")
			fdb.ReadData(func(data string) error {
				require.Contains(t, data, d[i].ID, "data incorrecta")
				return nil
			})
		}

	})

	t.Run("validar se regresen todos los registros", func(t *testing.T) {
		pages, err := dao.GetAllPages()
		t.Log(pages)
		require.Len(t, pages, 2, "no estan todos los registros")
		require.NoError(t, err, "error inesperado")
	})

	t.Run("validar el update de los datos", func(t *testing.T) {
		id := d[0].ID
		p := entities.Page{
			ID:     id,
			Name:   "testing",
			URL:    "aurl",
			Status: true,
		}
		require.NoError(t, dao.Update(&p), "error inesperado")
		pages, err := dao.GetAllPages()
		require.NoError(t, err, "error inesperado")
		for _, v := range pages {
			if v.Name == p.Name {
				return
			}
		}
		t.Fatal("no se edito el nombre")
	})

	t.Run("validar la busqueda por url", func(t *testing.T) {
		url := d[0].URL
		page, err := dao.FindPageByURL(url)
		require.NoError(t, err, "error inesperado")
		require.Equal(t, url, page.URL, "no corresponden las url de las paginas")
	})

	t.Run("validar el borrado de los datos", func(t *testing.T) {
		id := d[0].ID
		err := dao.Delete(id)
		require.NoError(t, err, "error inesperado")

		pages, err := dao.GetAllPages()
		require.NoError(t, err, "error inesperado")
		for _, p := range pages {
			if p.ID == id {
				t.Fatal("no debe existir el elemento en el documento")
			}
		}
	})

	t.Cleanup(func() {
		err := os.Remove(dbFile)
		if err != nil {
			log.Fatal(err)
		}
	})

}

package dao

import (
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/manicar2093/web-monitor/connections"
	"github.com/manicar2093/web-monitor/entities"
	"github.com/stretchr/testify/require"
)

func TestPhraseDao(t *testing.T) {
	dbFile := "./.testPhraseDao.json"
	fdb := connections.NewFileDatabase(dbFile)
	dao := NewPhraseDao(fdb)

	d := []entities.Phrase{
		{
			ID:     uuid.New().String(),
			Phrase: "phrase1",
		},
		{
			ID:     uuid.New().String(),
			Phrase: "phrase2",
		},
		{
			ID:     uuid.New().String(),
			Phrase: "phrase3",
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
		phrases, err := dao.GetAllPhrases()
		t.Log(phrases)
		require.Len(t, phrases, 3, "no estan todos los registros")
		require.NoError(t, err, "error inesperado")
	})

	t.Run("validar el update de los datos", func(t *testing.T) {
		id := d[0].ID
		p := entities.Phrase{
			ID:     id,
			Phrase: "Other phrase :D",
		}
		require.NoError(t, dao.Update(&p), "error inesperado")
		phrases, err := dao.GetAllPhrases()
		require.NoError(t, err, "error inesperado")
		for _, v := range phrases {
			if v.Phrase == p.Phrase {
				return
			}
		}
		t.Fatal("no se edito el nombre")
	})

	t.Run("validar el borrado de los datos", func(t *testing.T) {
		id := d[0].ID
		err := dao.Delete(id)
		require.NoError(t, err, "error inesperado")

		phrases, err := dao.GetAllPhrases()
		require.NoError(t, err, "error inesperado")
		for _, p := range phrases {
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

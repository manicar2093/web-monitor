package connections

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/manicar2093/web-monitor/db/entities"
	"github.com/stretchr/testify/require"
)

const fileName = ".test.json"

func deleteTestFile() {
	err := os.Remove(fileName)
	if err != nil {
		log.Fatal(err)
	}
}

// TestNewFileDatabase realiza la validación de las funciones privadas de la struct
func TestNewFileDatabase(t *testing.T) {
	var fdb *FileDatabase

	t.Run("Validar que el archivo se crea", func(t *testing.T) {
		fdb = NewFileDatabase(fileName)
		defer func() {
			d := recover()
			if d != nil {
				t.Fatal(d)
			}
		}()
		if fdb == nil {
			t.Fatal("la base no puede ser nil")
		}
	})

	t.Run("Valida la escritura del archivo", func(t *testing.T) {
		data := []byte("Hello from tests")
		err := fdb.writeFile(data)
		if err != nil {
			t.Fatal(err)
		}

		raw, err := fdb.readFile()
		if err != nil {
			t.Fatal(err)
		}

		if raw == "" {
			t.Fatal("la info no puede estar vacia")
		}
	})

	t.Cleanup(deleteTestFile)

}

func TestPublicMethods(t *testing.T) {
	fdb := NewFileDatabase(fileName)

	testingData := entities.Page{
		ID:     uuid.New().String(),
		Name:   "testing",
		URL:    "aurl",
		Status: http.StatusText(http.StatusOK),
	}

	t.Run("Escribir un objeto JSON", func(t *testing.T) {

		f := func(data string) (string, error) {
			j, err := json.Marshal(&testingData)
			if err != nil {
				return "", err
			}
			return string(j), nil
		}

		err := fdb.SaveData(f)
		if err != nil {
			t.Fatal(err)
		}

		s, err := fdb.readFile()
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(s, "testing") {
			t.Fatal("no hay data guardada")
		}
	})

	t.Run("Leer la información que se encuentra dentro del archivo", func(t *testing.T) {

		var d entities.Page

		f := func(data string) error {
			return json.Unmarshal([]byte(data), &d)
		}

		if err := fdb.ReadData(f); err != nil {
			t.Fatal(err)
		}

		require.Equal(t, testingData.ID, d.ID, "assertion error")
		require.Equal(t, testingData.Name, d.Name, "assertion error")
		require.Equal(t, testingData.Status, d.Status, "assertion error")
		require.Equal(t, testingData.URL, d.URL, "assertion error")
	})

	t.Cleanup(deleteTestFile)
}

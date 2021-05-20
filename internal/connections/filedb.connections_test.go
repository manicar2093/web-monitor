package connections

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
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

	t.Run("Escribir un objeto JSON", func(t *testing.T) {
		testingData := struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Age  int32  `json:"age"`
		}{
			ID:   uuid.New().String(),
			Name: "testing",
			Age:  27,
		}

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

	t.Cleanup(deleteTestFile)
}

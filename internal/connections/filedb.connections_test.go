package connections

import (
	"os"
	"testing"
)

// TestNewFileDatabase realiza la validación de las funciones privadas de la struct
func TestNewFileDatabase(t *testing.T) {
	var fdb *FileDatabase

	t.Run("Validar que el archivo se crea", func(t *testing.T) {
		fdb = NewFileDatabase(".test.json")
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

	t.Cleanup(func() {
		err := os.Remove(fdb.Path)
		if err != nil {
			t.Log("error al hacer cleanup")
			t.Fatal(err)
		}
	})

}

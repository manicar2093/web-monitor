package connections

import (
	"log"
	"os"
	"sync"
)

type FileDatabase struct {
	Path string
	lock *sync.Mutex
}

// ReadFunc recibe la información que se obtiene del archivo. Desde este se debe realizar la asignación de la info a un a struct que se necesite
type ReadFunc func(data string) error

// SaveFunc es para realizar los movimientos a la información antes de ser guardada. La info del archivo será sobreescrito por toda la data que se reciba
type SaveFunc func(data string) (string, error)

func NewFileDatabase(path string) *FileDatabase {
	err := validatePathExists(path)
	if err != nil {
		log.Println("error al validar el archivo solicitado")
		panic(err)
	}
	f := &FileDatabase{
		Path: path,
		lock: &sync.Mutex{},
	}

	err = f.writeFile([]byte("[]"))
	if err != nil {
		panic(err)
	}
	return f
}

func (f FileDatabase) ReadData(callback ReadFunc) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	data, err := f.readFile()
	if err != nil {
		return err
	}

	err = callback(data)
	if err != nil {
		return err
	}
	return nil
}

func (f FileDatabase) SaveData(callback SaveFunc) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	data, err := f.readFile()
	if err != nil {
		return err
	}

	data, err = callback(data)
	if err != nil {
		return err
	}
	err = f.writeFile([]byte(data))
	if err != nil {
		return err
	}
	return nil

}

func (f FileDatabase) readFile() (string, error) {
	content, err := os.ReadFile(f.Path)
	if err != nil {
		return "", err
	}
	return string(content), err
}

func (f FileDatabase) writeFile(data []byte) error {
	file, err := os.OpenFile(f.Path, os.O_WRONLY|os.O_TRUNC, 0777)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func validatePathExists(path string) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	return nil
}

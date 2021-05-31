package services

import (
	"log"
	"sync"
	"testing"

	"github.com/manicar2093/web-monitor/entities"
	"github.com/manicar2093/web-monitor/mocks"
	"github.com/stretchr/testify/mock"
)

func TestValidatePages(t *testing.T) {
	pagesDao := mocks.PageDao{}
	observer := mocks.Observer{}
	pages := []entities.Page{
		{ID: "id1", Name: "A name", URL: "https://google.com", Status: true},
		{ID: "id2", Name: "To fail", URL: "http://localhost:1000", Status: true},
	}
	seconds := 5

	wait := sync.WaitGroup{}

	pagesDao.On("GetAllPages").Return(pages, nil)
	observer.On("Notify", &Notification{PageID: pages[1].ID, Error: "Get \"http://localhost:1000\": dial tcp 127.0.0.1:1000: connect: connection refused", Cause: "Error on client :/"}).Run(func(args mock.Arguments) {
		log.Println("notificando")
		wait.Done()
	})

	pages[1].Status = false
	pagesDao.On("Update", &pages[0]).Return(nil)
	pagesDao.On("Update", &pages[1]).Return(nil)

	wait.Add(1)
	go func() {
		NewValidatorService(seconds, &pagesDao, &observer)
		t.Log("activado :/")
	}()
	wait.Wait()

	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(seconds+3))
	// defer cancel()

	// <-ctx.Done()
	// wait.Done()

}

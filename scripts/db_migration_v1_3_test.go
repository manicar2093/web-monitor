package scripts

import (
	"os"
	"testing"

	"github.com/manicar2093/web-monitor/config"
	"github.com/manicar2093/web-monitor/mocks"
	"github.com/stretchr/testify/mock"
)

func TestDatabaseMigrationV1_3(t *testing.T) {

	pages := `[{"id":"739ef00f-623d-4ed1-867c-fb5ef32eb960","name":"Test0","url":"http://localhost:7890","status":true},
		{"id":"739ef00f-623d-4ed1-867c-fb5ef32eb961","name":"Test1","url":"http://localhost:7890/dont","status":false},
		{"id":"739ef00f-623d-4ed1-867c-fb5ef32eb963","name":"Test3","url":"http://localhost:7890","status":true},
		{"id":"739ef00f-623d-4ed1-867c-fb5ef32eb964","name":"Test4","url":"http://localhost:7890/dont","status":false}]`

	f, e := os.OpenFile(config.PagesFile, os.O_RDWR|os.O_CREATE, 0755)
	if e != nil {
		t.Fatal(e)
	}
	defer f.Close()

	_, e = f.Write([]byte(pages))
	if e != nil {
		t.Fatal(e)
	}
	pageDao := mocks.PageDao{}
	pageDao.On("Save", mock.AnythingOfType("entities.Page")).Return(nil)
	DatabaseMigrationV1_3(&pageDao)
	pageDao.AssertExpectations(t)
}

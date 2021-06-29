package scripts

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/manicar2093/web-monitor/config"
	"github.com/manicar2093/web-monitor/connections"
	"github.com/manicar2093/web-monitor/dao"
	"github.com/manicar2093/web-monitor/entities"
)

// DatabaseMigrationV1_3 validates if a antique versión of the DB is been used to change it to the new one
func DatabaseMigrationV1_3(pageDao dao.PageDao) {
	if !Exists(config.PagesFile) {
		return
	}

	conn := connections.NewFileDatabase(config.PagesFile)
	var d []map[string]interface{}
	e := conn.ReadData(func(data string) error {
		return json.Unmarshal([]byte(data), &d)
	})

	if e != nil {
		panic(e)
	}

	if !(len(d) > 0) {
		return
	}

	for _, item := range d {

		isWorking, status, code := item["status"].(bool), http.StatusText(http.StatusOK), http.StatusOK

		if !isWorking {
			status = http.StatusText(http.StatusNotFound)
			code = http.StatusNotFound
		}
		page := entities.Page{
			ID:        item["id"].(string),
			Name:      item["name"].(string),
			URL:       item["url"].(string),
			IsWorking: isWorking,
			Status:    status,
			Code:      code,
			Recovered: false,
		}

		if e := pageDao.Save(page); e != nil {
			panic(e)
		}
	}

	conn = nil
	e = os.Remove(config.PagesFile)
	if e != nil {
		panic(e)
	}

}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

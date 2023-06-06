package store

import (
	"encoding/json"
	"github.com/DmitryOdintsov/workingWithGit/internal/models"
	"log"
	"net/http"
)

type DatasetSupport struct {
	Data []*models.SupportData
	//sync.Mutex
}

func NewDatasetSupport() *DatasetSupport {
	return &DatasetSupport{}
}

func (d *DatasetSupport) GetSupportAPI() *DatasetSupport {
	urlPath := "http://127.0.0.1:8383/support"
	req, err := http.Get(urlPath)
	if err != nil {
		log.Println("Отсутствует соединение с сервером", err)
		return d
	}
	if req.StatusCode == http.StatusOK {
		err = json.NewDecoder(req.Body).Decode(&d.Data)
		if err != nil {
			log.Println(err)
			return d
		}
	} else if req.StatusCode == http.StatusInternalServerError {
		log.Println("произошла ошибка!")
	}
	return d
}

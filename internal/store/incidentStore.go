package store

import (
	"encoding/json"
	"github.com/DmitryOdintsov/workingWithGit/internal/models"
	"log"
	"net/http"
	"sort"
)

type DatasetIncident struct {
	Data []*models.IncidentData
}

func NewDatasetIncident() *DatasetIncident {
	return &DatasetIncident{}
}

func (d *DatasetIncident) GetIncident() *DatasetIncident {
	urlPath := "http://127.0.0.1:8383/accendent"
	req, err := http.Get(urlPath)
	if err != nil {
		log.Println(err)
		return d
	}
	if req.StatusCode != http.StatusOK {
		return d
	}
	err = json.NewDecoder(req.Body).Decode(&d.Data)
	if err != nil {
		log.Println(err)
		return d
	}
	return d
}

func (d *DatasetIncident) SortedIncidentStatus() []models.IncidentData {
	var data []models.IncidentData
	byteSms, err := json.Marshal(d.Data)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(byteSms, &data)
	sort.Slice(data, func(i, j int) bool { return data[i].Status < data[j].Status })
	return data
}

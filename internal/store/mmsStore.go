package store

import (
	"encoding/json"
	"github.com/DmitryOdintsov/workingWithGit/internal/models"
	"github.com/DmitryOdintsov/workingWithGit/pkg/alfa2"
	"log"
	"net/http"
	"sort"
)

type DatasetMMS struct {
	Data []*models.MMS
}

func NewDatasetMMS() *DatasetMMS {
	return &DatasetMMS{}
}

func (d *DatasetMMS) ParseMMSApi() *DatasetMMS {
	urlPath := "http://127.0.0.1:8383/mms"
	req, err := http.Get(urlPath)
	if err != nil {
		log.Println("Отсутствует соединение с сервером", err)
	}
	if req.StatusCode != http.StatusOK {
		return d
	}
	err = json.NewDecoder(req.Body).Decode(&d.Data)
	if err != nil {
		log.Println(err)
		return d
	}
	d.Validate()
	return d
}

func (d *DatasetMMS) Validate() {
	for i, v := range d.Data {
		if ok := v.ValidateMMS(); !ok {
			d.Data = append(d.Data[:i], d.Data[i+1:]...)
		}
	}
}

func (d *DatasetMMS) ChangeСountryMMS() {
	mms := d.Data
	alfa := alfa2.Alpha2
	for _, v := range mms {
		v.Country = alfa[v.Country]
	}
}

func (d *DatasetMMS) SortedMMSProvider() []models.MMS {
	var data []models.MMS
	byteSms, err := json.Marshal(d.Data)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(byteSms, &data)
	sort.Slice(data, func(i, j int) bool { return data[i].Provider < data[j].Provider })
	return data
}

func (d *DatasetMMS) SortedMMSCountry() []models.MMS {
	var data []models.MMS
	byteMms, err := json.Marshal(d.Data)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(byteMms, &data)
	sort.Slice(data, func(i, j int) bool { return data[i].Country < data[j].Country })
	return data
}

package store

import (
	"encoding/json"
	"github.com/DmitryOdintsov/workingWithGit/internal/models"
	"github.com/DmitryOdintsov/workingWithGit/pkg/alfa2"
	"log"
	"sort"
	"sync"
)

type DatasetSMS struct {
	Data []*models.SMS
	sync.Mutex
}

func NewDatasetSMS() *DatasetSMS {
	return &DatasetSMS{}
}

func (d *DatasetSMS) ParseInStruct(data [][]string) *DatasetSMS {
	for _, line := range data {
		sms := models.SMS{}
		if len(line) == 4 {
			for j, field := range line {
				switch j {
				case 0:
					sms.Country = field
				case 1:
					sms.Bandwidth = field
				case 2:
					sms.ResponseTime = field
				case 3:
					sms.Provider = field
				}
			}
			d.Data = append(d.Data, &sms)
			d.Validate()

		}
	}
	return d
}

func (d *DatasetSMS) Validate() {
	for i, v := range d.Data {
		if ok := v.ValidateSMS(); !ok {
			d.Data = append(d.Data[:i], d.Data[i+1:]...)
		}
	}
}

func (d *DatasetSMS) ChangeСountrySMS() {
	sms := d.Data
	alfa := alfa2.Alpha2
	for _, v := range sms {
		v.Country = alfa[v.Country]
	}
}

func (d *DatasetSMS) SortedSMSProvider() []models.SMS {
	var data []models.SMS
	byteSms, err := json.Marshal(d.Data)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(byteSms, &data)
	sort.Slice(data, func(i, j int) bool { return data[i].Provider < data[j].Provider })
	return data
}

func (d *DatasetSMS) SortedSMSCountry() []models.SMS {
	var data []models.SMS
	byteSms, err := json.Marshal(d.Data)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(byteSms, &data)
	sort.Slice(data, func(i, j int) bool { return data[i].Country < data[j].Country })
	return data
}

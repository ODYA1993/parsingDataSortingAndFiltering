package store

import (
	"encoding/json"
	"github.com/DmitryOdintsov/workingWithGit/internal/models"
	"log"
	"sort"
	"strconv"
	"sync"
)

type DatasetEmail struct {
	Data []*models.EmailData
	sync.Mutex
}

func NewDatasetEmail() *DatasetEmail {
	return &DatasetEmail{}
}

func (d *DatasetEmail) ParseInStruct(data [][]string) *DatasetEmail {
	var sliceEmail []*models.EmailData
	for _, line := range data {
		if len(line) == 3 {
			email := models.EmailData{}
			for j, field := range line {
				switch j {
				case 0:
					email.Country = field
				case 1:
					email.Provider = field
				case 2:
					fieldInt, err := strconv.Atoi(field)
					if err != nil {
						log.Println(err)
					}
					email.DeliveryTime = fieldInt
				}
			}
			if ok := email.ValidateEmail(); ok {
				sliceEmail = append(sliceEmail, &email)
			}
			d.Data = sliceEmail
		}
	}
	return d
}

func (d *DatasetEmail) SortedEmailTime() []models.EmailData {
	var data []models.EmailData
	byteSms, err := json.Marshal(d.Data)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(byteSms, &data)
	sort.Slice(data, func(i, j int) bool { return data[i].DeliveryTime < data[j].DeliveryTime })
	return data
}

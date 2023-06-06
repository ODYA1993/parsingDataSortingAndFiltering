package store

import (
	"github.com/DmitryOdintsov/workingWithGit/internal/models"
	"log"
	"strconv"
)

type DatasetVoice struct {
	Data []*models.VoiceCall
	//sync.Mutex
}

func NewDatasetVoice() *DatasetVoice {
	return &DatasetVoice{}
}

func (d *DatasetVoice) ParseInStruct(data [][]string) {
	for _, line := range data {
		if len(line) == 8 {
			voice := models.VoiceCall{}
			for j, field := range line {
				switch j {
				case 0:
					voice.Country = field
				case 1:
					fieldInt, err := strconv.Atoi(field)
					if err != nil {
						log.Println(err)
					}
					voice.Load = fieldInt
				case 2:
					fieldInt, err := strconv.Atoi(field)
					if err != nil {
						log.Println(err)
					}
					voice.ResponseTime = fieldInt
				case 3:
					voice.Provider = field
				case 4:
					fieldFloat, err := strconv.ParseFloat(field, 64)
					if err != nil {
						log.Println(err)
					}
					voice.ConnectionStability = fieldFloat
				case 5:
					fieldInt, err := strconv.Atoi(field)
					if err != nil {
						log.Println(err)
					}
					voice.TTFB = fieldInt
				case 6:
					fieldInt, err := strconv.Atoi(field)
					if err != nil {
						log.Println(err)
					}
					voice.PurityOfCommunication = fieldInt
				case 7:
					fieldInt, err := strconv.Atoi(field)
					if err != nil {
						log.Println(err)
					}
					voice.MedianCallDuration = fieldInt
				}
			}
			if ok := voice.ValidateVoice(); ok {
				d.Data = append(d.Data, &voice)
			}
		}
	}
}

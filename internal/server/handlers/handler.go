package handler

import (
	"encoding/json"
	"github.com/DmitryOdintsov/workingWithGit/internal/models"
	"github.com/DmitryOdintsov/workingWithGit/internal/store"
	"log"
	"net/http"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) GetMMS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlPath := "http://127.0.0.1:8383/mms"
	req, err := http.Get(urlPath)
	if err != nil {
		log.Println(err)
	}
	datSet := store.NewDataset()
	if req.StatusCode != http.StatusOK {
		json.NewEncoder(w).Encode(datSet)
	}
	var data []*models.Data
	err = json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(datSet)
	}

	datSet.Data = append(datSet.Data, data...)
	datSet.Validate()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(datSet)
	return
}

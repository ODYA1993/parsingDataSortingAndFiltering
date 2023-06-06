package handlers

import (
	"encoding/json"
	"github.com/DmitryOdintsov/workingWithGit/internal/service"
	"github.com/DmitryOdintsov/workingWithGit/internal/store"
	"log"
	"net/http"
)

type Handler struct{}

func NewHandlers() *Handler {
	return &Handler{}
}

func (h *Handler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resultT := service.NewResultT(false)
	resultSet := service.NewResultSetT()
	_, err := store.GetResultData(resultSet)
	if err != nil {
		log.Println(err)
		resultT.Error = err.Error()
	} else {
		resultT.Status = true
		resultT.Data = *resultSet
	}
	byteResult, err := json.Marshal(resultT)
	w.WriteHeader(http.StatusOK)
	w.Write(byteResult)
}

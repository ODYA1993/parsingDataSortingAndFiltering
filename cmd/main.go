package main

import (
	"github.com/DmitryOdintsov/workingWithGit/internal/server"
	"github.com/DmitryOdintsov/workingWithGit/internal/server/handlers"
	"github.com/DmitryOdintsov/workingWithGit/internal/service"
	"github.com/DmitryOdintsov/workingWithGit/internal/store"
)

var (
	pathSMS     = "./src/simulator/sms.data"
	pathVoice   = "./src/simulator/voice.data"
	pathEmail   = "./src/simulator/email.data"
	pathBilling = "./src/simulator/billing.data"
)

func main() {
	hand := handlers.NewHandlers()
	resultSet := service.NewResultSetT()
	store.GetResultData(resultSet)
	server.Run(hand)
}

package models

import (
	"github.com/DmitryOdintsov/workingWithGit/pkg/alfa2"
)

type SMS struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

type MMS struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

type VoiceCall struct {
	Country               string  `json:"country"`
	Load                  int     `json:"load"`
	ResponseTime          int     `json:"response_time"`
	Provider              string  `json:"provider"`
	ConnectionStability   float64 `json:"connectionStability"`
	TTFB                  int     `json:"TTFB"`
	PurityOfCommunication int     `json:"purityOfCommunication"`
	MedianCallDuration    int     `json:"medianCallDuration"`
}

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

func (s *SMS) ValidateSMS() bool {
	switch s.Provider {
	case "Topolo", "Rond", "Kildy":
		if ok := alfa2.IsAlpha2(s.Country); ok {
			return true
		}
	}
	return false
}

func (m *MMS) ValidateMMS() bool {
	switch m.Provider {
	case "Topolo", "Rond", "Kildy":
		if ok := alfa2.IsAlpha2(m.Country); ok {
			return true
		}
	}
	return false
}

func (v *VoiceCall) ValidateVoice() bool {
	switch v.Provider {
	case "TransparentCalls", "E-Voice", "JustPhone":
		if ok := alfa2.IsAlpha2(v.Country); ok {
			return true
		}
	}
	return false
}

func (e *EmailData) ValidateEmail() bool {
	switch e.Provider {
	case "Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live",
		"RediffMail", "GMX", "Proton Mail", "Yandex", "Mail.ru":
		if ok := alfa2.IsAlpha2(e.Country); ok {
			return true
		}
	}
	return false
}

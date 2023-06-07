package models

import (
	"github.com/DmitryOdintsov/workingWithGit/pkg/alfa2"
)

type SMS struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

type MMS struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

type VoiceCall struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float64 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
	//Country               string  `json:"country"`
	//Load                  int     `json:"load"`
	//ResponseTime          int     `json:"response_time"`
	//Provider              string  `json:"provider"`
	//ConnectionStability   float64 `json:"connection_stability"`
	//TTFB                  int     `json:"ttfb"`
	//PurityOfCommunication int     `json:"purity_of_communication"`
	//MedianCallDuration    int     `json:"median_call_duration"`
}

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

type BillingData struct {
	CreateCustomer bool `json:"createCustomer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraudControl"`
	CheckoutPage   bool `json:"checkoutPage"`
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

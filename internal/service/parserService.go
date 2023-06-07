package service

import (
	"encoding/json"
	"fmt"
	"github.com/DmitryOdintsov/workingWithGit/internal/models"
	"github.com/DmitryOdintsov/workingWithGit/internal/store"
	"github.com/DmitryOdintsov/workingWithGit/pkg/parse"
	"log"
)

var (
	pathSMS     = "./simulator/sms.data"
	pathVoice   = "./simulator/voice.data"
	pathEmail   = "./simulator/email.data"
	pathBilling = "./simulator/billing.data"
)

type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

func NewResultT(stat bool) *ResultT {
	return &ResultT{Status: stat}
}

type ResultSetT struct {
	SMS       [][]models.SMS                  `json:"sms"`
	MMS       [][]models.MMS                  `json:"mms"`
	VoiceCall []models.VoiceCall              `json:"voice_call"`
	Email     map[string][][]models.EmailData `json:"email"`
	Billing   models.BillingData              `json:"billing"`
	Support   [][]int                         `json:"support"`
	Incidents []models.IncidentData           `json:"incident"`
}

func NewResultSetT() *ResultSetT {
	return &ResultSetT{}
}

func (r *ResultSetT) GetSMS() error {
	dataSetSMS := store.NewDatasetSMS()
	smsData := parse.ParsingFile(pathSMS)
	dataSetSMS.ParseInStruct(smsData).ChangeСountrySMS()

	sortProvider := dataSetSMS.SortedSMSProvider()
	sortCountry := dataSetSMS.SortedSMSCountry()
	r.SMS = append(r.SMS, sortProvider)
	r.SMS = append(r.SMS, sortCountry)
	fmt.Println("SMS данные из GET-запроса:")
	for _, v := range r.SMS {
		y, _ := json.Marshal(v)
		fmt.Println(string(y))
	}
	return nil
}

func (r *ResultSetT) GetMMS() error {
	dataSetMMS := store.NewDatasetMMS()
	dataSetMMS.ParseMMSApi().ChangeСountryMMS()

	sortProvider := dataSetMMS.SortedMMSProvider()
	sortCountry := dataSetMMS.SortedMMSCountry()
	r.MMS = append(r.MMS, sortProvider)
	r.MMS = append(r.MMS, sortCountry)
	fmt.Println("MMS данные из GET-запроса:")
	for _, v := range r.MMS {
		y, _ := json.Marshal(v)
		fmt.Println(string(y))
	}
	return nil
}

func (r *ResultSetT) GetVoice() error {
	dataSetVoice := store.NewDatasetVoice()
	voice := parse.ParsingFile(pathVoice)
	dataSetVoice.ParseInStruct(voice)

	var data []models.VoiceCall
	byteSms, err := json.Marshal(dataSetVoice.Data)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(byteSms, &data)
	if err != nil {
		log.Println(err)
	}
	r.VoiceCall = append(r.VoiceCall, data...)
	fmt.Println("VOICE данные из файла:")
	for _, v := range r.VoiceCall {
		y, _ := json.Marshal(v)
		fmt.Println(string(y))
	}
	return nil
}

func (r *ResultSetT) GetEmail() error {
	email := parse.ParsingFile(pathEmail)
	dataEmail := store.NewDatasetEmail()
	data := dataEmail.ParseInStruct(email).SortedEmailTime()

	r.Email = map[string][][]models.EmailData{}

	for _, v := range data {
		r.Email[v.Country] = append(r.Email[v.Country], data[:3])
		r.Email[v.Country] = append(r.Email[v.Country], data[len(data)-3:])
		break
	}
	fmt.Println("Email данные из файла:")
	for i, v := range r.Email {
		for _, j := range v {
			fmt.Println(i, j)
		}
	}
	return nil
}

func (r *ResultSetT) GetBilling() error {
	dataBilling := store.NewDatasetBilling()
	_, billing := parse.ParseFileBit(pathBilling)
	dataBilling.ParseInStruct(billing)
	for _, v := range dataBilling.Data {
		r.Billing = *v
	}
	fmt.Println("Billing данные из файла:\n", r.Billing)
	return nil
}

func (r *ResultSetT) GetSupport() error {
	dataSupport := store.NewDatasetSupport()
	dataSupport.GetSupportAPI()
	avgPerTicket := 60 / 18
	for _, v := range dataSupport.Data {
		if v.ActiveTickets < 9 {
			r.Support = append(r.Support, []int{1, v.ActiveTickets * avgPerTicket})
		} else if v.ActiveTickets > 9 && v.ActiveTickets < 16 {
			r.Support = append(r.Support, []int{2, v.ActiveTickets * avgPerTicket})
		} else {
			r.Support = append(r.Support, []int{3, v.ActiveTickets * avgPerTicket})
		}
	}
	fmt.Println("Support данные из GET-запроса:\n", r.Support)
	return nil
}

func (r *ResultSetT) GetIncident() error {
	dataIncident := store.NewDatasetIncident()
	incident := dataIncident.GetIncident().SortedIncidentStatus()
	//sortData := sortedIncidentStatus(incidents)
	r.Incidents = append(r.Incidents, incident...)
	fmt.Println("Incident данные из GET-запроса:")
	for _, v := range r.Incidents {
		y, _ := json.Marshal(v)
		fmt.Println(string(y))
	}
	return nil
}

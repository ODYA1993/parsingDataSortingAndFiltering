package store

import (
	"log"
)

type IResult interface {
	GetSMS() error
	GetMMS() error
	GetVoice() error
	GetEmail() error
	GetBilling() error
	GetSupport() error
	GetIncident() error
}

func GetResultData(data IResult) (IResult, error) {
	err := data.GetSMS()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = data.GetMMS()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = data.GetVoice()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = data.GetEmail()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = data.GetBilling()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = data.GetSupport()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = data.GetIncident()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}

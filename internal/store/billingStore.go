package store

import (
	"github.com/DmitryOdintsov/workingWithGit/internal/models"
)

type DatasetBilling struct {
	Data []*models.BillingData
}

func NewDatasetBilling() *DatasetBilling {
	return &DatasetBilling{}
}

func (d *DatasetBilling) ParseInStruct(data []bool) {
	billing := new(models.BillingData)
	for i, v := range data {
		switch i {
		case 0:
			billing.CreateCustomer = v
		case 1:
			billing.Purchase = v
		case 2:
			billing.Payout = v
		case 3:
			billing.Recurring = v
		case 4:
			billing.FraudControl = v
		case 5:
			billing.CheckoutPage = v
		}
	}
	d.Data = append(d.Data, billing)
}

package paymill

import (
	"net/url"
)

type Offer struct {
	Id                string   "id"
	Name              string   "name"
	Amount            Numeric  "amount"
	Currency          string   "currency"
	Interval          Interval "interval"
	TrialPeriodDays   Numeric  "trial_period_days"
	CreatedAt         Numeric  "created_at"
	UpdatedAt         Numeric  "updated_at"
	SubscriptionCount struct {
		Active   Numeric `json:"active"`
		Inactive Numeric `json:"inactive"`
	} `json:"subscription_count"`
}

func (o *Offer) ResourceId() string {
	return o.Id
}

func (o *Offer) ResourceName() string {
	return "offers"
}

func (o *Offer) CreateValues() url.Values {
	return url.Values{
		"amount":            []string{string(o.Amount)},
		"currency":          []string{o.Currency},
		"interval":          []string{string(o.Interval)},
		"name":              []string{o.Name},
		"trial_period_days": []string{string(o.TrialPeriodDays)},
	}
}

func (o *Offer) UpdateValues() url.Values {
	return url.Values{
		"name": []string{o.Name},
	}
}

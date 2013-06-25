package paymill

type Transaction struct {
	Id               string        `json:"id"`
	Amount           string        `json:"amount"`
	OriginAmount     int           `json:"origin_amount"`
	Status           string        `json:"status"`
	Currency         string        `json:"currency"`
	Description      string        `json:"description"`
	LiveMode         bool          `json:"livemode"`
	Refunds          []*Refund     `json:"refunds"`
	Payment          *Payment      `json:"payment"`
	Client           *Client       `json:"client"`
	Preauthorization interface{}   `json:"preauthorization"`
	CreatedAt        int           `json:"created_at"`
	UpdatedAt        int           `json:"updated_at"`
	ResponseCode     int           `json:"response_code"`
	ShortId          string        `json:"short_id"`
	Invoices         []interface{} `json:"invoices"`
	Fees             []interface{} `json:"fees"`
}

func (t *Transaction) ResourceName() string {
	return "transactions"
}

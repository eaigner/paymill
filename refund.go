package paymill

type Refund struct {
	Id           string      `json:"id"`
	Amount       string      `json:"amount"`
	Status       string      `json:"status"`
	Description  string      `json:"description"`
	LiveMode     bool        `json:"livemode"`
	CreatedAt    int         `json:"created_at"`
	UpdatedAt    int         `json:"updated_at"`
	ResponseCode int         `json:"response_code"`
	Transaction  interface{} `json:"transaction"`
}

func (r *Refund) ResourceName() string {
	return "refunds"
}

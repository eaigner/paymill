package paymill

type Preauthorization struct {
	Id        string   `json:"id"`
	Amount    string   `json:"amount"`
	Currency  string   `json:"currency"`
	Status    string   `json:"status"`
	LiveMode  bool     `json:"livemode"`
	CreatedAt int      `json:"created_at"`
	UpdatedAt int      `json:"updated_at"`
	Payment   *Payment `json:"payment"`
	Client    *Client  `json:"client"`
}

func (p *Preauthorization) ResourceName() string {
	return "preauthorizations"
}

package paymill

type PaymentType string

const (
	PaymentTypeCredit PaymentType = "creditcard"
	PaymentTypeDebit  PaymentType = "debit"
)

type CardType string

const (
	CardTypeVisa       CardType = "visa"
	CardTypeMastercard CardType = "mastercard"
)

type Payment struct {
	Id          string      `json:"id"`
	Type        PaymentType `json:"type"`
	Client      string      `json:"client"`
	CardType    CardType    `json:"card_type"` // Credit payment fields
	Country     string      `json:"country"`
	ExpireMonth Numeric     `json:"expire_month"`
	ExpireYear  Numeric     `json:"expire_year"`
	CardHolder  string      `json:"card_holder"`
	Last4       string      `json:"last4"`
	Code        string      `json:"code"` // Debit payment fields
	Account     string      `json:"account"`
	Holder      string      `json:"holder"`
	CreatedAt   Numeric     `json:"created_at"`
	UpdatedAt   Numeric     `json:"updated_at"`
}

func (p *Payment) ResourceName() string {
	return "payments"
}

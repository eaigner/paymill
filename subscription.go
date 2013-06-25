package paymill

type Subscription struct {
	Id                string   `jons:"id"`
	Offer             *Offer   `jons:"offer"`
	LiveMode          bool     `jons:"livemode"`
	CancelAtPeriodEnd bool     `jons:"cancel_at_period_end"`
	TrialStart        int      `jons:"trial_start"`
	TrialEnd          int      `jons:"trial_end"`
	NextCaptureAt     int      `jons:"next_capture_at"`
	CreatedAt         int      `jons:"created_at"`
	UpdatedAt         int      `jons:"updated_at"`
	CanceledAt        int      `jons:"canceled_at"`
	Payment           *Payment `jons:"payment"`
	Client            *Client  `jons:"client"`
}

func (s *Subscription) ResourceName() string {
	return "subscriptions"
}

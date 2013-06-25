package paymill

import (
	"net/url"
	"time"
)

type Client struct {
	Id           string     `json:"id"`
	Email        string     `json:"email"`
	Description  string     `json:"description"`
	CreatedAt    Numeric    `json:"created_at"`
	UpdatedAt    Numeric    `json:"updated_at"`
	Payment      []*Payment `json:"payment"`
	Subscription string     `json:"subscription"`
}

func (c *Client) ResourceId() string {
	return c.Id
}

func (c *Client) ResourceName() string {
	return "clients"
}

func (c *Client) CreateValues() url.Values {
	return url.Values{
		"email":       []string{c.Email},
		"description": []string{c.Description},
	}
}

func (c *Client) UpdateValues() url.Values {
	return c.CreateValues()
}

func (c *Client) Created() time.Time {
	return time.Unix(c.CreatedAt.Int64(), 0)
}

func (c *Client) Updated() time.Time {
	return time.Unix(c.UpdatedAt.Int64(), 0)
}

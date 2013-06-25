package paymill

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Paymill struct {
	PrivateKey string
}

func NewPaymill(key string) *Paymill {
	return &Paymill{
		PrivateKey: key,
	}
}

func (p *Paymill) URL() string {
	return "https://api.paymill.com/v2"
}

func (p *Paymill) Create(r Resource) error {
	return p.Request("POST", r.ResourceName()+"?"+r.CreateValues().Encode(), &Result{r})
}

func (p *Paymill) Get(r Resource) error {
	return p.Request("GET", r.ResourceName()+"/"+r.ResourceId(), &Result{r})
}

func (p *Paymill) Update(r Resource) error {
	return p.Request("PUT", r.ResourceName()+"/"+r.ResourceId()+"?"+r.UpdateValues().Encode(), &Result{r})
}

func (p *Paymill) Delete(r Resource) error {
	return p.Request("DELETE", r.ResourceName()+"/"+r.ResourceId(), nil)
}

func (p *Paymill) Request(method, pathAndQuery string, out interface{}) error {
	fmt.Println("REQ:", method, p.URL()+"/"+pathAndQuery) // TODO: remove
	req, err := http.NewRequest(method, p.URL()+"/"+pathAndQuery, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(p.PrivateKey, "")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return errors.New(http.StatusText(resp.StatusCode))
	}
	if out != nil {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(b)) // TODO: remove
		return json.Unmarshal(b, &out)
	}
	return nil
}

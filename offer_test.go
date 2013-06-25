package paymill

import (
	"testing"
)

func TestOffer(t *testing.T) {
	pm := paymill()
	o := &Offer{
		Amount:          "99",
		Interval:        "1 MONTH",
		Name:            "TestSubscription",
		TrialPeriodDays: "14",
		Currency:        "EUR",
	}

	// Create
	err := pm.Create(o)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(o)

	if o.Id == "" {
		t.Fatal()
	}

	// Update
	o.Name = "TestSubscription2"
	err = pm.Update(o)
	if err != nil {
		t.Fatal(err)
	}

	// Get
	o2 := &Offer{
		Id: o.Id,
	}
	err = pm.Get(o2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(o2)

	if o2.Name != "TestSubscription2" {
		t.Fatal()
	}

	// Delete
	err = pm.Delete(o)
	if err != nil {
		t.Fatal(err)
	}

	// Try to get deleted offer
	err = pm.Get(o2)
	if err == nil {
		t.Fatal()
	}
}

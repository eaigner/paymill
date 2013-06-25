package paymill

import (
	"reflect"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	pm := paymill()
	now := time.Now()
	c := &Client{
		Email:       "a@b.com",
		Description: "Test Client",
	}

	// Create
	err := pm.Create(c)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(c)

	if c.Id == "" {
		t.Fatal()
	}
	if c.Created().Unix() < now.Unix()-100 {
		t.Fatal(now.Unix())
	}
	if c.Updated().Unix() < now.Unix()-100 {
		t.Fatal(now.Unix())
	}

	// Get
	c2 := &Client{
		Id: c.Id,
	}
	err = pm.Get(c2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(c2)

	if !reflect.DeepEqual(c, c2) {
		t.Fatal()
	}

	// Update (create a new object to verify if fields are set correctly)
	c3 := &Client{
		Id:          c.Id,
		Email:       "c@d.com",
		Description: "desc",
	}
	err = pm.Update(c3)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(c3)

	if !c3.Created().Equal(c.Created()) {
		t.Fatal()
	}

	// Delete
	err = pm.Delete(c)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(c)

	// Try to get deleted client
	c4 := &Client{
		Id: c.Id,
	}
	err = pm.Get(c4)
	if err == nil {
		t.Fatal()
	}

	t.Log(c4)

	if c4.Email != "" {
		t.Fatal()
	}
}

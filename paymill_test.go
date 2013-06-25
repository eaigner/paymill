package paymill

import (
	"os"
)

func paymill() *Paymill {
	return &Paymill{
		PrivateKey: os.Getenv("PAYMILL_KEY"),
	}
}

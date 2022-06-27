package device

import (
	"errors"

	"github.com/94peter/sterna/util"
)

type NewDevice struct {
	Mac		string
	Model	string
	Description		string
	NewID	string	
}

func (d *NewDevice) Valid() error {
	if !util.IsMAC(d.Mac) {
		return errors.New("invalid mac (400)")
	}

	return nil
}
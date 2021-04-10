package core

import (
	"github.com/denisbrodbeck/machineid"
	"github.com/ip-rw/ransack/ent"
)

func LookupMachine() (*ent.Machine, error) {
	id, err := machineid.ID()
	if err != nil {
		return nil, err
	}

}

package db

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Machine struct {
	gorm.Model
	MachineID string
	//Prefixes  []net.IPNet

	Keys []Key
}

type Key struct {
	gorm.Model
	Data      []byte
	Path      string
	UserID    uint
	MachineID uint
}

type User struct {
	gorm.Model
	Username string
	Password string
	Machines []Machine
}

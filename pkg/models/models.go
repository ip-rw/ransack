package models

import (
	"golang.org/x/crypto/ssh"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
)

type Machine struct {
	gorm.Model
	MachineID string `gorm:"unique;"`
	Hostname  string
	IP        string
	Users     []*User     `gorm:"many2many:user_machines;"`
	Passwords []*Password `gorm:"many2many:password_machines;"`
	Keys      []*Key
}

type Login struct {
	gorm.Model
	UserID uint
}

type Key struct {
	gorm.Model
	Data []byte
	Path string

	UserID    uint
	MachineID uint

	signer ssh.Signer
}

func ParseKey(path string, data []byte) (key *Key, err error) {
	var signer ssh.Signer
	if data == nil {
		if data, err = ioutil.ReadFile(path); err != nil {
			return nil, err
		} else {
			if signer, err = ssh.ParsePrivateKey(data); err != nil {
				return nil, err
			}
		}
	}
	return &Key{
		Data:   data,
		Path:   path,
		signer: signer,
	}, nil
}

type Password struct {
	gorm.Model
	Password string

	Users    []*User    `gorm:"many2many:user_passwords;"`
	Machines []*Machine `gorm:"many2many:password_machines;"`
}

type User struct {
	gorm.Model
	Username string

	Machines  []*Machine  `gorm:"many2many:user_machines;"`
	Passwords []*Password `gorm:"many2many:user_passwords;"`
}

type Credential struct {
	gorm.Model
	UserID     uint
	User       *User
	PasswordID uint
	Password   *Password
}

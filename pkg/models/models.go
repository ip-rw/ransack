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
	Keys      []*Key      `gorm:"many2many:key_machines;"`
}

type Key struct {
	gorm.Model
	Data []byte
	Path string

	Users    []*User    `gorm:"many2many:key_users;"`
	Machines []*Machine `gorm:"many2many:key_machines;"`

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
	Keys      []*Machine  `gorm:"many2many:key_users;"`
}

type Target struct {
	gorm.Model
	HostID       uint
	Host         *Host
	CredentialID uint
	Credential   *Credential
	Valid        bool
}

type Host struct {
	gorm.Model
	IP          string
	Port        int
	Fingerprint string
}

type Credential struct {
	gorm.Model
	UserID     uint
	User       *User
	PasswordID uint
	Password   *Password
	KeyID      uint
	Key        *Key
}

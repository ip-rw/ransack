package main

import (
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"github.com/ip-rw/ransack/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func main() {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("ransack.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	// Auto Migrate
	fmt.Println(db.AutoMigrate(&models.Machine{}, &models.Password{}, &models.User{}, &models.Key{}, &models.Credential{}, &models.Login{}))
	var m models.Machine
	//var c models.Credential
	m.Hostname, _ = os.Hostname()
	m.MachineID, _ = machineid.ID()
	fmt.Println(db.Create(&m))
}

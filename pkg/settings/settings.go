package settings

import (
	"github.com/go-ini/ini"
	"log"
)

type Database struct {
	Type      string
	Username  string
	Password  string
	Name      string
	Host      string
	Port      uint16
	MaxIdle   uint16
	MaxActive uint16
}

var DatabaseSetting = &Database{}

func SetUp() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Failed to parse app.ini, %v", err)
	}

	if err := cfg.Section("database").MapTo(DatabaseSetting); err != nil {
		log.Fatalf("Failed to map to Database, %v", err)
	}
}

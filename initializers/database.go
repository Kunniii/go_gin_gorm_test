package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseInfo struct {
	host     string
	user     string
	password string
	dbname   string
	port     int
	sslmode  string
	TimeZone string
}

func (dbInfo *DatabaseInfo) toString() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", dbInfo.host, dbInfo.user, dbInfo.password, dbInfo.dbname, dbInfo.port, dbInfo.sslmode, dbInfo.TimeZone)
}

var db *gorm.DB

func ConnectDB(dbInfo DatabaseInfo) {
	var err error
	dsn := dbInfo.toString()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database!")
	}
}

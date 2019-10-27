package db

import (
	"github.com/naormalca/api-management/config"
	"github.com/naormalca/api-management/db/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
var DBService Service

func Load() {
	var myDBConf = config.Main.Database
	// setup the database
	dbc, err := newDatabase(&myDBConf)
	if err != nil {
		panic(err.Error())
	}

	// create the database service
	DBService = NewService(dbc)
}

func Close() {

}
// Returns a new Database Client Connection
func newDatabase(config *config.Database) (*gorm.DB, error) {
	db, err := gorm.Open(config.Dialect, getInfo(config))
	if err != nil {
		return nil, err
	}

	if err := initDatabase(db, config); err != nil {
		return nil, err
	}

	return db, nil
}

// initializes the database
func initDatabase(db *gorm.DB, config *config.Database) error {
	db.LogMode(config.Debug)

	// auto migrate
	models := []interface{}{
		&models.Account{},
	}
	if err := db.AutoMigrate(models...).Error; err != nil {
		return err
	}

	return nil
}

func getInfo(config *config.Database) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Source)
	return psqlInfo
}
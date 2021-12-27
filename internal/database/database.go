package database

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	stdLog "log"
	"os"
	"reflect"
	"time"
)

var (
	Db            *gorm.DB
	IsInitialized bool
)

func InitDatabase(dsn string) (*gorm.DB, error) {

	silentLogger := gormLogger.New(
		stdLog.New(os.Stdout, "\r\n", stdLog.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold: time.Second,       // Slow SQL threshold
			LogLevel:      gormLogger.Silent, // Log level
			Colorful:      false,             // Disable color
		},
	)

	// Open a connection with the database, otherwise quit the main process.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: silentLogger,
	})

	// Update the global DbHandle instance.
	Db = db

	if err == nil {
		IsInitialized = true
	}

	// Finally return the instance of the db we created.
	return Db, err
}

// Automigrates a gorm.Model interface.
// This simply calls AutoMigrate on the model argument.
// Additional logging.
func AutoMigrate(model interface{}) {
	err := Db.AutoMigrate(model)

	if err != nil {
		log.Error(fmt.Sprintf("Unable to migrate model %s", reflect.TypeOf(model)))
		log.Error(err.Error())
		return
	}

	log.Info(fmt.Sprintf("Migrated model of type %s", reflect.TypeOf(model)))
}

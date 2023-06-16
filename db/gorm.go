package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDatabaseConfig struct {
	Host                   string
	User                   string
	Password               string
	DbName                 string
	Port                   int
	SslMode                bool
	TimeZone               string
	DisableImplicitPrepare bool
}

type GormDatabase struct {
	inner *gorm.DB
}

func NewGormDatabase(config GormDatabaseConfig) *GormDatabase {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", config.Host, config.User, config.Password, config.DbName, config.Port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: config.DisableImplicitPrepare, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database")
	}
	return &GormDatabase{inner: db}
}

func (db *GormDatabase) Inner() *gorm.DB {
	return db.inner
}

func (db *GormDatabase) Migrate(obj any) error {
	return db.inner.AutoMigrate(obj)
}

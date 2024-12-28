package database

import (
	"echo-base/internal/model"
	utils "echo-base/utils/connection"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type Configuration struct {
	ConnectionIdle     int
	ConnectionMax      int
	Dsn                string
	Dialect            string
	Prefix             string
	Driver             gorm.Dialector
	ConnectionLifeTime time.Duration
}

type ConnectionInterface interface {
	Config() *Configuration
	Open() (*gorm.DB, error)
	Close() error
}

type Connection struct {
	config   *Configuration
	instance *gorm.DB
}

func NewConnection() (ConnectionInterface, error) {
	var driver gorm.Dialector

	dsn := os.Getenv(FlagDSN)
	if dsn == "" {
		dsn = DefaultDsn
	}
	dialect := os.Getenv(FlagDialect)
	if dialect == "" {
		dialect = DefaultDialect
	}
	prefix := os.Getenv(FlagPrefix)
	if prefix == "" {
		prefix = DefaultPrefix
	}
	switch dialect {
	case DialectMySql:
		driver = mysql.Open(dsn)
	case DialectMsSql:
		driver = sqlserver.Open(dsn)
	case DialectPostgres:
		driver = postgres.Open(dsn)
	case DialectSqLite:
		driver = sqlite.Open(dsn)
	default:
		return nil, errors.New("unsupported database dialect")
	}

	config := &Configuration{
		ConnectionIdle:     DefaultConIdle,
		ConnectionMax:      DefaultConMax,
		Dsn:                dsn,
		Dialect:            dialect,
		Prefix:             prefix,
		Driver:             driver,
		ConnectionLifeTime: time.Duration(DefaultConLifeTime) * time.Minute,
	}

	return &Connection{config: config}, nil
}

func (d *Connection) Open() (*gorm.DB, error) {
	connection, err := utils.ConnectionURLBuilder(os.Getenv("DB_DRIVER"))
	if err != nil {
		log.Println("fail create connection builder")
		panic(err)
	}
	orm, err := gorm.Open(postgres.Open(connection))
	if nil != err {
		return nil, err
	}

	sqlDB, err := orm.DB()
	if nil != err {
		return nil, err
	}

	d.instance = orm

	sqlDB.SetConnMaxLifetime(d.config.ConnectionLifeTime)
	sqlDB.SetMaxIdleConns(d.config.ConnectionIdle)
	sqlDB.SetMaxOpenConns(d.config.ConnectionMax)

	AutoMigrate(orm)

	return orm, nil
}

func (d *Connection) Close() error {
	if d.instance == nil {
		return errors.New("database instance is not initialized")
	}

	sqlDB, err := d.instance.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		return err
	}

	return nil
}

func (d *Connection) Config() *Configuration {
	return d.config
}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(model.User{})
}

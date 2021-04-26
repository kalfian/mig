package config

import (
	"accounting_migration/config/logger"
	"fmt"

	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Orm *gorm.DB
)

type Database struct {
	Engine            string
	Host              string
	User              string
	Password          string
	Schema            string
	Port              int
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool

	MaxIdleConns int
	MaxOpenConns int
	MaxLifeTime  int
}

func LoadDatabaseConfig() Database {
	conf := Database{
		Engine:            viper.GetString("DB_ENGINE"),
		Host:              viper.GetString("DB_HOST"),
		User:              viper.GetString("DB_USERNAME"),
		Password:          viper.GetString("DB_PASSWORD"),
		Schema:            viper.GetString("DB_SCHEMA"),
		Port:              viper.GetInt("DB_PORT"),
		ReconnectRetry:    viper.GetInt("DB_CONNECTION_RETRY"),
		ReconnectInterval: viper.GetInt64("DB_RECONNECT_INTERVAL"),
		DebugMode:         viper.GetBool("DB_DEBUG"),

		MaxIdleConns: viper.GetInt("DB_MAX_IDLE"),
		MaxOpenConns: viper.GetInt("DB_OPEN_CONNS"),
		MaxLifeTime:  viper.GetInt("DB_MAX_LIFETIME"),
	}

	return conf
}

func DBConnect() *gorm.DB {
	conf := LoadDatabaseConfig()
	dbConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Schema,
	)
	inst, err := gorm.Open(
		conf.Engine,
		dbConfig,
	)

	if err != nil {
		panic(err)
	}

	inst.LogMode(conf.DebugMode)

	inst.DB().SetMaxIdleConns(conf.MaxIdleConns)
	inst.DB().SetMaxOpenConns(conf.MaxOpenConns)
	inst.DB().SetConnMaxLifetime(time.Duration(conf.MaxLifeTime) * time.Second)

	if conf.DebugMode {
		return inst.Debug()
	}

	logger.Info().Msg("Database Started!")
	return inst
}

func OpenDB() {
	Orm = DBConnect()
}

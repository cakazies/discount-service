package infra

import (
	"discount-service/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Infra interface {
	Config() *viper.Viper
	ConnectionDB() *DB
	ConfigInfra() *models.Config
}

type infraCtx struct {
	configFilePath string
}

type DB struct {
	Read, Write *sqlx.DB
}

// New construct new infrastructure object manager
func New(configFilePath string) Infra {
	fpath, err := filepath.Abs(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(fpath); err != nil {
		log.Fatalf("config file path %s not found", configFilePath)
	}

	return &infraCtx{
		configFilePath: configFilePath,
	}
}

var (
	cfgOnce sync.Once
	cfg     *viper.Viper
)

func (c *infraCtx) Config() *viper.Viper {
	cfgOnce.Do(func() {
		viper.SetConfigFile(c.configFilePath)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}

		cfg = viper.GetViper()
	})

	return cfg
}

var (
	dbOnce sync.Once
	db     DB
)

func (c *infraCtx) ConnectionDB() *DB {
	dbOnce.Do(func() {
		pgConfig := c.Config().Sub("db_write")
		dbConfWrite := models.DBConfig{
			Host:         pgConfig.GetString("host"),
			Port:         pgConfig.GetInt("port"),
			Database:     pgConfig.GetString("database"),
			User:         pgConfig.GetString("user"),
			Password:     pgConfig.GetString("password"),
			MaxLifeTime:  pgConfig.GetInt("max_lifetime"),
			MaxIdleConns: pgConfig.GetInt("max_idle_con"),
			MaxOpenConn:  pgConfig.GetInt("max_open_con"),
		}
		db.Write = connectPostgres(dbConfWrite)

		pgConfig = c.Config().Sub("db_read")
		dbConfRead := models.DBConfig{
			Host:         pgConfig.GetString("host"),
			Port:         pgConfig.GetInt("port"),
			Database:     pgConfig.GetString("database"),
			User:         pgConfig.GetString("user"),
			Password:     pgConfig.GetString("password"),
			MaxLifeTime:  pgConfig.GetInt("max_lifetime"),
			MaxIdleConns: pgConfig.GetInt("max_idle_con"),
			MaxOpenConn:  pgConfig.GetInt("max_open_con"),
		}

		db.Read = connectPostgres(dbConfRead)
	})
	return &db
}

func connectPostgres(dbConf models.DBConfig) *sqlx.DB {
	url := generatePostgreURL(dbConf)
	dbs, err := sqlx.Open("postgres", url)
	if err != nil {
		log.Fatalf("error conection Database url : %s error : %v", url, err)
	}

	err = dbs.Ping()
	if err != nil {
		log.Fatalf("error ping Database url : %s error : %v", url, err)
	}

	return dbs
}

func generatePostgreURL(dbAcc models.DBConfig) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d  sslmode=disable extra_float_digits=-1 connect_timeout=%d", dbAcc.User, dbAcc.Password, dbAcc.Database, dbAcc.Host, dbAcc.Port, dbAcc.TimeOut)
}

var (
	configOnce sync.Once
	config     *models.Config
)

func (c *infraCtx) ConfigInfra() *models.Config {
	configOnce.Do(func() {
		resp := &models.Config{}
		config = resp

	})

	return config
}

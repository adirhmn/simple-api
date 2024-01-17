package dbpool

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	// DBPools
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type ConnectionConfig struct {
	SetMaxOpenConns int
	SetMaxIdleConns int
	Selector        Selector
}

type Selector interface {
	Next([]*sql.DB) *sql.DB
}

type DBPools struct {
	readDBs  []*sql.DB
	writeDb  *sql.DB
	selector Selector
}

func (c *DBPools) GetWriteDB() *sql.DB {
	return c.writeDb
}

func (c *DBPools) GetReadDB() *sql.DB {
	return c.selector.Next(c.readDBs)
}

func (c *DBPools) Shutdown() error {

	if err := c.writeDb.Close(); err != nil {
		return err
	}

	for _, rdb := range c.readDBs {
		if err := rdb.Close(); err != nil {
			return err
		}
	}

	return nil
}

func New(writeCfg *Config, readCfgs []*Config, cfg *ConnectionConfig) (*DBPools, error) {

	logrus.Info("create db")

	dbPools := &DBPools{}

	if cfg.Selector == nil {
		dbPools.selector = NewRoundRobinSelector()
	}

	writeDB, err := createDB(writeCfg, cfg)
	if err != nil {
		return nil, err
	}

	dbPools.writeDb = writeDB

	for _, readCfg := range readCfgs {
		readDB, err := createDB(readCfg, cfg)
		if err != nil {
			return nil, err
		}
		dbPools.readDBs = append(dbPools.readDBs, readDB)
	}

	return dbPools, nil
}

func createDB(cfg *Config, connCfg *ConnectionConfig) (*sql.DB, error) {
	datasource := "host=" + cfg.DBHost + " port=" + cfg.DBPort + " user=" + cfg.DBUser + " dbname=" + cfg.DBName + " password=" + cfg.DBPassword + " sslmode=disable"
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		return nil, err
	}

	if connCfg.SetMaxOpenConns > 0 {
		db.SetMaxOpenConns(connCfg.SetMaxOpenConns)
	}

	if connCfg.SetMaxIdleConns > 0 {
		db.SetMaxIdleConns(connCfg.SetMaxIdleConns)
	}

	// set how long connection before it destroy
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

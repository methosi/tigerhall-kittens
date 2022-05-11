package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/methosi/tigerhall-kittens/storage/connection"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Storage struct {
	logger logrus.FieldLogger
	db     *sqlx.DB
}

// NewStorageFromConfig returns a new Storage from config.
func NewStorageFromConfig(config *viper.Viper) (*Storage, error) {
	dbString, err := connection.NewDBStringFromConfig(config)
	if err != nil {
		return nil, err
	}
	return NewStorage(dbString)
}

// NewStorage returns a new Storage from the provides psql database string
func NewStorage(dbstring string) (*Storage, error) {
	db, err := sqlx.Connect("postgres", dbstring)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to postgres :%v :%s", err, dbstring)
	}
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(10 * time.Minute)
	return &Storage{logger: logrus.New(), db: db}, nil
}

// GetDBConn returns the underlying sql.DB object
func (s *Storage) GetDBConn() *sql.DB {
	return s.db.DB
}

package storage

import (
	"context"
	"database/sql"
)

type Storage interface {
	GetDBConn() *sql.DB

	TigerStore

	TigerSightingStore
}

type TigerStore interface {
	InsertTiger(context.Context, *Tiger) (string, error)
	RetrieveTigers(...FilterOption) ([]*Tiger, error)
}

type TigerSightingStore interface {
	InsertTigerSighting()
	RetrieveTigerSightings(...FilterOption) ([]*TigerSighting, error)
}

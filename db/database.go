package db

import (
	"context"
	"database/sql"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var dbInstance *bun.DB
var once sync.Once

func InitLogger(db *bun.DB, debug bool, level int) {
	db.AddQueryHook(
		bundebug.NewQueryHook(
			bundebug.WithEnabled(debug),
			bundebug.WithVerbose(level == 2),
		),
	)
}

func InitDB() *bun.DB {
	once.Do(func() {
		dsn := "postgres://postgres:postgres@localhost:5432/zaman_now_category?sslmode=disable"
		sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		dbInstance = bun.NewDB(sqlDb, pgdialect.New())
		InitLogger(dbInstance, true, 2)
	})

	return dbInstance
}

// GetConn return database connection instance
func GetConn() *bun.DB {
	return dbInstance
}

func OpenConnection() int {
	return dbInstance.Stats().OpenConnections
}

func QueryContext(ctx context.Context, query string) (ids []int, err error) {
	var qr *sql.Rows
	db := GetConn()
	qr, err = db.QueryContext(ctx, query)
	// handle error masih belom selesai
	if err != nil {
		// return nil, utils.WrapError(err)
		return
	}

	for qr.Next() {
		var id int
		err = qr.Scan(&id)
		if err == nil {
			ids = append(ids, id)
		}
	}
	return
}

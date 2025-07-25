package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunotel"

	"github.com/uptrace/bun"
)

type DatabaseDialect string

// default database connection shared between different stores
var defaultDB *DB

const (
	DialectPostgres DatabaseDialect = "pg"
	DialectSQLite   DatabaseDialect = "sqlite"
)

// DBConfig config of database
type DBConfig struct {
	// database vendor to use, valid value: pg, sqlite
	Dialect DatabaseDialect `json:"dialect" comment:"database vendor to use, valid value: pg, sqlite"`
	// e.g.: postgresql://starhub:starhub@localhost:5433/starhub?sslmode=disable
	DSN string `json:"dsn" comment:"e.g.: postgresql://starhub:starhub@localhost:5433/starhub?sslmode=disable"`
}

// DB is where all database operation lives
type DB struct {
	// Operator is where database access/write methods implemented
	Operator

	// the underlying *bun.DB
	BunDB *bun.DB
}

// Operator is where database access/write methods implemented
// so that we don't have to write the same method on both DB and Transaction.
type Operator struct {
	Core bun.IDB
}

func GetDB() *DB {
	return defaultDB
}

func InitDB(config DBConfig) {
	bg, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	defaultDB, err = NewDB(bg, config)
	if err != nil {
		log.Fatal("failed to initialize database", err.Error())
	}
}

// NewDB initializes a DB via config
func NewDB(ctx context.Context, config DBConfig) (db *DB, err error) {
	var bunDB *bun.DB

	switch config.Dialect {
	case DialectPostgres:
		sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.DSN)))
		bunDB = bun.NewDB(sqlDB, pgdialect.New(), bun.WithDiscardUnknownColumns())
	case DialectSQLite:
		var sqlDB *sql.DB
		sqlDB, err = sql.Open(sqliteshim.ShimName, config.DSN)
		if err != nil {
			err = fmt.Errorf("sql.Open: %w", err)
			return
		}

		// in-memory database is deleted when the connection is closed.
		// we should avoid this.
		// ref: https://www.sqlite.org/inmemorydb.html
		if strings.Contains(config.DSN, ":memory:") ||
			strings.Contains(config.DSN, "mode=memory") {
			sqlDB.SetMaxIdleConns(100)
			sqlDB.SetConnMaxLifetime(0)
		}

		// SQLite allows multiple readers but only a single writer at any one time.
		// If you have multiple connections to the same DB, you will inevitably run into database is locked.
		// This is an unavoidable consequence of sqlite's locking model,
		// and therefore something your application needs to deal with.
		// ref: https://github.com/mattn/go-sqlite3/issues/274#issuecomment-232942571
		//
		// Tip: Don't write through DB and TX at the same time.
		bunDB = bun.NewDB(sqlDB, sqlitedialect.New(), bun.WithDiscardUnknownColumns())
	default:
		err = fmt.Errorf("unknown database dialect %q", config.Dialect)
		return
	}

	bunDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.FromEnv("DB_DEBUG"),
	))
	bunDB.AddQueryHook(bunotel.NewQueryHook(
		bunotel.WithDBName("csghub"),
		bunotel.WithFormattedQueries(true),
	))

	err = bunDB.PingContext(ctx)
	if err != nil {
		err = fmt.Errorf("pinging %s database: %w", config.Dialect, err)
		return
	}

	db = &DB{
		Operator: Operator{Core: bunDB},
		BunDB:    bunDB,
	}

	db.BunDB.RegisterModel((*RepositoryTag)(nil))
	db.BunDB.RegisterModel((*CollectionRepository)(nil))
	return
}

// RunInTx runs the function in a transaction.
// If the function returns an error, the transaction is rolled back.
// Otherwise, the transaction is committed.
func (db *DB) RunInTx(ctx context.Context, fn func(ctx context.Context, tx Operator) error) error {
	return db.BunDB.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		op := Operator{Core: tx}
		return fn(ctx, op)
	})
}

// Close closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server to finish.
// It is rare to Close a DB,
// as the DB handle is meant to be long-lived and shared between many goroutines.
func (db *DB) Close() error {
	return db.BunDB.Close()
}

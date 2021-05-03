package app

import (
	"context"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

type Databases struct {
	PrimaryDB   *pg.DB
	SecondaryDB *pg.DB
}

type loggerHook struct {
	instance string
	logger   *zap.Logger
}

func NewDatabases(primaryDsn string, secondaryDsn string, logger *zap.Logger) (*Databases, error) {
	primaryDB, err := newDb(primaryDsn)
	if err != nil {
		return nil, err
	}

	primaryDB.AddQueryHook(loggerHook{"primary DB",logger})

	secondaryDB, err := getSecondaryDb(primaryDB, secondaryDsn)
	if err != nil {
		return nil, err
	}

	if secondaryDB != primaryDB {
		secondaryDB.AddQueryHook(loggerHook{"secondary DB",logger})
	}

	return &Databases{
		PrimaryDB:   primaryDB,
		SecondaryDB: secondaryDB,
	}, nil
}

func (databases *Databases) Close() error {
	if err := databases.SecondaryDB.Close(); err != nil {
		return err
	}

	if err := databases.PrimaryDB.Close(); err != nil {
		return err
	}

	return nil
}

func newDb(dsn string) (*pg.DB, error) {
	opts, err := pg.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opts)

	return db, nil
}

func getSecondaryDb(primaryDb *pg.DB, secondaryDsn string) (*pg.DB, error) {
	if secondaryDsn == "" {
		return primaryDb, nil
	}

	return newDb(secondaryDsn)
}

func (h loggerHook) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	query, err := q.FormattedQuery()

	if err == nil {
		h.logger.Debug(h.instance, zap.ByteString("query", query))
	}

	return c, nil
}

func (h loggerHook) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	return nil
}

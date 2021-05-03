package migration

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/go-pg/pg/v10"
	"os"
)

type Args struct {
	IsUpdate    bool
	IsCreate    bool
	Directory   string
	DatabaseDsn string
}

func NewArgs(opts *docopt.Opts) (*Args, error) {
	isUpdate, _ := opts.Bool("update")
	isCreate, _ := opts.Bool("create")

	directory, _ := opts.String("--dir")
	if _, err := os.Stat(directory); err != nil {
		return nil, fmt.Errorf("invalid argument --dir; %s", err)
	}

	databaseDsn, _ := opts.String("<db_dsn>")
	if databaseDsn != "" {
		if _, err := pg.ParseURL(databaseDsn); err != nil {
			return nil, fmt.Errorf("invalid parameter <db_dsn>; format: postgres://user:pass@host:port/db?option=value")
		}
	}

	return &Args{
		IsUpdate:    isUpdate,
		IsCreate:    isCreate,
		Directory:   directory,
		DatabaseDsn: databaseDsn,
	}, nil
}

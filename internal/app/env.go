package app

import (
	"go.uber.org/zap"
)

type Env struct {
	Logger       *zap.Logger
	Databases    *Databases
	Storages     *Storages
	Repositories *Repositories
	Services     *Services
}

func NewEnv(args *Args) (*Env, error) {
	logger, err := NewLogger(args.LogLevel)
	if err != nil {
		return nil, err
	}

	databases, err := NewDatabases(args.PrimaryDatabaseDSN, args.SecondaryDatabaseDSN, logger)
	if err != nil {
		return nil, err
	}

	storages, err := NewStorages(args.ImageStorageURI)
	if err != nil {
		return nil, err
	}

	repos := NewRepositories(databases)

	services := NewServices(repos, storages)

	env := Env{
		Logger:       logger,
		Databases:    databases,
		Storages:     storages,
		Repositories: repos,
		Services:     services,
	}

	return &env, nil
}

func (env *Env) Close() error {
	if err := env.Databases.Close(); err != nil {
		return nil
	}

	return nil
}

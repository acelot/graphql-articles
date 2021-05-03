package main

import (
	"context"
	"github.com/acelot/articles/internal/feature/dbmigration"
	"github.com/acelot/articles/internal/migration"
	"github.com/docopt/docopt-go"
	"github.com/go-pg/pg/v10"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const migrationFileExt string = ".sql"

const usage = `Migration

Usage:
  migration [--dir=<path>] update <db_dsn>

Options:
  -h --help        Show this screen.
     --dir=<path>  Migrations directory [default: ./migrations].`

func main() {
	opts, _ := docopt.ParseDoc(usage)

	args, err := migration.NewArgs(&opts)
	if err != nil {
		log.Fatalf("cannot parse args: %v", err)
	}

	if args.IsUpdate {
		updateCommand(args.Directory, args.DatabaseDsn)
	}
}

func updateCommand(directory string, databaseDsn string) {
	pgOpts, _ := pg.ParseURL(databaseDsn)
	pgOpts.PoolSize = 1

	// Check connection
	log.Print("connecting to DB...")

	db := pg.Connect(pgOpts)
	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("- failed: %s", databaseDsn)
	}

	log.Print("- ok")

	repo := dbmigration.NewRepository(db)

	// Check table
	log.Print("ensuring migration table existing...")

	if err := repo.EnsureTable(); err != nil {
		log.Fatalf("- failed: %v", err)
	}

	log.Print("- ok")

	// Make list of migration files
	log.Print("loading migration files from directory...")

	fileNames, err := getAllMigrationFileNames(directory)
	if err != nil {
		log.Fatalf("- failed: %v", err)
	}

	log.Printf("- %d files loaded", len(fileNames))

	// Get all applied migrations from DB
	log.Print("loading applied migrations from DB...")

	migrations, err := repo.Find(context.Background(), dbmigration.FindFilterIsAppliedOnly(true))
	if err != nil {
		log.Fatalf("- failed: %v", err)
	}

	log.Printf("- %d migrations loaded", len(migrations))

	// Check migrations integrity
	log.Print("checking migrations integrity...")

	if len(migrations) > len(fileNames) {
		log.Fatalf("- failed: the DB contains migrations that aren't in the file system")
	}

	for i, name := range fileNames[:len(migrations)] {
		if name != migrations[i].Name {
			log.Fatalf(
				`- failed: migrations applying order is violated - expected "%s", actual "%s"`,
				name,
				migrations[i].Name,
			)
		}

		log.Printf("- migration %s was applied on %s", name, migrations[i].AppliedAt.String())
	}

	// Applying migrations
	applied := 0

	for _, name := range fileNames[len(migrations):] {
		log.Printf("applying migration %s...", name)

		filePath := filepath.Join(directory, name) + migrationFileExt

		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("- failed: %v", err)
		}

		migrationItem := dbmigration.DBMigration{
			AppliedAt: nil,
			Name:      name,
		}

		if err := repo.Create(&migrationItem); err != nil {
			log.Fatalf("- failed: %v", err)
		}

		if _, err := db.Exec(string(bytes)); err != nil {
			log.Fatalf("- failed: %v", err)
		}

		appliedAt := time.Now()
		migrationItem.AppliedAt = &appliedAt

		if err := repo.Update(&migrationItem); err != nil {
			log.Fatalf("- failed: %v", err)
		}

		log.Printf("- ok")
		applied++
	}

	log.Printf("- %d migrations applied", applied)
}

func getAllMigrationFileNames(directory string) ([]string, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return []string{}, err
	}

	var fileNames []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != migrationFileExt {
			continue
		}

		fileName := strings.TrimSuffix(file.Name(), migrationFileExt)

		fileNames = append(fileNames, fileName)
	}

	sort.Strings(fileNames)

	return fileNames, nil
}

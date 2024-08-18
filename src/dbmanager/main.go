package dbmanager

import (
	"crispy-vault/src/vault"
	"database/sql"
	"path/filepath"

	"context"
	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseManager struct {
	ctx   context.Context
	vault *vault.Vault
	DB    *sql.DB
}

func (man *DatabaseManager) SetContext(ctx context.Context) {
	man.ctx = ctx
}
func (man *DatabaseManager) Provide(vault *vault.Vault) {
	man.vault = vault
}

func (man *DatabaseManager) Bootstrap() {
	log.Print(man.vault.BaseDir)
	dbPath := filepath.Join(man.vault.BaseDir, "db.sql")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	man.DB = db

	man.setupDb()

	man.seedDb()
}

//go:embed scripts/setup.sql
var setupScript string

//go:embed scripts/seed-data.sql
var seedScript string

func (man *DatabaseManager) setupDb() {
	dieIfErr(man.DB.Exec(setupScript))

	log.Print("Database setup successfully")
}

func (man *DatabaseManager) seedDb() {
	dieIfErr(man.DB.Exec(seedScript))

	log.Print("Database seeded")
}

func dieIfErr(res sql.Result, err error) {
	if err != nil {
		log.Fatal(err)
	}
}

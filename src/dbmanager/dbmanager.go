package dbmanager

import (
	"crispy-vault/src/vault"
	"path/filepath"

	"context"
	_ "embed"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseManager struct {
	ctx   context.Context
	vault *vault.Vault
	DB    *sqlx.DB
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
	db, err := sqlx.Open("sqlite3", dbPath)
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
	_, err := man.DB.Exec(setupScript)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Database setup successfully")
}

func (man *DatabaseManager) seedDb() {
	_, err := man.DB.Exec(seedScript)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Database seeded")
}

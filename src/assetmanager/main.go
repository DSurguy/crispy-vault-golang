package assetmanager

import (
	"crispy-vault/src/dbmanager"
	"crispy-vault/src/vault"

	"context"
)

type AssetManager struct {
	ctx   context.Context
	vault *vault.Vault
	dbman *dbmanager.DatabaseManager
}

func (man *AssetManager) SetContext(ctx context.Context) {
	man.ctx = ctx
}

func (man *AssetManager) CreateAsset(name string) {
	// man.dbman.DB.Prepare("INSERT INTO asset (name, uuid, last_update) VALUES (?1, ?2 , datetime('now'))")
}

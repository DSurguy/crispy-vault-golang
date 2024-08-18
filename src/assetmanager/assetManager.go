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
func (man *AssetManager) Provide(vault *vault.Vault, dbman *dbmanager.DatabaseManager) {
	man.vault = vault
	man.dbman = dbman
}

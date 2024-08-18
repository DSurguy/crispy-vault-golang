package vault

import (
	"context"
	"log"
	"os"
	"path/filepath"
)

type Vault struct {
	ctx     context.Context
	BaseDir string
}

func (vault *Vault) SetContext(ctx context.Context) {
	vault.ctx = ctx
}

// TODO prod: Don't clear, use config to ensure vault exists
func (vault *Vault) Bootstrap() {
	dataDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}

	vault.BaseDir = filepath.Join(dataDir, "com.surgingforward.crispyvault")

	// ensure app dir exists and is empty "com.surgingforward.crispyvault"
	clearDirErr := os.RemoveAll(vault.BaseDir)
	if clearDirErr != nil {
		log.Fatal(clearDirErr)
	}
	dirSetupErr := os.MkdirAll(filepath.Join(vault.BaseDir, "assets"), 0755)
	if dirSetupErr != nil {
		log.Fatal(dirSetupErr)
	}

	log.Print("Directory created")
}

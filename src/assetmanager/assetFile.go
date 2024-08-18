package assetmanager

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type AssetFile struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Extension   string `json:"extension"`
}

type CreateAssetFilePayload struct {
	AssetUuid   string `json:"assetUuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
}

type CreateAssetFileResponse struct {
	AssetFile AssetFile `json:"assetFile"`
	Err       string    `json:"error"`
}

// TODO: transaction
func (man *AssetManager) CreateAssetFile(payload CreateAssetFilePayload) CreateAssetFileResponse {
	uuid := uuid.New().String()
	extension := filepath.Ext(payload.Path)
	destPath := filepath.Join(man.vault.BaseDir, "assets", payload.AssetUuid, fmt.Sprintf("%s.%s", uuid, extension))
	err := copyFile(payload.Path, destPath)
	if err != nil {
		return CreateAssetFileResponse{
			Err: err.Error(),
		}
	}
	_, err = man.dbman.DB.ExecContext(
		man.ctx,
		"INSERT INTO asset_file (uuid, name, description, extension, last_update) VALUES (?, ?, ?, ?, datetime('now'))",
		uuid,
		payload.Name,
		payload.Description,
		extension,
	)
	if err != nil {
		return CreateAssetFileResponse{
			Err: err.Error(),
		}
	}

	_, err = man.dbman.DB.ExecContext(
		man.ctx,
		"INSERT INTO asset_to_asset_file (asset_id, asset_file_id) VALUES (?, ?)",
		payload.AssetUuid,
		uuid,
	)
	if err != nil {
		return CreateAssetFileResponse{
			Err: err.Error(),
		}
	}
	return CreateAssetFileResponse{
		AssetFile: AssetFile{
			Uuid:        uuid,
			Name:        payload.Name,
			Description: payload.Description,
			Extension:   extension,
		},
	}
}

func copyFile(sourcePath string, destPath string) error {
	sourceFileStat, err := os.Stat(sourcePath)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", sourcePath)
	}

	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	err = dest.Chmod(0755)
	if err != nil {
		return err
	}
	defer dest.Close()
	_, err = io.Copy(dest, source)
	return err
}

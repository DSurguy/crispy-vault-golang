package assetmanager

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

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

func (man *AssetManager) CreateAssetFile(payload CreateAssetFilePayload) CreateAssetFileResponse {
	uuid := uuid.New().String()
	extension := filepath.Ext(payload.Path)
	assetPath := filepath.Join(man.vault.BaseDir, "assets", payload.AssetUuid)
	err := os.MkdirAll(assetPath, 0755)
	if err != nil {
		return CreateAssetFileResponse{
			Err: err.Error(),
		}
	}
	destPath := filepath.Join(assetPath, fmt.Sprintf("%s.%s", uuid, extension))
	err = copyFile(payload.Path, destPath)
	if err != nil {
		return CreateAssetFileResponse{
			Err: err.Error(),
		}
	}

	tx, err := man.dbman.DB.BeginTx(man.ctx, nil)
	if err != nil {
		return CreateAssetFileResponse{
			Err: err.Error(),
		}
	}
	_, err = tx.ExecContext(
		man.ctx,
		"INSERT INTO asset_file (uuid, name, description, extension, last_update) VALUES (?, ?, ?, ?, datetime('now'))",
		uuid,
		payload.Name,
		payload.Description,
		extension,
	)
	if err != nil {
		_ = tx.Rollback()
		return CreateAssetFileResponse{
			Err: err.Error(),
		}
	}

	_, err = tx.ExecContext(
		man.ctx,
		"INSERT INTO asset_to_asset_file (asset_id, asset_file_id) VALUES (?, ?)",
		payload.AssetUuid,
		uuid,
	)
	if err != nil {
		_ = tx.Rollback()
		return CreateAssetFileResponse{
			Err: err.Error(),
		}
	}
	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
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

type ListAssetFilesResponse struct {
	AssetFiles []AssetFile `json:"assetFiles"`
	Err        string      `json:"error"`
}

func (man *AssetManager) ListAssetFiles(assetUuid string, offset int) ListAssetFilesResponse {
	query := "SELECT uuid, name, description, extension " +
		"FROM asset_file " +
		"INNER JOIN asset_to_asset_file aaf ON asset_file.uuid = aaf.asset_file_id AND aaf.asset_id = ?" +
		"ORDER BY last_update DESC " +
		"LIMIT 20 OFFSET ?"
	rows, err := man.dbman.DB.Query(query, assetUuid, offset)
	if err != nil {
		return ListAssetFilesResponse{
			Err: err.Error(),
		}
	}

	files := make([]AssetFile, 0)
	for rows.Next() {
		var (
			uuid        string
			name        string
			description string
			extension   string
		)
		err = rows.Scan(&uuid, &name, &description, &extension)
		if err != nil {
			return ListAssetFilesResponse{
				Err: err.Error(),
			}
		}
		files = append(files, AssetFile{
			Uuid:        uuid,
			Name:        name,
			Description: description,
			Extension:   extension,
		})
	}
	return ListAssetFilesResponse{
		AssetFiles: files,
	}
}

type EditAssetFilePayload struct {
	AssetUuid   string `json:"assetUuid"`
	FileUuid    string `json:"fileUuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
}

type EditAssetFileResponse struct {
	AssetFile AssetFile `json:"assetFile"`
	Err       string    `json:"error"`
}

func (man *AssetManager) EditAssetFile(payload EditAssetFilePayload) EditAssetFileResponse {
	// TODO: validate that the file belongs to this asset
	var err error
	if payload.Path != "" {
		err = man.editAndReplaceFile(payload)
	} else {
		err = man.editFile(payload)
	}
	if err != nil {
		return EditAssetFileResponse{
			Err: err.Error(),
		}
	}

	file, err := man.getAssetFileFromDb(payload.FileUuid)
	if err != nil {
		return EditAssetFileResponse{
			Err: err.Error(),
		}
	}

	return EditAssetFileResponse{
		AssetFile: AssetFile{
			Uuid:        file.Uuid,
			Name:        file.Name,
			Description: file.Description,
			Extension:   file.Extension,
		},
	}
}

func (man *AssetManager) editFile(payload EditAssetFilePayload) error {
	query := "UPDATE asset_file SET " +
		"name = ?, " +
		"description = ?, " +
		"last_update = datetime('now') " +
		"WHERE uuid = ?"
	_, err := man.dbman.DB.Exec(query, payload.Name, payload.Description, payload.FileUuid)
	if err != nil {
		return err
	}
	return nil
}

func (man *AssetManager) editAndReplaceFile(payload EditAssetFilePayload) error {
	assetPath := filepath.Join(man.vault.BaseDir, "assets", payload.AssetUuid)

	// Remove the existing file first
	currentFile, err := man.getAssetFileFromDb(payload.FileUuid)
	if err != nil {
		return err
	}
	existingPath := filepath.Join(assetPath, fmt.Sprintf("%s.%s", payload.FileUuid, currentFile.Extension))
	err = os.RemoveAll(existingPath)
	if err != nil {
		return err
	}

	// Then copy the new file
	extension := filepath.Ext(payload.Path)
	destPath := filepath.Join(assetPath, fmt.Sprintf("%s.%s", payload.FileUuid, extension))
	err = copyFile(payload.Path, destPath)
	if err != nil {
		return err
	}

	// last update the database
	query := "UPDATE asset_file SET " +
		"name = ?, " +
		"description = ?, " +
		"extension = ?, " +
		"last_update = datetime('now') " +
		"WHERE uuid = ?"
	_, err = man.dbman.DB.Exec(query, payload.Name, payload.Description, extension, payload.FileUuid)
	if err != nil {
		return err
	}
	return nil
}

func (man *AssetManager) getAssetFileFromDb(fileUuid string) (*AssetFile, error) {
	row := man.dbman.DB.QueryRow("SELECT uuid, name, description, extension FROM asset_file WHERE uuid = ?", fileUuid)
	var (
		uuid        string
		name        string
		description string
		extension   string
	)
	err := row.Scan(&uuid, &name, &description, &extension)
	if err != nil {
		return nil, err
	}

	return &AssetFile{
		Uuid:        uuid,
		Name:        name,
		Description: description,
		Extension:   extension,
	}, nil
}

type DeleteAssetFilePayload struct {
	AssetUuid string `json:"assetUuid"`
	FileUuid  string `json:"fileUuid"`
}

type DeleteAssetFileResponse struct {
	Err string `json:"error"`
}

// TODO: cascade
func (man *AssetManager) DeleteAssetFile(payload DeleteAssetFilePayload) DeleteAssetFileResponse {
	currentFile, err := man.getAssetFileFromDb(payload.FileUuid)
	if err != nil {
		return DeleteAssetFileResponse{
			Err: err.Error(),
		}
	}

	// Remove from DB
	_, err = man.dbman.DB.ExecContext(man.ctx, "DELETE FROM asset_file WHERE uuid = ?", payload.FileUuid)
	if err != nil {
		return DeleteAssetFileResponse{
			Err: err.Error(),
		}
	}

	// Remove from file system
	assetPath := filepath.Join(man.vault.BaseDir, "assets", payload.AssetUuid)
	existingPath := filepath.Join(assetPath, fmt.Sprintf("%s.%s", payload.FileUuid, currentFile.Extension))
	err = os.RemoveAll(existingPath)
	if err != nil {
		return DeleteAssetFileResponse{
			Err: err.Error(),
		}
	}

	return DeleteAssetFileResponse{}
}

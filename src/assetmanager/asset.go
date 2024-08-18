package assetmanager

import "github.com/google/uuid"

type Asset struct {
	Uuid       string `json:"uuid"`
	Name       string `json:"name"`
	LastUpdate string `json:"lastUpdate"`
}

type CreateAssetResponse struct {
	Id  string `json:"id"`
	Err string `json:"error"`
}

func (man *AssetManager) CreateAsset(name string) CreateAssetResponse {
	id := uuid.New().String()
	_, err := man.dbman.DB.ExecContext(
		man.ctx,
		"INSERT INTO asset (name, uuid, last_update) VALUES (?1, ?2 , datetime('now'))",
		name,
		id,
	)
	if err != nil {
		return CreateAssetResponse{
			Err: err.Error(),
		}
	}
	return CreateAssetResponse{
		Id: id,
	}
}

type ListAssetsResponse struct {
	Assets []Asset `json:"assets"`
	Err    string  `json:"error"`
}

func (man *AssetManager) ListAssets() ListAssetsResponse {
	rows, err := man.dbman.DB.Query(
		"SELECT name, uuid, last_update FROM asset ORDER BY last_update DESC LIMIT 20",
	)
	if err != nil {
		return ListAssetsResponse{
			Err: err.Error(),
		}
	}

	assets := make([]Asset, 0)
	for rows.Next() {
		var name string
		var rowUuid string
		var lastUpdate string
		if err := rows.Scan(&name, &rowUuid, &lastUpdate); err != nil {
			return ListAssetsResponse{
				Err: err.Error(),
			}
		}
		var asset = Asset{
			Name:       name,
			Uuid:       rowUuid,
			LastUpdate: lastUpdate,
		}
		assets = append(assets, asset)
	}

	return ListAssetsResponse{
		Assets: assets,
	}
}

type GetAssetResponse struct {
	Asset Asset  `json:"asset"`
	Err   string `json:"error"`
}

func (man *AssetManager) GetAsset(uuid string) GetAssetResponse {
	row := man.dbman.DB.QueryRow("SELECT name, uuid, last_update FROM asset WHERE uuid = ?1", uuid)
	var name string
	var rowUuid string
	var lastUpdate string
	if err := row.Scan(&name, &rowUuid, &lastUpdate); err != nil {
		return GetAssetResponse{
			Err: err.Error(),
		}
	}
	asset := Asset{
		Name:       name,
		Uuid:       uuid,
		LastUpdate: lastUpdate,
	}
	return GetAssetResponse{
		Asset: asset,
	}
}

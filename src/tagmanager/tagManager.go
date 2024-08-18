package tagmanager

import (
	"context"
	"crispy-vault/src/dbmanager"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type TagManager struct {
	ctx   context.Context
	dbman *dbmanager.DatabaseManager
}

func (man *TagManager) SetContext(ctx context.Context) {
	man.ctx = ctx
}
func (man *TagManager) Provide(dbman *dbmanager.DatabaseManager) {
	man.dbman = dbman
}

type TagSearchPayload struct {
	Search string   `json:"search"`
	Omit   []string `json:"omit"`
}

type TagSearchResponse struct {
	Tags []string `json:"tags"`
	Err  string   `json:"error"`
}

func (man *TagManager) TagSearch(payload TagSearchPayload) TagSearchResponse {
	var (
		rows *sql.Rows
		err  error
	)
	if len(payload.Omit) > 0 {
		arg := map[string]interface{}{
			"match": payload.Search,
			"omit":  payload.Omit,
		}
		query, args, err := sqlx.Named("SELECT text FROM tag_fts WHERE text MATCH :match AND text NOT IN (:omit) ORDER BY rank LIMIT 20", arg)
		if err != nil {
			return TagSearchResponse{
				Err: err.Error(),
			}
		}
		query, args, err = sqlx.In(query, args...)
		if err != nil {
			return TagSearchResponse{
				Err: err.Error(),
			}
		}
		query = man.dbman.DB.Rebind(query)
		rows, err = man.dbman.DB.Query(query, args...)
		if err != nil {
			return TagSearchResponse{
				Err: err.Error(),
			}
		}
	} else {
		rows, err = man.dbman.DB.Query("SELECT text FROM tag_fts WHERE text MATCH ? ORDER BY rank LIMIT 20", payload.Search)
		if err != nil {
			return TagSearchResponse{
				Err: err.Error(),
			}
		}
	}

	tags := make([]string, 0)
	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return TagSearchResponse{
				Err: err.Error(),
			}
		}
		tags = append(tags, tag)
	}

	return TagSearchResponse{
		Tags: tags,
	}
}

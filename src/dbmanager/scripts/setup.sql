-- name: create-asset-table
CREATE TABLE IF NOT EXISTS asset (
    uuid TEXT UNIQUE ON CONFLICT ROLLBACK NOT NULL,
    name TEXT NOT NULL,
    last_update TEXT
);

-- name: create-asset-fts-table
CREATE VIRTUAL TABLE IF NOT EXISTS asset_fts using fts5(uuid, name, tokenize=trigram, content=asset, content_rowid=rowid);

-- name: create-asset-insert-trigger
CREATE TRIGGER IF NOT EXISTS asset_insert AFTER INSERT ON asset BEGIN
  INSERT INTO asset_fts(rowid, uuid, name) VALUES (new.rowid, new.uuid, new.name);
END;
-- name: create-asset-delete-trigger
CREATE TRIGGER IF NOT EXISTS asset_delete AFTER DELETE ON asset BEGIN
  INSERT INTO asset_fts(asset_fts, rowid, uuid, name) VALUES('delete', old.rowid, old.uuid, old.name);
END;
-- name: create-asset-update-trigger
CREATE TRIGGER IF NOT EXISTS asset_update AFTER UPDATE ON asset BEGIN
  INSERT INTO asset_fts(asset_fts, rowid, uuid, name) VALUES('delete', old.rowid, old.uuid, old.name);
  INSERT INTO asset_fts(rowid, uuid, name) VALUES (new.rowid, new.uuid, new.name);
END;

-- name: create-tag-table
CREATE TABLE IF NOT EXISTS tag (
    text TEXT UNIQUE ON CONFLICT ROLLBACK NOT NULL 
);

-- name: create-tag-fts-table
CREATE VIRTUAL TABLE IF NOT EXISTS tag_fts USING fts5(text, tokenize=trigram, content=tag, content_rowid=rowid);

-- name: create-tag-insert-trigger
CREATE TRIGGER IF NOT EXISTS tag_insert AFTER INSERT ON tag BEGIN
  INSERT INTO tag_fts(rowid, text) VALUES (new.rowid, new.text);
END;
-- name: create-tag-delete-trigger
CREATE TRIGGER IF NOT EXISTS tag_delete AFTER DELETE ON tag BEGIN
  INSERT INTO tag_fts(tag_fts, rowid, text) VALUES('delete', old.rowid, old.text);
END;
-- name: create-tag-update-trigger
CREATE TRIGGER IF NOT EXISTS tag_update AFTER UPDATE ON tag BEGIN
  INSERT INTO tag_fts(tag_fts, rowid, text) VALUES('delete', old.rowid, old.text);
  INSERT INTO tag_fts(rowid, text) VALUES (new.rowid, new.text);
END;

-- name: create-asset-file-table
CREATE TABLE IF NOT EXISTS asset_file (
    uuid TEXT UNIQUE ON CONFLICT ROLLBACK NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    extension TEXT NOT NULL,
    last_update TEXT
);

-- name: create-tag-asset-assoc-table
CREATE TABLE IF NOT EXISTS tag_to_asset (
    asset_id TEXT,
    tag_id TEXT,
    PRIMARY KEY (asset_id, tag_id) ON CONFLICT ROLLBACK
);

-- name: create-asset-file-assoc-table
CREATE TABLE IF NOT EXISTS asset_to_asset_file (
    asset_id TEXT,
    asset_file_id TEXT,
    PRIMARY KEY (asset_id, asset_file_id) ON CONFLICT ROLLBACK
);
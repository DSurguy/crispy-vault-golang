package main

import (
	"context"
	"embed"

	"crispy-vault/src/assetmanager"
	"crispy-vault/src/dbmanager"
	"crispy-vault/src/tagmanager"
	"crispy-vault/src/vault"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	vault := vault.Vault{}
	vault.Bootstrap()
	db := dbmanager.DatabaseManager{}
	db.Provide(&vault)
	db.Bootstrap()
	assetManager := assetmanager.AssetManager{}
	assetManager.Provide(&vault, &db)
	tagManager := tagmanager.TagManager{}
	tagManager.Provide(&db)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Crispy Vault",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			db.SetContext(ctx)
			vault.SetContext(ctx)
			assetManager.SetContext(ctx)
		},
		Bind: []interface{}{
			app,
			&assetManager,
			&tagManager,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

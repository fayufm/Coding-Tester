package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed tubiao.ico
var icon []byte

func main() {
	// 创建应用程序实例
	app := NewApp()

	// 创建应用程序配置
	err := wails.Run(&options.App{
		Title:  "Networ Tester",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			IsZoomControlEnabled: false,
			ZoomFactor:           1.0,
			Theme:                windows.SystemDefault,
		},
		LogLevel:           logger.ERROR,
		LogLevelProduction: logger.ERROR,
	})

	if err != nil {
		log.Fatal(err)
	}
}

package main

//go:generate go tool goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest

import (
	"os"
	"path/filepath"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/files"
	"github.com/portapps/portapps/v3/pkg/log"
)

type config struct {
	Cleanup bool `yaml:"cleanup" mapstructure:"cleanup"`
}

var (
	app *portapps.App
	cfg *config
)

func init() {
	var err error

	// Default config
	cfg = &config{
		Cleanup: false,
	}

	// Init app
	if app, err = portapps.NewWithCfg("vscode-portable", "Visual Studio Code", cfg); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	if err := os.MkdirAll(app.DataPath, os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("Cannot create data directory.")
	}
	app.Process = filepath.Join(app.AppPath, "Code.exe")
	app.Args = []string{
		"--log debug",
	}

	// Cleanup on exit
	if cfg.Cleanup {
		defer func() {
			files.Cleanup(filepath.Join(os.Getenv("APPDATA"), "Code"))
		}()
	}

	os.Setenv("VSCODE_APPDATA", filepath.Join(app.DataPath, "appdata"))
	if !app.Config().Common.DisableLog {
		os.Setenv("VSCODE_LOGS", filepath.Join(app.DataPath, "logs"))
	}
	os.Setenv("VSCODE_EXTENSIONS", filepath.Join(app.DataPath, "extensions"))

	defer app.Close()
	app.Launch(os.Args[1:])
}

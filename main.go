//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"
	"path"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/utl"
)

type config struct {
	Cleanup bool `yaml:"cleanup" mapstructure:"cleanup"`
}

var (
	app *App
	cfg *config
)

func init() {
	var err error

	// Default config
	cfg = &config{
		Cleanup: false,
	}

	// Init app
	if app, err = NewWithCfg("vscode-portable", "Visual Studio Code", cfg); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	app.Process = utl.PathJoin(app.AppPath, "Code.exe")
	app.Args = []string{
		"--log debug",
	}

	// Cleanup on exit
	if cfg.Cleanup {
		defer func() {
			utl.Cleanup([]string{
				path.Join(os.Getenv("APPDATA"), "Code"),
			})
		}()
	}

	utl.OverrideEnv("VSCODE_APPDATA", utl.PathJoin(app.DataPath, "appdata"))
	utl.OverrideEnv("VSCODE_LOGS", utl.PathJoin(app.DataPath, "logs"))
	utl.OverrideEnv("VSCODE_EXTENSIONS", utl.PathJoin(app.DataPath, "extensions"))

	app.Launch(os.Args[1:])
}

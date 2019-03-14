//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/utl"
)

var (
	app *App
)

func init() {
	var err error

	// Init app
	if app, err = New("vscode-portable", "Visual Studio Code"); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	app.Process = utl.PathJoin(app.AppPath, "Code.exe")
	app.Args = []string{
		"--log debug",
	}

	utl.OverrideEnv("VSCODE_APPDATA", utl.PathJoin(app.DataPath, "appdata"))
	utl.OverrideEnv("VSCODE_LOGS", utl.PathJoin(app.DataPath, "logs"))
	utl.OverrideEnv("VSCODE_EXTENSIONS", utl.PathJoin(app.DataPath, "extensions"))

	app.Launch(os.Args[1:])
}

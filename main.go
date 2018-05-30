//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"

	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "vscode-portable"
	Papp.Name = "Visual Studio Code"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = CreateFolder(AppPathJoin("data"))
	Papp.Process = PathJoin(Papp.AppPath, "Code.exe")
	Papp.Args = []string{"--log debug"}
	Papp.WorkingDir = PathJoin(Papp.AppPath)

	OverrideEnv("VSCODE_APPDATA", PathJoin(Papp.DataPath, "appdata"))
	OverrideEnv("VSCODE_LOGS", PathJoin(Papp.DataPath, "logs"))
	OverrideEnv("VSCODE_EXTENSIONS", PathJoin(Papp.DataPath, "extensions"))

	Launch(os.Args[1:])
}

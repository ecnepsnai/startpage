package startpage

import (
	"path"

	"github.com/ecnepsnai/logtic"
)

var log = logtic.Connect("startpage")

func Start() {
	fsSetup()
	SetupLogtic()
	LoadOptions()
	ModuleSetup()
	ModuleRefresh()
	StartCron()
	StartRouter()
}

func SetupLogtic() {
	logtic.Log.FilePath = path.Join(Directories.Logs, "startpage.log")
	logtic.Open()
}

func Stop() {
	ModuleTeardown()
}

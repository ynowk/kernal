package portapps

import (
	"github.com/ynowk/kernal/pkg/utl"
)

// ElectronAppPath returns the app electron path
func (app *App) ElectronAppPath() string {
	electronAppFolder, err := utl.FindElectronAppFolder("app-", app.AppPath)
	if err != nil {
		app.FatalBoxLog(err.Error())
	}
	return utl.PathJoin(app.AppPath, electronAppFolder)
}

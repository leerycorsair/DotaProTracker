package delete_widgets

import (
	"local/internal/logger"
	"local/internal/managers"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

type DeleteTeamPlayerBox struct {
	mainWindow        fyne.Window
	teamPlayerManager managers.TeamPlayerManager
}

func (obj *DeleteTeamPlayerBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.teamPlayerManager = managers.TeamPlayerMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")

	handler := func() {
		err := obj.teamPlayerManager.Delete(idEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}

	btn := widget.NewButton("Удалить", handler)

	box := container.NewCenter(container.NewVBox(idEntry,
		btn))
	return box
}

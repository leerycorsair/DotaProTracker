package delete_widgets

import (
	"local/internal/logger"
	"local/internal/managers"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

type DeleteTournamentBox struct {
	mainWindow        fyne.Window
	tournamentManager managers.TournamentManager
}

func (obj *DeleteTournamentBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.tournamentManager = managers.TournamentMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")

	handler := func() {
		err := obj.tournamentManager.Delete(idEntry.Text)
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

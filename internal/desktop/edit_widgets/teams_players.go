package edit_widgets

import (
	"local/internal/logger"
	"local/internal/managers"
	"local/internal/models"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

type EditTeamPlayerBox struct {
	mainWindow        fyne.Window
	teamPlayerManager managers.TeamPlayerManager
}

func (obj *EditTeamPlayerBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.teamPlayerManager = managers.TeamPlayerMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")
	teamIdEntry := widget.NewEntry()
	teamIdEntry.SetPlaceHolder("ID Команды...")
	playerIdEntry := widget.NewEntry()
	playerIdEntry.SetPlaceHolder("ID Игрока...")
	contractDateEntry := widget.NewEntry()
	contractDateEntry.SetPlaceHolder("Дата контракта...")
	contractTimeEntry := widget.NewEntry()
	contractTimeEntry.SetPlaceHolder("Продолжительность...")

	handler := func() {
		info := models.TeamPlayerView{
			Id:           idEntry.Text,
			TeamId:       teamIdEntry.Text,
			PlayerId:     playerIdEntry.Text,
			ContractDate: contractDateEntry.Text,
			ContractTime: contractTimeEntry.Text,
		}
		err := obj.teamPlayerManager.Edit(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}

	btn := widget.NewButton("Изменить", handler)

	box := container.NewCenter(container.NewVBox(idEntry,
		teamIdEntry,
		playerIdEntry,
		contractDateEntry,
		contractTimeEntry,
		btn))
	return box
}

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

type EditMatchBox struct {
	mainWindow   fyne.Window
	matchManager managers.MatchManager
}

func (obj *EditMatchBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.matchManager = managers.MatchMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")
	tournamentIdEntry := widget.NewEntry()
	tournamentIdEntry.SetPlaceHolder("ID Турнира...")
	rTeamIdEntry := widget.NewEntry()
	rTeamIdEntry.SetPlaceHolder("ID Команды Света...")
	dTeamIdEntry := widget.NewEntry()
	dTeamIdEntry.SetPlaceHolder("ID Команды Тьмы...")
	durationEntry := widget.NewEntry()
	durationEntry.SetPlaceHolder("Продолжительность...")
	winnerEntry := widget.NewEntry()
	winnerEntry.SetPlaceHolder("Победитель...")

	handler := func() {
		info := models.MatchView{
			Id:           idEntry.Text,
			TournamentId: tournamentIdEntry.Text,
			RTeamId:      rTeamIdEntry.Text,
			DTeamId:      dTeamIdEntry.Text,
			Duration:     durationEntry.Text,
			Winner:       winnerEntry.Text,
		}
		err := obj.matchManager.Edit(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}

	btn := widget.NewButton("Изменить", handler)

	box := container.NewCenter(container.NewVBox(idEntry,
		tournamentIdEntry,
		rTeamIdEntry,
		dTeamIdEntry,
		durationEntry,
		winnerEntry,
		btn))
	return box
}

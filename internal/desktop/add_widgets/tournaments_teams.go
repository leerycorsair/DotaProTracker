package add_widgets

import (
	"local/internal/logger"
	"local/internal/managers"
	"local/internal/models"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

type AddTournamentTeamBox struct {
	mainWindow            fyne.Window
	tournamentTeamManager managers.TournamentTeamManager
}

func (obj *AddTournamentTeamBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.tournamentTeamManager = managers.TournamentTeamMan
	tournamentEntry := widget.NewEntry()
	tournamentEntry.SetPlaceHolder("ID Турнира...")
	teamIdEntry := widget.NewEntry()
	teamIdEntry.SetPlaceHolder("ID Команды...")
	participationTypeEntry := widget.NewEntry()
	participationTypeEntry.SetPlaceHolder("Способ участия...")
	isWinnerEntry := widget.NewEntry()
	isWinnerEntry.SetPlaceHolder("Является ли победителем...")

	handler := func() {
		info := models.TournamentTeamView{
			Id:                "1",
			TournamentId:      tournamentEntry.Text,
			TeamId:            teamIdEntry.Text,
			ParticipationType: participationTypeEntry.Text,
			IsWinner:          isWinnerEntry.Text,
		}
		err := obj.tournamentTeamManager.Add(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}
	btn := widget.NewButton("Добавить", handler)

	box := container.NewCenter(container.NewVBox(tournamentEntry,
		teamIdEntry,
		participationTypeEntry,
		isWinnerEntry,
		btn))
	return box
}

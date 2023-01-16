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

type EditMatchPerfomanceBox struct {
	mainWindow             fyne.Window
	matchPerfomanceManager managers.MatchPerfomanceManager
}

func (obj *EditMatchPerfomanceBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.matchPerfomanceManager = managers.MatchPerfomanceMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")
	matchIdEntry := widget.NewEntry()
	matchIdEntry.SetPlaceHolder("ID Матча...")
	playerIdEntry := widget.NewEntry()
	playerIdEntry.SetPlaceHolder("ID Игрока...")
	teamEntry := widget.NewEntry()
	teamEntry.SetPlaceHolder("Команда...")
	heroEntry := widget.NewEntry()
	heroEntry.SetPlaceHolder("Герой...")
	killsEntry := widget.NewEntry()
	killsEntry.SetPlaceHolder("Убийства...")
	deathsEntry := widget.NewEntry()
	deathsEntry.SetPlaceHolder("Смерти...")
	assistsEntry := widget.NewEntry()
	assistsEntry.SetPlaceHolder("Помощи...")
	networthEntry := widget.NewEntry()
	networthEntry.SetPlaceHolder("Ценность...")
	gpmEntry := widget.NewEntry()
	gpmEntry.SetPlaceHolder("ЗВМ...")
	xpmEntry := widget.NewEntry()
	xpmEntry.SetPlaceHolder("ОПМ...")
	dmgEntry := widget.NewEntry()
	dmgEntry.SetPlaceHolder("Урон...")
	healEntry := widget.NewEntry()
	healEntry.SetPlaceHolder("Лечение...")
	bldEntry := widget.NewEntry()
	bldEntry.SetPlaceHolder("Урон по постройкам...")

	handler := func() {
		info := models.MatchPerfomanceView{
			Id:       idEntry.Text,
			MatchId:  matchIdEntry.Text,
			PlayerId: playerIdEntry.Text,
			Team:     teamEntry.Text,
			Hero:     heroEntry.Text,
			Kills:    killsEntry.Text,
			Deaths:   deathsEntry.Text,
			Assists:  assistsEntry.Text,
			Networth: networthEntry.Text,
			GPM:      gpmEntry.Text,
			XPM:      xpmEntry.Text,
			DMG:      dmgEntry.Text,
			Heal:     healEntry.Text,
			BLD:      bldEntry.Text,
		}
		err := obj.matchPerfomanceManager.Edit(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}

	btn := widget.NewButton("Изменить", handler)

	box := container.NewCenter(container.NewVBox(idEntry,
		matchIdEntry,
		playerIdEntry,
		teamEntry,
		heroEntry,
		killsEntry,
		deathsEntry,
		assistsEntry,
		networthEntry,
		gpmEntry,
		xpmEntry,
		dmgEntry,
		healEntry,
		bldEntry,
		btn))
	return box
}

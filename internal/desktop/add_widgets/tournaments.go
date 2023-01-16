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

type AddTournamentBox struct {
	mainWindow        fyne.Window
	tournamentManager managers.TournamentManager
}

func (obj *AddTournamentBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.tournamentManager = managers.TournamentMan
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Название...")
	tierEntry := widget.NewEntry()
	tierEntry.SetPlaceHolder("Уровень...")
	prizePoolEntry := widget.NewEntry()
	prizePoolEntry.SetPlaceHolder("Призовой фонд...")
	dateStartEntry := widget.NewEntry()
	dateStartEntry.SetPlaceHolder("Дата начала...")
	durationEntry := widget.NewEntry()
	durationEntry.SetPlaceHolder("Продолжительность...")
	dpcPointsEntry := widget.NewEntry()
	dpcPointsEntry.SetPlaceHolder("DPC очки...")
	locationEntry := widget.NewEntry()
	locationEntry.SetPlaceHolder("Место проведения...")

	handler := func() {
		info := models.TournamentView{
			Id:        "1",
			Name:      nameEntry.Text,
			Tier:      tierEntry.Text,
			PrizePool: prizePoolEntry.Text,
			DateStart: dateStartEntry.Text,
			Duration:  durationEntry.Text,
			DPCPoints: dpcPointsEntry.Text,
			Location:  locationEntry.Text,
		}
		err := obj.tournamentManager.Add(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}
	btn := widget.NewButton("Добавить", handler)

	box := container.NewCenter(container.NewVBox(nameEntry,
		tierEntry,
		prizePoolEntry,
		dateStartEntry,
		durationEntry,
		dpcPointsEntry,
		locationEntry,
		btn))
	return box
}

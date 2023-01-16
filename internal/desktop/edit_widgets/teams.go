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

type EditTeamBox struct {
	mainWindow  fyne.Window
	teamManager managers.TeamManager
}

func (obj *EditTeamBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.teamManager = managers.TeamMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Название...")
	createdAtEntry := widget.NewEntry()
	createdAtEntry.SetPlaceHolder("Дата создания...")
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Почта...")
	totalEarningsEntry := widget.NewEntry()
	totalEarningsEntry.SetPlaceHolder("Общий заработок...")
	regionEntry := widget.NewEntry()
	regionEntry.SetPlaceHolder("Регион...")
	tierEntry := widget.NewEntry()
	tierEntry.SetPlaceHolder("Уровень...")

	handler := func() {
		info := models.TeamView{
			Id:            idEntry.Text,
			Name:          nameEntry.Text,
			CreatedAt:     createdAtEntry.Text,
			Email:         emailEntry.Text,
			TotalEarnings: totalEarningsEntry.Text,
			Region:        regionEntry.Text,
			Tier:          tierEntry.Text,
		}
		err := obj.teamManager.Edit(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}

	btn := widget.NewButton("Изменить", handler)

	box := container.NewCenter(container.NewVBox(idEntry,
		nameEntry,
		createdAtEntry,
		emailEntry,
		totalEarningsEntry,
		regionEntry,
		tierEntry,
		btn))
	return box
}

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

type AddTeamBox struct {
	mainWindow  fyne.Window
	teamManager managers.TeamManager
}

func (obj *AddTeamBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.teamManager = managers.TeamMan
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
			Id:            "1",
			Name:          nameEntry.Text,
			CreatedAt:     createdAtEntry.Text,
			Email:         emailEntry.Text,
			TotalEarnings: totalEarningsEntry.Text,
			Region:        regionEntry.Text,
			Tier:          tierEntry.Text,
		}
		err := obj.teamManager.Add(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}

	btn := widget.NewButton("Добавить", handler)

	box := container.NewCenter(container.NewVBox(nameEntry,
		createdAtEntry,
		emailEntry,
		totalEarningsEntry,
		regionEntry,
		tierEntry,
		btn))
	return box
}

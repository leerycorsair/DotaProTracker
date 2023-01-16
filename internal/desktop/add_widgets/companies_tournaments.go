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

type AddCompanyTournamentBox struct {
	mainWindow               fyne.Window
	companyTournamentManager managers.CompanyTournamentManager
}

func (obj *AddCompanyTournamentBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.companyTournamentManager = managers.CompanyTournamentMan
	companyIdEntry := widget.NewEntry()
	companyIdEntry.SetPlaceHolder("ID Компании...")
	tournamentIdEntry := widget.NewEntry()
	tournamentIdEntry.SetPlaceHolder("ID Турнира...")
	depositEntry := widget.NewEntry()
	depositEntry.SetPlaceHolder("Депозит...")

	handler := func() {
		info := models.CompanyTournamentView{
			Id:           "1",
			CompanyId:    companyIdEntry.Text,
			TournamentId: tournamentIdEntry.Text,
			Deposit:      depositEntry.Text,
		}
		err := obj.companyTournamentManager.Add(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}
	btn := widget.NewButton("Добавить", handler)

	box := container.NewCenter(container.NewVBox(companyIdEntry,
		tournamentIdEntry,
		depositEntry,
		btn))
	return box
}

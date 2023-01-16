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

type EditCompanyTournamentBox struct {
	mainWindow               fyne.Window
	companyTournamentManager managers.CompanyTournamentManager
}

func (obj *EditCompanyTournamentBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.companyTournamentManager = managers.CompanyTournamentMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")
	companyIdEntry := widget.NewEntry()
	companyIdEntry.SetPlaceHolder("ID Компании...")
	tournamentIdEntry := widget.NewEntry()
	tournamentIdEntry.SetPlaceHolder("ID Турнира...")
	depositEntry := widget.NewEntry()
	depositEntry.SetPlaceHolder("Депозит...")

	handler := func() {
		info := models.CompanyTournamentView{
			Id:           idEntry.Text,
			CompanyId:    companyIdEntry.Text,
			TournamentId: tournamentIdEntry.Text,
			Deposit:      depositEntry.Text,
		}
		err := obj.companyTournamentManager.Edit(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}
	btn := widget.NewButton("Изменить", handler)

	box := container.NewCenter(container.NewVBox(idEntry,
		companyIdEntry,
		tournamentIdEntry,
		depositEntry,
		btn))
	return box
}

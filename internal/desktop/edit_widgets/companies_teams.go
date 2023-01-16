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

type EditCompanyTeamBox struct {
	mainWindow         fyne.Window
	companyTeamManager managers.CompanyTeamManager
}

func (obj *EditCompanyTeamBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.companyTeamManager = managers.СompanyTeamMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")
	companyIdEntry := widget.NewEntry()
	companyIdEntry.SetPlaceHolder("ID Компании...")
	teamIdEntry := widget.NewEntry()
	teamIdEntry.SetPlaceHolder("ID Команды...")
	contractDateEntry := widget.NewEntry()
	contractDateEntry.SetPlaceHolder("Дата контракта...")
	contractTimeEntry := widget.NewEntry()
	contractTimeEntry.SetPlaceHolder("Продолжительность...")

	handler := func() {
		info := models.CompanyTeamView{
			Id:           idEntry.Text,
			CompanyId:    companyIdEntry.Text,
			TeamId:       teamIdEntry.Text,
			ContractDate: contractDateEntry.Text,
			ContractTime: contractTimeEntry.Text,
		}
		err := obj.companyTeamManager.Edit(info)
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
		teamIdEntry,
		contractDateEntry,
		contractTimeEntry,
		btn))
	return box
}

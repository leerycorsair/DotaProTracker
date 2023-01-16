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

type EditCompanyBox struct {
	mainWindow     fyne.Window
	companyManager managers.CompanyManager
}

func (obj *EditCompanyBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.companyManager = managers.CompanyMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Название...")
	countryEntry := widget.NewEntry()
	countryEntry.SetPlaceHolder("Страна...")
	websiteEntry := widget.NewEntry()
	websiteEntry.SetPlaceHolder("Сайт...")
	revenueEntry := widget.NewEntry()
	revenueEntry.SetPlaceHolder("Оборот...")
	industryEntry := widget.NewEntry()
	industryEntry.SetPlaceHolder("Индустрия..")

	handler := func() {
		info := models.CompanyView{
			Id:       idEntry.Text,
			Name:     nameEntry.Text,
			Country:  countryEntry.Text,
			Website:  websiteEntry.Text,
			Revenue:  revenueEntry.Text,
			Industry: industryEntry.Text,
		}
		err := obj.companyManager.Edit(info)
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
		countryEntry,
		websiteEntry,
		revenueEntry,
		industryEntry,
		btn))
	return box
}

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

type AddCompanyBox struct {
	mainWindow     fyne.Window
	companyManager managers.CompanyManager
}

func (obj *AddCompanyBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.companyManager = managers.CompanyMan
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Название...")
	countryEntry := widget.NewEntry()
	countryEntry.SetPlaceHolder("Страна...")
	websiteEntry := widget.NewEntry()
	websiteEntry.SetPlaceHolder("Сайт...")
	revenueEntry := widget.NewEntry()
	revenueEntry.SetPlaceHolder("Оборот...")
	industryEntry := widget.NewEntry()
	industryEntry.SetPlaceHolder("Индустрия...")

	handler := func() {
		info := models.CompanyView{
			Id:       "1",
			Name:     nameEntry.Text,
			Country:  countryEntry.Text,
			Website:  websiteEntry.Text,
			Revenue:  revenueEntry.Text,
			Industry: industryEntry.Text,
		}
		err := obj.companyManager.Add(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}

	btn := widget.NewButton("Добавить", handler)

	box := container.NewCenter(container.NewVBox(nameEntry,
		countryEntry,
		websiteEntry,
		revenueEntry,
		industryEntry,
		btn))
	return box
}

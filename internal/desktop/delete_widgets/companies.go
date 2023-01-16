package delete_widgets

import (
	"local/internal/logger"
	"local/internal/managers"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

type DeleteCompanyBox struct {
	mainWindow     fyne.Window
	companyManager managers.CompanyManager
}

func (obj *DeleteCompanyBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.companyManager = managers.CompanyMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")

	handler := func() {
		err := obj.companyManager.Delete(idEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}

	btn := widget.NewButton("Удалить", handler)

	box := container.NewCenter(container.NewVBox(idEntry,
		btn))
	return box
}

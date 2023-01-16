package user_widgets

import (
	"local/internal/logger"
	"local/internal/managers"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

type DeleteUserBox struct {
	mainWindow  fyne.Window
	userManager managers.UserManager
}

func (obj *DeleteUserBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.userManager = managers.UserMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")

	handler := func() {
		err := obj.userManager.Delete(idEntry.Text)
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

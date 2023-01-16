package user_widgets

import (
	"local/internal/logger"
	"local/internal/managers"
	"local/internal/models"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

type EditUserBox struct {
	mainWindow  fyne.Window
	userManager managers.UserManager
}

func (obj *EditUserBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.userManager = managers.UserMan
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID...")
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Имя...")
	birthdateEntry := widget.NewEntry()
	birthdateEntry.SetPlaceHolder("Дата рождения...")
	loginEntry := widget.NewEntry()
	loginEntry.SetPlaceHolder("Логин...")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Пароль...")
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Почта...")
	privilegeLevelEntry := widget.NewEntry()
	privilegeLevelEntry.SetPlaceHolder("Уровень привилегий...")

	handler := func() {
		info := models.UserView{
			Id:             idEntry.Text,
			Name:           nameEntry.Text,
			Birthdate:      birthdateEntry.Text,
			Login:          loginEntry.Text,
			Password:       passwordEntry.Text,
			Email:          emailEntry.Text,
			PrivilegeLevel: privilegeLevelEntry.Text,
		}
		err := obj.userManager.Edit(info)
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
		birthdateEntry,
		loginEntry,
		passwordEntry,
		emailEntry,
		privilegeLevelEntry,
		btn))
	return box
}

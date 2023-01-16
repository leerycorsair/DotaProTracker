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

type AddPlayerBox struct {
	mainWindow    fyne.Window
	playerManager managers.PlayerManager
}

func (obj *AddPlayerBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.playerManager = managers.PlayerMan
	nicknameEntry := widget.NewEntry()
	nicknameEntry.SetPlaceHolder("Прозвище...")
	realNameEntry := widget.NewEntry()
	realNameEntry.SetPlaceHolder("Имя...")
	birthdateEntry := widget.NewEntry()
	birthdateEntry.SetPlaceHolder("Дата рождения...")
	countryEntry := widget.NewEntry()
	countryEntry.SetPlaceHolder("Страна...")
	mmrEntry := widget.NewEntry()
	mmrEntry.SetPlaceHolder("Рейтинг...")
	roleEntry := widget.NewEntry()
	roleEntry.SetPlaceHolder("Роль...")
	sigHeroEntry := widget.NewEntry()
	sigHeroEntry.SetPlaceHolder("Избранный герой...")

	handler := func() {
		info := models.PlayerView{
			Id:            "1",
			Nickname:      nicknameEntry.Text,
			Realname:      realNameEntry.Text,
			Birthdate:     birthdateEntry.Text,
			Country:       countryEntry.Text,
			MMR:           mmrEntry.Text,
			Role:          roleEntry.Text,
			SignatureHero: sigHeroEntry.Text,
		}
		err := obj.playerManager.Add(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			dialog.NewInformation("Информация", "Успешно", obj.mainWindow)
		}
	}
	btn := widget.NewButton("Добавить", handler)

	box := container.NewCenter(container.NewVBox(nicknameEntry,
		realNameEntry,
		birthdateEntry,
		countryEntry,
		mmrEntry,
		roleEntry,
		sigHeroEntry,
		btn))
	return box
}

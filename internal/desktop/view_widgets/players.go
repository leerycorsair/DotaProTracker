package view_widgets

import (
	"fmt"
	"local/internal/logger"
	"local/internal/managers"
	"local/internal/models"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

type ViewPlayerBox struct {
	mainWindow    fyne.Window
	PlayerManager managers.PlayerManager
}

func PlayersToData(recs []*models.PlayerView) [][]string {
	data := [][]string{}
	data = append(data, []string{"Id", "Nickname", "Realname", "Birthdate", "Country", "MMR", "Role", "SignatureHero"})
	for i := 0; i < len(recs); i++ {
		data = append(data, []string{recs[i].Id, recs[i].Nickname, recs[i].Realname, recs[i].Birthdate, recs[i].Country, recs[i].MMR, recs[i].Role, recs[i].SignatureHero})
	}
	return data
}

func (obj *ViewPlayerBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.PlayerManager = managers.PlayerMan

	handler1 := func() {
		recs, err := obj.PlayerManager.GetAll()
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := PlayersToData(recs)
			l := widget.NewTable(
				func() (int, int) {
					return len(data), len(data[0])
				},
				func() fyne.CanvasObject {
					item := widget.NewLabel("LongLongLongLongLongLongLong")
					return item
				},
				func(cell widget.TableCellID, item fyne.CanvasObject) {
					item.(*widget.Label).SetText(fmt.Sprintf("%s", data[cell.Row][cell.Col]))
				})
			d := dialog.NewCustom("Таблица", "OK", l, obj.mainWindow)
			d.Resize(fyne.NewSize(1800, 1000))
			d.Show()
		}
	}

	btn1 := widget.NewButton("Показать все", handler1)

	playerIdEntry := widget.NewEntry()
	playerIdEntry.SetPlaceHolder("ID...")
	handler2 := func() {
		recs, err := obj.PlayerManager.FindById(playerIdEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := PlayersToData(recs)
			l := widget.NewTable(
				func() (int, int) {
					return len(data), len(data[0])
				},
				func() fyne.CanvasObject {
					item := widget.NewLabel("LongLongLongLongLongLongLong")
					return item
				},
				func(cell widget.TableCellID, item fyne.CanvasObject) {
					item.(*widget.Label).SetText(fmt.Sprintf("%s", data[cell.Row][cell.Col]))
				})
			d := dialog.NewCustom("Таблица", "OK", l, obj.mainWindow)
			d.Resize(fyne.NewSize(1500, 1000))
			d.Show()
		}
	}
	btn2 := widget.NewButton("Поиск", handler2)

	PlayerNameEntry := widget.NewEntry()
	PlayerNameEntry.SetPlaceHolder("Имя...")

	handler3 := func() {
		recs, err := obj.PlayerManager.FindByName(PlayerNameEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := PlayersToData(recs)
			l := widget.NewTable(
				func() (int, int) {
					return len(data), len(data[0])
				},
				func() fyne.CanvasObject {
					item := widget.NewLabel("LongLongLongLongLongLongLong")
					return item
				},
				func(cell widget.TableCellID, item fyne.CanvasObject) {
					item.(*widget.Label).SetText(fmt.Sprintf("%s", data[cell.Row][cell.Col]))
				})
			d := dialog.NewCustom("Таблица", "OK", l, obj.mainWindow)
			d.Resize(fyne.NewSize(1800, 1000))
			d.Show()
		}
	}
	btn3 := widget.NewButton("Поиск", handler3)

	PlayerBirthdateEntry := widget.NewEntry()
	PlayerBirthdateEntry.SetPlaceHolder("Год рождения...")

	handler4 := func() {
		recs, err := obj.PlayerManager.FindByBirthdate(PlayerBirthdateEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := PlayersToData(recs)
			l := widget.NewTable(
				func() (int, int) {
					return len(data), len(data[0])
				},
				func() fyne.CanvasObject {
					item := widget.NewLabel("LongLongLongLongLongLongLong")
					return item
				},
				func(cell widget.TableCellID, item fyne.CanvasObject) {
					item.(*widget.Label).SetText(fmt.Sprintf("%s", data[cell.Row][cell.Col]))
				})
			d := dialog.NewCustom("Таблица", "OK", l, obj.mainWindow)
			d.Resize(fyne.NewSize(1800, 1000))
			d.Show()
		}
	}
	btn4 := widget.NewButton("Поиск", handler4)

	box := container.NewCenter(container.NewVBox(btn1,
		playerIdEntry,
		btn2,
		PlayerNameEntry,
		btn3,
		PlayerBirthdateEntry,
		btn4))
	return box
}

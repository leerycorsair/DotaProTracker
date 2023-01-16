package user_widgets

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

type ViewUserBox struct {
	mainWindow  fyne.Window
	UserManager managers.UserManager
}

func UsersToData(recs []*models.UserView) [][]string {
	data := [][]string{}
	data = append(data, []string{"Id", "Name", "Birthdate", "Login", "Email", "PrivilegeLevel"})
	for i := 0; i < len(recs); i++ {
		data = append(data, []string{recs[i].Id, recs[i].Name, recs[i].Birthdate, recs[i].Login, recs[i].Email, recs[i].PrivilegeLevel})
	}
	return data
}

func (obj *ViewUserBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.UserManager = managers.UserMan

	handler1 := func() {
		recs, err := obj.UserManager.GetAll()
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := UsersToData(recs)
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

	UserIdEntry := widget.NewEntry()
	UserIdEntry.SetPlaceHolder("ID...")
	handler2 := func() {
		recs, err := obj.UserManager.FindById(UserIdEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := UsersToData(recs)
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
	btn2 := widget.NewButton("Поиск", handler2)

	box := container.NewCenter(container.NewVBox(btn1,
		UserIdEntry,
		btn2))
	return box
}

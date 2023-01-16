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

type ViewTeamBox struct {
	mainWindow  fyne.Window
	TeamManager managers.TeamManager
}

func TeamsToData(recs []*models.TeamView) [][]string {
	data := [][]string{}
	data = append(data, []string{"Id", "Name", "CreatedAt", "Email", "TotalEarnings", "Region", "Tier"})
	for i := 0; i < len(recs); i++ {
		data = append(data, []string{recs[i].Id, recs[i].Name, recs[i].CreatedAt, recs[i].Email, recs[i].TotalEarnings, recs[i].Region, recs[i].Tier})
	}
	return data
}

func (obj *ViewTeamBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.TeamManager = managers.TeamMan

	handler1 := func() {
		recs, err := obj.TeamManager.GetAll()
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := TeamsToData(recs)
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

	btn1 := widget.NewButton("Показать все", handler1)

	TeamIdEntry := widget.NewEntry()
	TeamIdEntry.SetPlaceHolder("ID...")
	handler2 := func() {
		recs, err := obj.TeamManager.FindById(TeamIdEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := TeamsToData(recs)
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

	TeamNameEntry := widget.NewEntry()
	TeamNameEntry.SetPlaceHolder("Название...")

	handler3 := func() {
		recs, err := obj.TeamManager.FindByName(TeamNameEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := TeamsToData(recs)
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

	box := container.NewCenter(container.NewVBox(btn1,
		TeamIdEntry,
		btn2,
		TeamNameEntry,
		btn3))
	return box
}

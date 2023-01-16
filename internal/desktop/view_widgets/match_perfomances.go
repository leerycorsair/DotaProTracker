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

type ViewMatchPerfomanceBox struct {
	mainWindow             fyne.Window
	MatchPerfomanceManager managers.MatchPerfomanceManager
}

func MatchPerfomancesToData(recs []*models.MatchPerfomanceView) [][]string {
	data := [][]string{}
	data = append(data, []string{"Id", "MatchId", "PlayerId", "Team", "Hero", "Kills", "Deaths", "Assists", "Networth", "GPM", "XPM", "DMG", "Heal", "BLD"})
	for i := 0; i < len(recs); i++ {
		data = append(data, []string{recs[i].Id, recs[i].MatchId, recs[i].PlayerId, recs[i].Team, recs[i].Hero, recs[i].Kills, recs[i].Deaths, recs[i].Assists, recs[i].Networth, recs[i].GPM, recs[i].XPM, recs[i].DMG, recs[i].Heal, recs[i].BLD})
	}
	return data
}

func (obj *ViewMatchPerfomanceBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.MatchPerfomanceManager = managers.MatchPerfomanceMan

	handler1 := func() {
		recs, err := obj.MatchPerfomanceManager.GetAll()
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := MatchPerfomancesToData(recs)
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

	MatchPerfomanceIdEntry := widget.NewEntry()
	MatchPerfomanceIdEntry.SetPlaceHolder("ID...")
	handler2 := func() {
		recs, err := obj.MatchPerfomanceManager.FindById(MatchPerfomanceIdEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := MatchPerfomancesToData(recs)
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

	HeroEntry := widget.NewEntry()
	HeroEntry.SetPlaceHolder("Герой...")

	handler3 := func() {
		recs, err := obj.MatchPerfomanceManager.FindByHero(HeroEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := MatchPerfomancesToData(recs)
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
		MatchPerfomanceIdEntry,
		btn2,
		HeroEntry,
		btn3))
	return box
}

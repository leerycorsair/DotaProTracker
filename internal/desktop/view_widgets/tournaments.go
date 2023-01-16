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

type ViewTournamentBox struct {
	mainWindow        fyne.Window
	TournamentManager managers.TournamentManager
}

func TournamentsToData(recs []*models.TournamentView) [][]string {
	data := [][]string{}
	data = append(data, []string{"Id", "Name", "Tier", "PrizePool", "DateStart", "Duration", "DPCPoints", "Location"})
	for i := 0; i < len(recs); i++ {
		data = append(data, []string{recs[i].Id, recs[i].Name, recs[i].Tier, recs[i].PrizePool, recs[i].DateStart, recs[i].Duration, recs[i].DPCPoints, recs[i].Location})
	}
	return data
}

func (obj *ViewTournamentBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.TournamentManager = managers.TournamentMan

	handler1 := func() {
		recs, err := obj.TournamentManager.GetAll()
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := TournamentsToData(recs)
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

	TournamentIdEntry := widget.NewEntry()
	TournamentIdEntry.SetPlaceHolder("ID...")
	handler2 := func() {
		recs, err := obj.TournamentManager.FindById(TournamentIdEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := TournamentsToData(recs)
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

	TournamentNameEntry := widget.NewEntry()
	TournamentNameEntry.SetPlaceHolder("Название...")

	handler3 := func() {
		recs, err := obj.TournamentManager.FindByName(TournamentNameEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := TournamentsToData(recs)
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

	TournamentYearEntry := widget.NewEntry()
	TournamentYearEntry.SetPlaceHolder("Год...")

	handler4 := func() {
		recs, err := obj.TournamentManager.FindByYear(TournamentYearEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := TournamentsToData(recs)
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
		TournamentIdEntry,
		btn2,
		TournamentNameEntry,
		btn3,
		TournamentYearEntry,
		btn4))

	return box
}

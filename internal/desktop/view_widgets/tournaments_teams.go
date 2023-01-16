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

type ViewTournamentTeamBox struct {
	mainWindow            fyne.Window
	TournamentTeamManager managers.TournamentTeamManager
}

func TournamentTeamToData(recs []*models.TournamentTeamView) [][]string {
	data := [][]string{}
	data = append(data, []string{"Id", "TournamentId", "TeamId", "ParticipationType", "IsWinner"})
	for i := 0; i < len(recs); i++ {
		data = append(data, []string{recs[i].Id, recs[i].TournamentId, recs[i].TeamId, recs[i].ParticipationType, recs[i].IsWinner})
	}
	return data
}

func (obj *ViewTournamentTeamBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.TournamentTeamManager = managers.TournamentTeamMan

	handler1 := func() {
		recs, err := obj.TournamentTeamManager.GetAll()
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := TournamentTeamToData(recs)
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

	TournamentTeamIdEntry := widget.NewEntry()
	TournamentTeamIdEntry.SetPlaceHolder("ID...")
	handler2 := func() {
		recs, err := obj.TournamentTeamManager.FindById(TournamentTeamIdEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := TournamentTeamToData(recs)
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
		TournamentTeamIdEntry,
		btn2))
	return box
}

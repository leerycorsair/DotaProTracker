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

type ViewTeamPlayerBox struct {
	mainWindow        fyne.Window
	TeamPlayerManager managers.TeamPlayerManager
}

func TeamPlayersToData(recs []*models.TeamPlayerView) [][]string {
	data := [][]string{}
	data = append(data, []string{"Id", "TeamId", "PlayerId", "ContractDate", "ContractTime"})
	for i := 0; i < len(recs); i++ {
		data = append(data, []string{recs[i].Id, recs[i].TeamId, recs[i].PlayerId, recs[i].ContractDate, recs[i].ContractTime})
	}
	return data
}

func (obj *ViewTeamPlayerBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.TeamPlayerManager = managers.TeamPlayerMan

	handler1 := func() {
		recs, err := obj.TeamPlayerManager.GetAll()
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := TeamPlayersToData(recs)
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

	TeamPlayerIdEntry := widget.NewEntry()
	TeamPlayerIdEntry.SetPlaceHolder("ID...")
	handler2 := func() {
		recs, err := obj.TeamPlayerManager.FindById(TeamPlayerIdEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := TeamPlayersToData(recs)
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

	box := container.NewCenter(container.NewVBox(btn1,
		TeamPlayerIdEntry,
		btn2))
	return box
}

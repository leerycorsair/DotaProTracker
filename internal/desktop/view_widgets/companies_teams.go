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

type ViewCompanyTeamBox struct {
	mainWindow         fyne.Window
	CompanyTeamManager managers.CompanyTeamManager
}

func CompanyTeamsToData(recs []*models.CompanyTeamView) [][]string {
	data := [][]string{}
	data = append(data, []string{"Id", "CompanyId", "TeamId", "ContractDate", "ContractTime"})
	for i := 0; i < len(recs); i++ {
		data = append(data, []string{recs[i].Id, recs[i].CompanyId, recs[i].TeamId, recs[i].ContractDate, recs[i].ContractTime})
	}
	return data
}

func (obj *ViewCompanyTeamBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.CompanyTeamManager = managers.СompanyTeamMan

	handler1 := func() {
		recs, err := obj.CompanyTeamManager.GetAll()
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := CompanyTeamsToData(recs)
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

	CompanyTeamIdEntry := widget.NewEntry()
	CompanyTeamIdEntry.SetPlaceHolder("ID...")
	handler2 := func() {
		recs, err := obj.CompanyTeamManager.FindById(CompanyTeamIdEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := CompanyTeamsToData(recs)
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
		CompanyTeamIdEntry,
		btn2))
	return box
}

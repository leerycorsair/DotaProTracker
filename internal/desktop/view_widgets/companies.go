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

type ViewCompanyBox struct {
	mainWindow     fyne.Window
	CompanyManager managers.CompanyManager
}

func CompanysToData(recs []*models.CompanyView) [][]string {
	data := [][]string{}
	data = append(data, []string{"Id", "Name", "Country", "Website", "Revenue", "Industry"})
	for i := 0; i < len(recs); i++ {
		data = append(data, []string{recs[i].Id, recs[i].Name, recs[i].Country, recs[i].Website, recs[i].Revenue, recs[i].Industry})
	}
	return data
}

func (obj *ViewCompanyBox) New(mainWindow fyne.Window, managers managers.Managers) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.CompanyManager = managers.CompanyMan

	handler1 := func() {
		recs, err := obj.CompanyManager.GetAll()
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := CompanysToData(recs)
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

	CompanyIdEntry := widget.NewEntry()
	CompanyIdEntry.SetPlaceHolder("ID...")
	handler2 := func() {
		recs, err := obj.CompanyManager.FindById(CompanyIdEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := CompanysToData(recs)
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

	CompanyNameEntry := widget.NewEntry()
	CompanyNameEntry.SetPlaceHolder("Название...")

	handler3 := func() {
		recs, err := obj.CompanyManager.FindByName(CompanyNameEntry.Text)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			data := CompanysToData(recs)
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
		CompanyIdEntry,
		btn2,
		CompanyNameEntry,
		btn3))
	return box
}

package login_widgets

import (
	"local/internal/desktop"
	"local/internal/managers"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type loginApp struct {
	app        fyne.App
	mainWindow fyne.Window
	managers   managers.Managers
	dApp       desktop.DesktopApp
}

func NewLoginWindow(app fyne.App, appTitle string, managers managers.Managers, dApp desktop.DesktopApp) loginApp {
	mainWindow := loginApp{app, app.NewWindow(appTitle), managers, dApp}
	return mainWindow
}

func (lApp loginApp) Init() {
	lApp.mainWindow.Resize(fyne.NewSize(800, 800))

	h_box := container.NewHBox((&RegUserBox{}).New(lApp.mainWindow, lApp.managers),
		widget.NewLabel("              "),
		(&LoginUserBox{}).New(lApp.app, lApp.mainWindow, lApp.managers, lApp.dApp))

	box := container.NewCenter(h_box)
	lApp.mainWindow.SetContent(box)
}

func (lApp loginApp) ShowAndRun() {
	lApp.mainWindow.ShowAndRun()
}

func (lApp loginApp) Show() {
	lApp.mainWindow.Show()
}

func (lApp loginApp) Run() {
	lApp.mainWindow.Show()
}

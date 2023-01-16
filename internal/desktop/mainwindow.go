package desktop

import (
	"local/internal/desktop/add_widgets"
	"local/internal/desktop/delete_widgets"
	"local/internal/desktop/edit_widgets"
	"local/internal/desktop/user_widgets"
	"local/internal/desktop/view_widgets"
	"local/internal/managers"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
)

type DesktopApp struct {
	mainWindow fyne.Window
	managers   managers.Managers
}

func NewMainWindow(app fyne.App, appTitle string, managers managers.Managers) DesktopApp {
	mainWindow := DesktopApp{app.NewWindow(appTitle), managers}
	return mainWindow
}

func (dApp DesktopApp) Init(privilege_lvl int) {
	dApp.mainWindow.Resize(fyne.NewSize(800, 800))
	if privilege_lvl == 1 {
		tabs := container.NewAppTabs(
			container.NewTabItem("Просмотр", view_widgets.NewViewTab(dApp.mainWindow, dApp.managers)),
		)
		tabs.SetTabLocation(container.TabLocationLeading)
		dApp.mainWindow.SetContent(tabs)
	} else if privilege_lvl == 2 {
		tabs := container.NewAppTabs(
			container.NewTabItem("Просмотр", view_widgets.NewViewTab(dApp.mainWindow, dApp.managers)),
			container.NewTabItem("Добавить", add_widgets.NewAddTab(dApp.mainWindow, dApp.managers)),
			container.NewTabItem("Редактировать", edit_widgets.NewEditTab(dApp.mainWindow, dApp.managers)),
			container.NewTabItem("Удалить", delete_widgets.NewDeleteTab(dApp.mainWindow, dApp.managers)),
		)
		tabs.SetTabLocation(container.TabLocationLeading)
		dApp.mainWindow.SetContent(tabs)
	} else {
		tabs := container.NewAppTabs(
			container.NewTabItem("Просмотр", view_widgets.NewViewTab(dApp.mainWindow, dApp.managers)),
			container.NewTabItem("Добавить", add_widgets.NewAddTab(dApp.mainWindow, dApp.managers)),
			container.NewTabItem("Редактировать", edit_widgets.NewEditTab(dApp.mainWindow, dApp.managers)),
			container.NewTabItem("Удалить", delete_widgets.NewDeleteTab(dApp.mainWindow, dApp.managers)),
			container.NewTabItem("Пользователи", user_widgets.NewUserTab(dApp.mainWindow, dApp.managers)),
		)
		tabs.SetTabLocation(container.TabLocationLeading)
		dApp.mainWindow.SetContent(tabs)
	}

}

func (dApp DesktopApp) ShowAndRun() {
	dApp.mainWindow.ShowAndRun()
}

func (dApp DesktopApp) Show() {
	dApp.mainWindow.Show()
}

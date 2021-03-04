const {app, Tray, Notification, ipcMain} = require("electron")
const database = require("./database")
const {trayMenu} = require("./tray")
const {createWindowFunction} = require("./windowsCreator")

let mainWindow, tray, startNotification;


function initTrayApp() {
    tray = new Tray("Logo.png")
    tray.setTitle("Ha iniiado el Web Monitor")
    tray.setToolTip("Web Monitor")
    tray.setContextMenu(trayMenu)

    tray.on("click", createWindowFunction({
        width: 1000,
        height: 800,
        minWidth: 800,
        minHeight: 600,
        show: false,
        webPreferences: {
            nodeIntegration: true,
        },
    }, "./vistas/main/index.html", mainWindow))

    startNotification = new Notification({
        title: "Web Monitor Iniciado",
        subtitle:"Da click para ver el estado de tus páginas",
        body: "Se ha comenzado con el monitoreo de tus páginas :)",
        icon: "Logo.png"
    })
    startNotification.show()
    startNotification.on("click", createWindowFunction({
        width: 1000,
        height: 800,
        minWidth: 800,
        minHeight: 600,
        show: false,
        webPreferences: {
            nodeIntegration: true,
        },
    }, "./vistas/main/index.html", mainWindow))
}

app.whenReady().then(initTrayApp)

app.on("window-all-closed", (e) => {
    e.preventDefault()
})

ipcMain.handle("getAllPages", (e) => {
    return database.getAllPages()
})
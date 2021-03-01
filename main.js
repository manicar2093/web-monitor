const {app, Tray, Menu, BrowserWindow, Notification} = require("electron")

let mainWindow, tray, startNotification, finalClose;

let trayMenu = Menu.buildFromTemplate([
    {label: "Salir", click: () => {
        finalClose = true
        app.quit()
    }}
])

function initTrayApp() {
    tray = new Tray("Logo.png")
    tray.setToolTip("Web Monitor")
    tray.setContextMenu(trayMenu)

    tray.on("click", createMainWindow)

    startNotification = new Notification({
        title: "Web Monitor Iniciado",
        subtitle:"Da click para ver el estado de tus páginas",
        body: "Se ha comenzado con el monitoreo de tus páginas :)",
        icon: "Logo.png"
    })
    startNotification.show()
    startNotification.on("click", createMainWindow)
}

function createMainWindow() {
    if(mainWindow) {
        mainWindow.show()
        return
    }
    mainWindow = new BrowserWindow({
        maxWidth: 800,
        minWidth: 800,
        maxHeight: 600,
        minHeight: 600,
        show: false,
        webPreferences: {
            nodeIntegration: true,
        },
    })

    mainWindow.webContents.openDevTools()

    mainWindow.loadFile("index.html")
    mainWindow.on("ready-to-show", mainWindow.show)
    mainWindow.on("close", (e) => {
        if(finalClose) {
            return
        }
        e.preventDefault()
        mainWindow.hide()
    })
}

app.whenReady().then(initTrayApp)

app.on("window-all-closed", () => {
    if(process.platform !== "darwin") {
        app.quit()
    }
})
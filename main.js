const {app, Tray, Menu, BrowserWindow, Notification, dialog, ipcMain} = require("electron")
const database = require("./database")

let mainWindow, tray, startNotification;


let trayMenu = Menu.buildFromTemplate([
    {label: "Salir", click: async () => {

        let options = ["Si", "Cancelar"]

        let r = await dialog.showMessageBox(mainWindow, {
            message: "¿Deseas cerrar la aplicación?",
            buttons: options,
            title: "Cerrar Web Monitor",
            icon: "Logo.png"
        })

        if(options[r.response] === "Si") {

            app.quit()
        }
    }}
])

function initTrayApp() {
    tray = new Tray("Logo.png")
    tray.setTitle("Ha iniiado el Web Monitor")
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
    
    mainWindow = new BrowserWindow({
        width: 1000,
        height: 800,
        minWidth: 800,
        minHeight: 600,
        show: false,
        webPreferences: {
            nodeIntegration: true,
        },
    })

    mainWindow.webContents.openDevTools()

    mainWindow.loadFile("index.html")
    mainWindow.on("ready-to-show", mainWindow.show)

}

app.whenReady().then(initTrayApp)

app.on("window-all-closed", (e) => {
    e.preventDefault()
})

ipcMain.handle("getAllPages", (e) => {
    return database.getAllPages()
})
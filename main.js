const {app, Tray, Notification, ipcMain, dialog, BrowserWindow} = require("electron")
const database = require("./dao")
const {trayMenu} = require("./tray")
const {createWindowFunction, getImageFromWindow} = require("./windowsCreator")

require("./cron")

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

ipcMain.handle("getPageById", (e, id) => {
    return database.getPageById(id)
})

ipcMain.handle("deletePage", async (e, id) => {
    try {
        database.deletePagina(id)
        return
    } catch (e){
        console.error(e)
        dialog.showErrorBox("Error al eliminar página", "Hubo un problema al eliminar la página solicitada.")
        return
    }
})

ipcMain.handle("savePage", async (e, data) => {
    let imageUrl = await getImageFromWindow(data.url)
    data.image = imageUrl
    data.status = true
    try {
        database.savePage(data)
        dialog.showMessageBox(BrowserWindow.getFocusedWindow(), {
            message:"Pagina guardada correctamente",
            title: "Éxito"
        })
    } catch (error) {
        console.error(error)
        dialog.showErrorBox("Error al eliminar página", "Hubo un problema al eliminar la página solicitada.")
        return
    }
})

ipcMain.handle("updatePage", async (e, data) => {
    data.image = await getImageFromWindow(data.url)
    data.status = true
    try {
        database.updatePage(data)
        dialog.showMessageBox(BrowserWindow.getFocusedWindow(), {
            message:" Página actualizada correctamente",
            title: "Éxito"
        })
        
    } catch (error) {
        console.error(error)
        dialog.showErrorBox("Error al actualizar página", "Hubo un problema al actualizar la página solicitada.")
        return
    }
})

ipcMain.handle("getAllFrases", (e)=>{
    return database.getAllFrases()
})


ipcMain.handle("getAllPhrases", e => {
    return database.getAllFrases()
})

ipcMain.handle("getPhraseById", (e, id) => {
    return database.getPhraseById(id)
})

ipcMain.handle("savePhrase", async (e, data) => {
    try {
        database.saveFrase(data.frase)
        dialog.showMessageBox(BrowserWindow.getFocusedWindow(), {
            title: "Éxito",
            message: "La frase se ha creado correctamente"
        })
    } catch (error) {
        console.error(error)
        dialog.showErrorBox("Error al crear frase", "Hubo un problema al crear la frase solicitada.")
        return
    }
})

ipcMain.handle("updatePhrase", async (e, data) => {
    try {
        database.updateFrase(data.id, data.frase)
        dialog.showMessageBox(BrowserWindow.getFocusedWindow(), {
            title: "Éxito",
            message: "Frase actualizada!"
        })
    } catch (error) {
        console.error(error)
        dialog.showErrorBox("Error al actualizar frase", "Hubo un problema al actualizar la frase solicitada.")
        return
    }

})

ipcMain.handle("deletePhrase", (e, id) => {
    try {
        database.deleteFrase(id)
        return
    } catch (error) {
        console.error(e)
        dialog.showErrorBox("Error al eliminar frase", "Hubo un problema al eliminar la frase solicitada.")
        return
    }
})
const { Menu, BrowserWindow, dialog, app } = require("electron")

let pageRegistryWin, phraseRegistryWin

const createDialogExitApp = async () => {

    let options = ["Cancelar","Si"]

    let r = await dialog.showMessageBox({
        message: "¿Deseas cerrar la aplicación?",
        buttons: options,
        title: "Cerrar Web Monitor",
        icon: "Logo.png",

    })

    if(options[r.response] === "Si") {

        app.quit()
    }
}

const createRegistryPageWindow = () => {
    pageRegistryWin = new BrowserWindow({
        width: 1000,
        height: 800,
        minWidth: 800,
        minHeight: 600,
        show: false,
        webPreferences: {
            nodeIntegration: true,
        },
    })

    pageRegistryWin.webContents.openDevTools()

    pageRegistryWin.loadFile("./vistas/pageRegistry/pageRegistry.html")
    pageRegistryWin.on("ready-to-show", pageRegistryWin.show)
    pageRegistryWin.on("closed", () => {
        pageRegistryWin = null
    })
}

const createRegistryPhraseWindow = () => {
    phraseRegistryWin = new BrowserWindow({
        width: 1000,
        height: 800,
        minWidth: 800,
        minHeight: 600,
        show: false,
        webPreferences: {
            nodeIntegration: true,
        },
    })

    phraseRegistryWin.webContents.openDevTools()

    phraseRegistryWin.loadFile("./vistas/phraseRegistry/phraseRegistry.html")
    phraseRegistryWin.on("ready-to-show", phraseRegistryWin.show)
    phraseRegistryWin.on("closed", () => {
        phraseRegistryWin = null
    })
}

const trayMenu = Menu.buildFromTemplate([
    {label: "Salir", click: createDialogExitApp},
    {label: "Registra Página", click: createRegistryPageWindow},
    //{label: "Registrar Frase", click: createRegistryPhraseWindow}
])

module.exports = {
    createDialogExitApp,
    createRegistryPageWindow,
    createRegistryPhraseWindow,
    trayMenu
}
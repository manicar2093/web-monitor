const { Menu, dialog, app } = require("electron")
const { createWindowFunction } = require("./windowsCreator")

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

const createRegistryPageWindow = createWindowFunction({
    width: 990,
    height: 700,
    minWidth: 990,
    minHeight: 700,
    maxWidth: 990,
    maxHeight: 700,
    maximizable: false,
    show: false,
    webPreferences: {
        nodeIntegration: true,
    },
}, "./vistas/pageRegistry/pageRegistry.html", pageRegistryWin)

const createRegistryPhraseWindow = createWindowFunction({
    width: 990,
    height: 700,
    minWidth: 990,
    minHeight: 700,
    maxWidth: 990,
    maxHeight: 700,
    maximizable: false,
    show: false,
    webPreferences: {
        nodeIntegration: true,
    },
}, "./vistas/phraseRegistry/phraseRegistry.html", phraseRegistryWin)

const trayMenu = Menu.buildFromTemplate([
    {label: "Registra Página", click: createRegistryPageWindow},
    {label: "Registrar Frase", click: createRegistryPhraseWindow},
    {label: "Salir", click: createDialogExitApp},
])

module.exports = {
    createDialogExitApp,
    createRegistryPageWindow,
    createRegistryPhraseWindow,
    trayMenu
}
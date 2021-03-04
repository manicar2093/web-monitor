const { BrowserWindow } = require("electron")

/**
 * 
 * Toma los datos solicitados y regresa una función para la creación de una ventana.
 * 
 * @param {Object} options Opciones para la creación de un BrowserWindow
 * @param {String} file Path donde se encuentra el archivo que se debe cargar en BrowserWindow
 * @param {Variable} target Variable donde se almacenara el objeto que se genere
 */
module.exports.createWindowFunction = (options, file, target) => {

    return () => {
        target = new BrowserWindow(options)

        target.webContents.openDevTools()

        target.loadFile(file)
        target.on("ready-to-show", target.show)
        target.on("closed", () => {
            target = null
        })
    }
}
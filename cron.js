const axios = require("axios")
const cron = require("node-cron")
const {Notification} = require("electron")
const {getAllPages, updatePage, getAllPagesWithStatusFalse, getAllFrases} = require("./dao")
const {createWindowFunction} = require("./windowsCreator")

function randomInteger(min, max) {
    max = max - 1
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

let notified = false

exports.setNotified = (data) => {
    notified = data
}

cron.schedule("*/40 * * * * *", async () => {
    let pages = await getAllPages()
    pages.forEach((curr)=> {
        axios.get(curr.url)
        .catch(error => {
            curr.status = false
            updatePage(curr)
        })
    })
    let pageError = await getAllPagesWithStatusFalse()
    if(pageError.length > 0 && !notified) {
        let frases = await getAllFrases()
        let index = randomInteger(0, frases.length)
        console.log(index, frases.length)
        let frase = frases[index]
        const notif = {
            title: "Hay paginas fuera de servicio",
            body: frase.frase
        }
        let mainWindow
        let errorNotif = new Notification(notif)
        errorNotif.show()
        errorNotif.on("click", createWindowFunction({
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
})
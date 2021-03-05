const { ipcRenderer } = require("electron")
let registros = document.getElementById("registrados")

let pageForm = document.getElementById("pageForm")
let pageName = document.getElementById("nombre")
let pageUrl = document.getElementById("url")
let pageId = document.getElementById("id")

let moveType = "SAVE"



async function deleteButton(id) {
    if (!confirm("¿Seguro quieres eliminar la frase?")) {
        return
    }
    await ipcRenderer.invoke("deletePage", id)
}

function savePage() {
    if (moveType == "UPDATE") {
        let data = {
            id: pageId.value,
            url: pageUrl.value,
            name: pageName.value,
        }
        // TODO : Terminar logica de guardado y actualización
    }
}

function updateButton(id) {
    console.log("A actualizar:",id)
}

const createRegistros = (rows)=> {
    registros.innerHTML = /*html*/ `
        <h1>Páginas Registradas </h1>
    `
    if (rows.length <= 0) {
        registros.innerHTML += /* html */`
        <div class="registry">
            <p class="text_center">Sin páginas registradas</p>
        </div>
        `
        return
    }
    rows.forEach(i => {
        registros.innerHTML += /* html */`
        <div class="registry">
            <p class="text_center italic">${i.name}</p>
            <p class="text_center italic">${i.url}</p>
            <div class="buttons">
                <button class="button button--primary" onclick="updateButton(${i.id})">Actualizar</button>
                <button class="button button--danger" onclick="deleteButton(${i.id})">Eliminar</button>
            </div>
        </div>
        `
    })
}



ipcRenderer.invoke("getAllPages").then(rows => {
    createRegistros(rows)
})



const { ipcRenderer } = require("electron")
let registros = document.getElementById("registrados")

let pageForm = document.getElementById("pageForm")
let pageName = document.getElementById("nombre")
let pageUrl = document.getElementById("url")
let pageId = document.getElementById("id")

const SAVE = "SAVE"
const UPDATE = "UPDATE"
let moveType = SAVE

pageForm.addEventListener("submit", async (e) => {
    e.preventDefault()
    await savePage()
    pageForm.reset()
})

async function deleteButton(id) {
    if (!confirm("¿Seguro quieres eliminar la frase?")) {
        return
    }
    await ipcRenderer.invoke("deletePage", id)
    await createRegistros()
}

async function savePage() {
    let data = {
        id: pageId.value,
        url: pageUrl.value,
        name: pageName.value,
    }
    switch(moveType) {
        case UPDATE:
            await ipcRenderer.invoke("updatePage", data)
            moveType = SAVE
            break
        case SAVE:
            await ipcRenderer.invoke("savePage", data)
            moveType = SAVE
            break
    }

    await createRegistros()
    
}

async function updateButton(id) {
    let registry = await ipcRenderer.invoke("getPageById")
    pageName.value = registry.name
    pageUrl.value = registry.url
    pageId.value = registry.id
    moveType = UPDATE
    
}

const createRegistros = async ()=> {
    let rows = await ipcRenderer.invoke("getAllPages")
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

createRegistros()

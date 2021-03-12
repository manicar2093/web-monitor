const { ipcRenderer } = require("electron")

let registrados = document.getElementById("registrados")
let fraseId = document.getElementById("id")
let frase = document.getElementById("frase")
let fraseForm = document.getElementById("fraseForm")

let submitButton = document.querySelector("button[type='submit']")

const SAVE = "SAVE"
const UPDATE = "UPDATE"
let moveType = SAVE

fraseForm.addEventListener("submit", async (e) => {
    e.preventDefault()
    submitButton.setAttribute("disabled", true)
    if (frase.value == ""){
        alert("Los datos no son validos. Favor de llenar de forma correcta")
        submitButton.removeAttribute("disabled")
        return
    }
    await savePage()
    fraseForm.reset()
    submitButton.removeAttribute("disabled")
})


async function savePage() {
    let data = {
        id: fraseId.value,
        frase: frase.value,
    }
    switch(moveType) {
        case UPDATE:
            await ipcRenderer.invoke("updatePhrase", data)
            break
        case SAVE:
            await ipcRenderer.invoke("savePhrase", data)
            break
    }

    moveType = SAVE
    await createRegisters()
    
}

const updateButton = async (id) => {
    let registry = await ipcRenderer.invoke("getPhraseById", id)
    frase.value = registry.frase
    fraseId.value = registry.id
    moveType = UPDATE
}

const deleteButton = async (id) => {
    if(!confirm("¿Estas segura de querer eliminar la frase?")) {
        return
    }
    await ipcRenderer.invoke("deletePhrase", id)
    await createRegisters()
}

const createRegisters = async () => {
    const phrases = await ipcRenderer.invoke("getAllPhrases")
    registrados.innerHTML = /*html*/ `<h1>Frases Registradas</h1>`
    phrases.forEach(i => {
        registrados.innerHTML += /*html*/ `
        <div class="registry">
            <p class="text_center italic">${i.frase}</p>
            <div class="buttons">
                <button class="button button--primary" onclick="updateButton(${i.id})">Actualizar</button>
                <button class="button button--danger" onclick="deleteButton(${i.id})">Eliminar</button>
            </div>
        </div>`
    })
}


createRegisters()



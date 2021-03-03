const {ipcRenderer} = require("electron")

const container = document.getElementById("pages")

const createCards = (rows) => {
    if(rows.length <= 0){
        container.innerHTML += `
        <div class="column">
            <div class="card success">
                <div class="card__title">
                    <h2 class="text_center">Aun no hay paginas registradas</h2>
                </div>
            </div>
        </div>
        `
        return
    }

    rows.forEach(i => {
        container.innerHTML += `
        <div class="column">
            <div class="card ${i.status ==true ? 'success' : 'danger'}">
                <div class="card__title">
                    <h2 class="text_center">${i.name}</h2>
                </div>
                <div class="card__img"> <img src="${i.image}"></div>
                <div class="card__content text_center">
                    <p>${i.url}</p>
                    <div class="buttons">
                        <button class="button button--primary">Actualizar</button>
                        <button class="button button--info">Visitar</button>
                        <!--button.button.button--danger Eliminar-->
                    </div>
                </div>
            </div>
        </div>
        `
    })
}
ipcRenderer.invoke("getAllPages").then((rows)=>{
    createCards(rows)
})
const sqlite = require("sqlite3").verbose()
const db = new sqlite.Database("database.db")

/**
 * Realiza la busqueda de todas las frases registradas en la base de datos. Regresa una promesa
 */
exports.getAllFrases = () => {
    return new Promise((res, rej) => {
        db.all("SELECT * FROM frases", (e, rows) => {
            if (e) rej(e)
            else res(rows)
        })
    })
}

/**
 * Obtiene una frase por su ID
 * @param {Number} id Identificador de la frase deseada
 * @returns Objeto con los datos de la frase
 */
 exports.getPhraseById = (id) => {
    return new Promise((res, rej)=>{
        db.get("SELECT * FROM frases WHERE id=?", [id], (e, row) => {
            if (e) rej(e)
            else res(row)
        })
    })
}

/**
 * 
 * Realiza el registro de una frase en la base de datos
 * 
 * @param {String} frase Dato que se guardará en la base de datos
 */
exports.saveFrase = (frase) => {
    return new Promise((res, rej) => {
        db.run("INSERT INTO frases (frase) VALUES (?)", [frase], (e, data) => {
            if (e) rej(e)
            else res(data)
        })
    })
}

/**
 * Elimina una frase de la base de datos
 * @param {Any} id Identificador de la frase que se debe eliminar
 */
exports.deleteFrase = (id) => {
    return new Promise((res, rej) => {
        db.run("DELETE FROM frases WHERE id = ?", id, (e, data) => {
            if (e) rej(e)
            else res(data)
        })
    })
}

/**
 * 
 * Realiza la actualización de una frase
 * 
 * @param {Any} id Identificador de la frase a actualizar
 * @param {String} frase Nueva frase
 */
exports.updateFrase = (id, frase) => {
    return new Promise((res, rej) => {
        db.run("UPDATE frases SET frase = ? WHERE id = ?", [frase, id], (e, data) => {
            if (e) rej(e)
            else res(data)
        })
    })
}

/**
 * Regresa todas las paginas que están en la base de datos
 */
exports.getAllPages = () => {
    return new Promise((res, rej) => {
        db.all("SELECT * FROM paginas", (e, rows) => {
            if (e) rej(e)
            else res(rows)
        })
    })
}

/**
 * Obtiene una pagina por su ID
 * @param {Number} id Identificador de la página deseada
 * @returns Objeto con los datos de la pagina
 */
exports.getPageById = (id) => {
    console.log("ID ha buscar:", id)
    return new Promise((res, rej)=>{
        db.get("SELECT * FROM paginas WHERE id=?", [id], (e, row) => {
            if (e) rej(e)
            else res(row)
        })
    })
}

/**
 * 
 * Realiza el guardado de una pagina
 * 
 * @param {Object} data URL, name, image y status que se debe registrar
 * @param {Function} callback Función que se corre
 */
exports.savePage = (data) => {

    return new Promise((res, rej)=> {

        db.run("INSERT INTO paginas (url,image,status, name) VALUES (?,?,?,?)", [data.url, data.image, data.status, data.name], (e, data) => {
            if (e) rej(e)
            else res(data)
        })
    })
}

/**
 * Elimina una página de la base de datos
 * @param {Any} id Identificador de la frase que se debe eliminar
 */
exports.deletePagina = (id) => {
    return new Promise((res, rej) => {
        db.run("DELETE FROM paginas WHERE id = ?", id, (e, data) => {
            if (e) rej(e)
            else res(data)
        })
    })
}

/**
 * Realiza el cambio de status del id que se envie
 * @param {Object} data status y id que se usaran para el update
 * @param {Function} callback Funcion que se corre
 */
exports.updatePage = (data) => {
    return new Promise((res, rej) => {
        db.run("UPDATE paginas SET url=?, image=?, status=?, name=? WHERE id = ?", [data.url, data.image, data.status, data.name, data.id], (e, data) => {
            if (e) rej(e)
            else res(data)
        })
    })
}


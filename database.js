const sqlite = require("sqlite3").verbose()
const db = new sqlite.Database("database.db")

exports.getAllFrases = (callback) => {
    db.all("SELECT * FROM frases", callback)
}

exports.saveFrase = (frase, callback) => {
    db.run("INSERT INTO frases (frase) VALUES (?)", [frase], callback)
}

exports.getAllPages = () => {
    return new Promise((res, rej) => {
        db.all("SELECT * FROM paginas", (e, rows) => {
            if (e) {
                rej(e)
            } else {
                res(rows)
            }

        })
    })
}

/**
 * 
 * @param {Array} data URL, name, image y status que se debe registrar
 * @param {Function} callback Función que se corre
 */
exports.savePage = (data, callback) => {
    db.run("INSERT INTO paginas (url,image,status, name) VALUES (?,?,?,?)", data, callback)
}

/**
 * Realiza el cambio de status del id que se envie
 * @param {Object} data status y id que se usaran para el update
 * @param {Function} callback Funcion que se corre
 */
exports.updatePage = (data, callback) => {
    db.run("UPDATE paginas SET status=? WHERE id = ?", [data.status, data.id], callback)
}


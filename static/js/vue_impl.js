const app = new Vue({
    el: "#app",
    delimiters: ['[[', ']]'],
    data: {
        registerPageForm: {
            name: '',
            url: '',
        },
        registerPhraseForm: {
            phrase: ''
        },
        activedPages: [],
        unactivedPages: [],
        pages: [],
        phrases: [],
        notifications_accepted: false,
        appClosed: false,
        evtSource: null,
        validationInterval: null,
    },
    methods: {
        async requestNotificationsPermissions() {
            const messageNegative = "Los permisos son necesarios para la aplicación"
            const result = await Notification.requestPermission()

            switch (result) {
                case "default":
                    alert(messageNegative)
                    await this.requestNotificationsPermissions()
                    break
                case "denied":
                    alert(messageNegative)
                    break
                case "granted":
                    this.notifications_accepted = true
                    break
            }
        },
        async getPages() {
            const data = await fetch('/pages/all')
            this.pages = await data.json()
            this.unactivedPages = this.pages.filter(i => i.is_working == false)
            this.activedPages = this.pages.filter(i => i.is_working == true)
        },
        async getPhrases() {
            const data = await fetch('/phrases/all')
            this.phrases = await data.json()
        },
        async show(data) {
            if (data == 'phrase_admin') {
                await this.getPhrases()
            }
            this.$refs[data].classList.toggle("is-active")
            // console.log(this.$refs[data].classList)
        },
        async createPage() {
            const onError = (accion) => alert(`ERROR AL ${accion} PAGINA.\nIntenta nuevamente.`)
            if (this.registerPageForm.id) {
                try {
                    const res = await fetch("/pages/update", {
                        body: JSON.stringify(this.registerPageForm),
                        method: 'PUT',
                        headers: {
                            'Content-Type': 'application/json'
                        }
                    })
                    if (res.status != 200) {
                        onError()
                        return
                    }
                    alert("¡Listo!\nPágina actualizada.")
                    await this.getPages()
                    this.clearForm('registerPageForm')
                } catch (error) {
                    console.log(error)
                    onError()
                }
                return
            }
            try {
                const res = await fetch("/pages/add", {
                    body: JSON.stringify(this.registerPageForm),
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    }
                })
                if (res.status != 200) {
                    console.log(error)
                    onError()
                    return
                }
                
                const {page} = await res.json()
                if(!page.is_working) {
                    alert("¡Listo!\nPágina registrada, pero no respondió")
                }else {
                    alert("¡Listo!\nPágina registrada.")
                }
                await this.getPages()
                this.clearForm('registerPageForm')
                this.show('page_admin')
            } catch (error) {
                console.log(error)
                onError("ACTUALIZAR")
            }

        },
        async createPhrase() {
            const onError = () => alert("ERROR AL REGISTRAR FRASE.\nIntenta nuevamente.")
            if (this.registerPhraseForm.id) {
                try {
                    const res = await fetch("/phrases/update", {
                        body: JSON.stringify(this.registerPhraseForm),
                        method: 'PUT',
                        headers: {
                            'Content-Type': 'application/json'
                        }
                    })
                    if (res.status != 200) {
                        onError()
                        return
                    }
                    alert("¡Listo!\nFrase actualizada.")
                    this.clearForm('registerPhraseForm')
                    await this.getPhrases()
                } catch (error) {
                    onError()
                }
                return
            }
            try {
                const res = await fetch("/phrases/add", {
                    body: JSON.stringify(this.registerPhraseForm),
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    }
                })
                if (res.status != 200) {
                    onError()
                    return
                }
                alert("¡Listo!\Frase registrada.")
                this.clearForm('registerPhraseForm')
                this.show('phrase_admin')
            } catch (error) {
                onError()
            }


        },
        clearForm(formID, askConfirmation = false) {
            if (askConfirmation && !confirm("¿Seguro desea borrar la información ya ingresada?")) {
                return
            }
            const forms = {
                registerPageForm: {
                    name: '',
                    url: '',
                    status: ''
                },
                registerPhraseForm: {
                    phrase: ''
                }
            }
            this[formID] = forms[formID]
        },
        async onErase(data) {
            if (data.type == "page") {
                await this.deletePage(data)
                return
            }

            await this.deletePhrase(data)
            return

        },
        onUpdateRequest(data) {
            if (data.type == 'page') {
                this.registerPageForm = data
                return
            }
            this.registerPhraseForm = data
        },
        async closeApp() {

            try {
                await fetch("/close", {
                    method: "POST"
                })
            } catch (error) {
                this.notifications_accepted = false
                this.appClosed = true
            }
        },
        async deletePage(data) {
            const onError = () => alert("ERROR AL BORRAR PAGINA.\nIntenta nuevamente.")
            try {
                if(!confirm("¡Estas seguro en eliminar la página?")){
                    return
                }
                const res = await fetch("/pages/delete", {
                    method: "DELETE",
                    body: JSON.stringify(data),
                    headers: {
                        'Content-Type': "application/json"
                    }
                });
                if (res.status != 200) {
                    onError()
                    return
                }
                alert("¡Listo!\nPágina eliminada.")
                await this.getPages()
            } catch (error) {
                console.error(error)
                onError()
                return
            }
        },
        async deletePhrase(data) {
            const onError = () => alert("ERROR AL BORRAR FRASE.\nIntenta nuevamente.")
            try {
                const res = await fetch("/phrases/delete", {
                    method: "DELETE",
                    body: JSON.stringify(data),
                    headers: {
                        'Content-Type': "application/json"
                    }
                });
                if (res.status != 200) {
                    onError()
                    return
                }
                alert("¡Listo!\Frase eliminada.")
                await this.getPhrases()
            } catch (error) {
                onError()
                return
            }
        },
        async sseHandler(e) {
            console.log("Creating notification")
            let data = JSON.parse(e.data)
            if(!data) { // avoid process when there is no data in streaming
                console.log("No data in strem")
                return
            }
            let {page} = data
            if(data.page.recovered) {
                new Notification("¡Recuperada!", {
                    body: `La pagina '${page.name}' ya respondio :D`
                })
            } else if (!data.page.is_working) {
                const phraseIndex = this.getRandomArbitrary()
                new Notification(this.phrases[phraseIndex].phrase, {
                    body: `La pagina '${page.name}' no responde`
                })
            } else if(data.page_id) {
                new Notification(`Validación fallida`, {
                    body: `Un error al validar la pagina '${page.name}'`
                })
            }
            await this.getPages()

        },
        getPageByID(id) {
            return this.pages.find((i) => i.id == id)
        },
        getRandomArbitrary() {
            return Math.floor(Math.random() * (this.phrases.length));
        },
        async validatePage(page) {
            const res = await fetch('/pages/validate', {
                body: JSON.stringify({
                    id: page.id,
                    url: page.url
                }),
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                }
            })
            const {page:PageRes} = await res.json()
            console.log(PageRes)
            if(PageRes.recovered) {
                alert(`La página '${PageRes.name}' fue recuperada :D`)
                await this.getPages()
                return
            } else if (!PageRes.is_working){
                alert(`La página '${PageRes.name}' aún no responde D:`)
            }
        },
        setIntervalHandler(){
            if(!this.unactivedPages.length > 0) {
                return
            }
            new Notification(`Recuerda...`, {
                body: `Aún hay ${this.unactivedPages.length} página(s) sin responder :/`
            })
        }
    },
    async created() {
        await this.requestNotificationsPermissions()
        await this.getPages()
        await this.getPhrases()

        this.evtSource = new EventSource("/sse/sse-validator");
        this.evtSource.onmessage = this.sseHandler
        // this.evtSource.onopen = (e) => console.log("Connected to SSE")
        // this.evtSource.onerror = (e) => console.error(e)
        this.validationInterval = setInterval(this.setIntervalHandler, 120000)
    }
})
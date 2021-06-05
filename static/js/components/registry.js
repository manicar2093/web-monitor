Vue.component("registry", {
    props: ["type", "status", "data", "show_options"],
    delimiters: ['[[', ']]'],
    template: /*html*/ `
    <div :class="generalClasses">
        <div v-if="status == 'down'">
            <div class="registry__time">
                <p>00:00:00</p>
            </div>
        </div>
        
        <div class="registry__page_name">
            <div v-if="type == 'phrase'" >
                [[ data.phrase ]]
            </div>

            <div v-else-if="type == 'page'">
                <a :href="data.url" target="_blank">[[ data.name ]]</a>
            </div>
            
        </div>
        <div class="registry__actions">
            <button v-if="type == 'page' && status == 'down'" @click="validatePage" class="button is-info" title="Revisamos el status de la página">Actualizar</button>
            <div v-if="show_options">
                <button class="button is-danger" @click="erase">Eliminar</button>
                <button class="button is-warning" @click="update">Editar</button>
            </div>
            
        </div>
    </div>
    `,
    methods: {
        validatePage() {
            this.$emit("verifyPage", {...this.data, type: this.type})
        },
        erase( ) {
            this.$emit("erase", {...this.data, type: this.type})
        },
        update() {
            this.$emit("update", {...this.data, type: this.type})
        }
    },
    computed: {
        generalClasses() {
            if(!this.status) {
                return "registry"
            }
            let registryType = `registry--${this.status}`
            return `registry ${registryType}`
        },
    }
})
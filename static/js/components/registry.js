Vue.component("registry", {
    props: ["type", "page_name", "url"],
    delimiters: ['[[', ']]'],
    template: /*html*/ `
    <div :class="generalClasses">
        <div v-if="verifyType">
            <div class="registry__time">
                <p>00:00:00</p>
            </div>
        </div>
        <div class="registry__page_name"><a :href="url" target="_blank">[[ page_name ]]</a></div>
        <div v-if="verifyType">
            <div class="registry__actions">
                <button @click="validatePage" class="button is-info" title="Revisamos el status de la página">Actualizar</button>
            </div>
        </div>
    </div>
    `,
    methods: {
        validatePage() {
            this.$emit("verifyPage", this.url)
        }
    },
    computed: {
        generalClasses() {
            if(!this.type) {
                return "registry"
            }
            let registryType = `registry--${this.type}`
            return `registry ${registryType}`
        },
        verifyType() {
            return this.type == 'down' || !this.type
        }
    }
})
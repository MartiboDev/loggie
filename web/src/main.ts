import { createApp } from "vue"
import { createPinia } from "pinia"
import PrimeVue from "primevue/config"

import "primevue/resources/themes/aura-dark-green/theme.css"
import "primeicons/primeicons.css"
import "./assets/main.scss"

import App from "./App.vue"
import router from "./router"

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(PrimeVue)

app.mount("#app")

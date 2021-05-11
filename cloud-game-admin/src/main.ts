import { createApp } from 'vue'
import vuetify from './plugins/vuetify'
import App from './App.vue'

// router-next
import { router } from './router'

const app = createApp(App)
app.use(vuetify)

app.mount('#app')

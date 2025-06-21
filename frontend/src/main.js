import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import '@fortawesome/fontawesome-free/css/all.min.css'
import { track, installAnalytics } from './services/analytics'

createApp(App).use(router).mount('#app')
const app = createApp(App)

app.config.globalProperties.$track = track
app.use(router)
installAnalytics(router)
app.mount('#app')
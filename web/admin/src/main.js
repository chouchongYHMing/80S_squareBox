import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import  './plugin/antui'

axios.defaults.baseURL = 'http://localhost:3000/api/v1'
Vue.prototype.$http = axios
Vue.config.productionTip = false

createApp(App).use(router).mount('#app')

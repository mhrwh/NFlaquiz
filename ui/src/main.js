import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import axios from 'axios'

import 'bootstrap/dist/js/bootstrap.min'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import "bootstrap-icons/font/bootstrap-icons.css"
import drag from "v-drag"
//import { loadFonts } from './plugins/webfontloader'
axios.defaults.withCredentials = true
//loadFonts()

createApp(App)
  .use(router)
  .use(store)
  .use(vuetify)
  .use(drag, {
    // global configuration
  })
  .mount('#app')

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import  store  from "./store"
import 'bootstrap/dist/js/bootstrap.min'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import "bootstrap-icons/font/bootstrap-icons.css"
import "../public/styles/overwrite.css"

createApp(App).use(store).use(router).mount('#app')

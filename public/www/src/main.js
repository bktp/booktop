import Vue from 'vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import App from './App.vue'
import router from './router'

import Bulma from 'bulma'

Vue.use(VueResource)
Vue.use(VueRouter)

Vue.http.options.credentials = true;

new Vue({
  el: '#app',
  router: router,
  render: h => h(App)
})

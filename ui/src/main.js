import Vue from 'vue'
import App from './App.vue'
import vueResource from 'vue-resource'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'

Vue.use(vueResource)

Vue.config.productionTip = false

Vue.use(Buefy)

new Vue({
  render: h => h(App),
}).$mount('#app')

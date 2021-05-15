import Vuelidate from 'vuelidate'
import Vue from 'vue';
import axios from 'axios'

import VueAxios from 'vue-axios'

Vue.use(VueAxios, axios)
Vue.use(Vuelidate)
import App from './App.vue';
import router from './router';
import store from './store';

Vue.config.productionTip = false;
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
Vue.use(BootstrapVue)
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin)
import './assets/scss/index.scss'

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');

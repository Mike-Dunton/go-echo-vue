import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import axios    from 'axios';
import VueAxios from 'vue-axios'
axios.defaults.baseURL = "http://dev.dntn.xyz:8000";
Vue.use(VueAxios, axios);

import auth            from '@websanova/vue-auth';
import authBasic       from '@websanova/vue-auth/drivers/auth/devise.js';
import httpVueResource from '@websanova/vue-auth/dist/drivers/http/axios.1.x.esm.js';
import routerVueRouter from '@websanova/vue-auth/dist/drivers/router/vue-router.2.x.esm.js';
import oauth2Google    from '@websanova/vue-auth/dist/drivers/oauth2/google.esm.js';
oauth2Google.params.client_id = '1058328058128-44m7tt45t86s9qbj8at3hv0ictvheo2m.apps.googleusercontent.com';
oauth2Google.params.scope = "openid " +
"https://www.googleapis.com/auth/userinfo.email " +
"https://www.googleapis.com/auth/userinfo.profile";

Vue.use(auth, {
  auth: authBasic,
  http: httpVueResource,
  router: routerVueRouter,
  oauth2: {
      google: oauth2Google
  },
  rolesKey: 'type',
  refreshData: {
    enabled: false
  }
});
 

import 'bulma/css/bulma.css'
import 'bulma-badge/dist/css/bulma-badge.min.css'

import { library } from '@fortawesome/fontawesome-svg-core'
import { faPlus, faUtensils } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faPlus, faUtensils)

Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.config.productionTip = false
Vue.config.devtools = true

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

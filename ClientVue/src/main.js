// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import Axios from 'axios'
import App from './App';
import helper from "./Plugins/helper";
import router from './router';
import swal from 'sweetalert2';
window.swal = swal;
import InputBase from "./components/Input";
import VeeValidate from 'vee-validate'
import PasswordMeter from "vue-simple-password-meter";
import { library } from '@fortawesome/fontawesome-svg-core';
import { faEye,faEyeSlash } from '@fortawesome/free-solid-svg-icons';
library.add(faEye,faEyeSlash);
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
Vue.use(VeeValidate);
Vue.component("Password",PasswordMeter);
Vue.component("Icons",FontAwesomeIcon);
Vue.component("InputBase",InputBase)

Vue.prototype.axios = Axios
Vue.prototype.helper = helper

Vue.config.productionTip = false


/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})

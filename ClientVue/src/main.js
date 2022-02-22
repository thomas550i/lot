import { createApp } from 'vue'
import App from './App.vue'
import router from "./router";
import RequiredComponents from "./Plugins/RequiredComponents";
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueSweetalert2 from 'vue-sweetalert2';
import 'sweetalert2/dist/sweetalert2.min.css';
import VueDatepickerUi from 'vue-datepicker-ui'
import 'vue-datepicker-ui/lib/vuedatepickerui.css';
import VueToast from 'vue-toast-notification';
import 'vue-toast-notification/dist/theme-sugar.css';

const appInstance = createApp(App);
appInstance.use(router);
appInstance.use(RequiredComponents);
appInstance.mount("#app");
appInstance.use(VueAxios, axios)
appInstance.use(VueSweetalert2);
appInstance.component('Datepicker', VueDatepickerUi)
appInstance.use(VueToast);

import Vue from 'vue'
import Router from 'vue-router'
import Register from "../views/Register"
import AuthLayout from "../layout/AuthLayout";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: AuthLayout
    },
    {
      path: '/Register',
      name: 'Register',
      component: Register
    },
  ]
})

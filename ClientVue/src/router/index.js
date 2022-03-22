import Vue from 'vue'
import Router from 'vue-router'
import Register from "../views/Register"
import EditProfile from "../views/EditProfile"
import UpcomingShows from "../views/UpcomingShows"
import AuthLayout from "../layout/AuthLayout";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: AuthLayout,
      children:[
        {
          path: '/',
          name: 'UpcomingShows',
          component: UpcomingShows,
        },
      ]
    },
    {
      path: '/Register',
      name: 'Register',
      component: Register
    },
    {
      path: '/edit',
      name: 'EditProfile',
      component: EditProfile
    },
  ]
})

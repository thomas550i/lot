import Vue from 'vue'
import Router from 'vue-router'
import Register from "../views/Register"
import EditProfile from "../views/EditProfile"
import UpcomingShows from "../views/UpcomingShows"
import SelectSlot from "../views/SelectSlot"
import SelectNumber from "../views/SelectNumber"
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
        {
          path: '/SelectSlot',
          name: 'SelectSlot',
          component: SelectSlot,
        },
        {
          path: '/SelectNumber',
          name: 'SelectNumber',
          component: SelectNumber,
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

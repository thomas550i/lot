import Vue from 'vue'
import Router from 'vue-router'
import Register from "../views/Register"
import EditProfile from "../views/EditProfile"
import UpcomingShows from "../views/UpcomingShows"
import SelectSlot from "../views/SelectSlot"
import SelectNumber from "../views/SelectNumber"
import Transaction from "../views/Transactions";
import AuthLayout from "../layout/AuthLayout";
import DashBoard from "../layout/DashBoard";

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
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashBoard,
      children:[
        {
          path: '/Transaction',
          name: 'Transaction',
          component: Transaction,
        },
      ]
    },
  ]
})

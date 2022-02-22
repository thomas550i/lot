import { createRouter, createWebHashHistory } from "vue-router";
import Login from "../views/Login"
import Register from "../views/Register"
import AuthLayout from "../layout/AuthLayout";


const routes = [
  {
    path: "/",
    name:"home",
    component: AuthLayout,
    children:[
      {
        path: "/Register",
        name:"Register",
        component: Register,
      },
      {
        path: "/Login",
        name:"Login",
        component: Login,
      },
    ],
  },
  
];

const router = createRouter({
  history: createWebHashHistory(),
  linkActiveClass: "active",
  routes,
});

export default router;

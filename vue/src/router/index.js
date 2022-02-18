import { createRouter, createWebHashHistory } from "vue-router";

import DashboardLayout from "@/layout/DashboardLayout";
import AuthLayout from "@/layout/AuthLayout";

import Dashboard from "../views/Dashboard.vue";
import Icons from "../views/Icons.vue";
import shows from "../views/shows.vue";
import Distributor from "../views/Distributor.vue";

import DataEntry from "../views/DataEntry.vue";
import approvedentry from "../views/approvedentry.vue";
import submitedentry from "../views/submitedentry.vue";
import myapprovedentry from "../views/myapprovedentry.vue";
import pendingapproval from "../views/pendingapproval.vue";
import sequencefinder from "../views/sequencefinder.vue";




import Login from "../views/Login.vue";
import Register from "../views/Register.vue";

const routes = [
  {
    path: "/",
    redirect: "/dashboard",
    component: DashboardLayout,
    children: [
      {
        path: "/dashboard",
        name: "dashboard",
        components: { default: Dashboard },
      },
      {
        path: "/Users",
        name: "icons",
        components: { default: Icons },
      },
      {
        path: "/shows",
        name: "shows",
        components: { default: shows },
      },
      {
        path: "/distributor",
        name: "distributor",
        components: { default: Distributor },
      },
      {
        path: "/dataentry",
        name: "dataentry",
        components: { default: DataEntry },
      },
      {
        path: "/submitedentry",
        name: "submitedentry",
        components: { default: submitedentry },
      },
      {
        path: "/approvedentry",
        name: "approvedentry",
        components: { default: approvedentry },
      },
      {
        path: "/myapprovedentry",
        name: "myapprovedentry",
        components: { default: myapprovedentry },
      },
      {
        path: "/pendingapproval",
        name: "pendingapproval",
        components: { default: pendingapproval },
      },
      {
        path: "/sequencefinder",
        name: "sequencefinder",
        components: { default: sequencefinder },
      },
    ],
  },
  {
    path: "/",
    redirect: "login",
    component: AuthLayout,
    children: [
      {
        path: "/login",
        name: "login",
        components: { default: Login },
      },
      {
        path: "/register",
        name: "register",
        components: { default: Register },
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

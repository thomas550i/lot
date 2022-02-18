<template>
  <div class="wrapper" :class="{ 'nav-open': $sidebar.showSidebar }">
    <side-bar
      :background-color="sidebarBackground"
      short-title="Argon"
      title="Argon"
    >
      <template v-slot:links>
        <sidebar-item
       
          :link="{
            name: 'Dashboard',
            icon: 'ni ni-tv-2 text-primary',
            path: '/dashboard',
          }"
        />

        <sidebar-item
         v-if="UR=='Super-Admin'"
          :link="{
            name: 'User Management',
            icon: 'ni ni-single-02 text-blue',
            path: '/users',
          }"
        />
        <sidebar-item
         v-if="UR=='Super-Admin'"
          :link="{
            name: 'Show Management',
            icon: 'ni ni-watch-time text-orange',
            path: '/shows',
          }"
        />
        <sidebar-item
         v-if="UR=='Super-Admin'"
          :link="{
            name: 'Distributor Management',
            icon: 'ni ni-single-02 text-yellow',
            path: '/distributor',
          }"
        />
        <sidebar-item
         v-if="UR=='Super-Admin'"
          :link="{
            name: 'Data Management and Sequence Finder',
            icon: 'ni ni-bullet-list-67 text-red',
            path: '/sequencefinder',
          }"
        />
        <sidebar-item
         v-if="UR=='Super-Admin'"
          :link="{
            name: 'Reports',
            icon: 'ni ni-single-copy-04 text-info',
            path: '/login',
          }"
        />

        <sidebar-item
         v-if="UR=='Quality-Checker'"
          :link="{
            name: 'Pending Approval',
            icon: 'ni ni-single-copy-04 text-info',
            path: '/pendingapproval',
          }"
        />

        <sidebar-item
         v-if="UR=='Quality-Checker'"
          :link="{
            name: 'Approved Entries',
            icon: 'ni ni-check-bold  text-info',
            path: '/myapprovedentry',
          }"
        />

  <sidebar-item
         v-if="UR=='Data-Entry-Operator'"
          :link="{
            name: 'Data Entry',
            icon: 'ni ni-bullet-list-67 text-info',
            path: '/dataentry',
          }"
        />

        <sidebar-item
         v-if="UR=='Data-Entry-Operator'"
          :link="{
            name: 'Submitted Entries',
            icon: 'ni ni-single-copy-04 text-info',
            path: '/submitedentry',
          }"
        />

        <sidebar-item
         v-if="UR=='Data-Entry-Operator'"
          :link="{
            name: 'Approved Entries',
            icon: 'ni ni-check-bold text-info',
            path: '/approvedentry',
          }"
        />

        <sidebar-item
        
          :link="{
            name: 'Logout',
            icon: 'ni ni-button-power text-red',
            path: '/login',
          }"
        />
       

       
      </template>
    </side-bar>
    <div class="main-content" :data="sidebarBackground">
      <dashboard-navbar></dashboard-navbar>

      <div @click="toggleSidebar">
        <!-- your content here -->
        <router-view></router-view>
        <content-footer v-if="!$route.meta.hideFooter"></content-footer>
      </div>
    </div>
  </div>
</template>
<script>
import DashboardNavbar from "./DashboardNavbar.vue";
import ContentFooter from "./ContentFooter.vue";
import helper from "../helper.js"

export default {
  components: {
    DashboardNavbar,
    ContentFooter,
  },
  data() {
    return {
      UR:helper.UR,
      sidebarBackground: "vue", //vue|blue|orange|green|red|primary
    };
  },
  beforeCreate(){
    helper.checklogin()
  },
  methods: {
    toggleSidebar() {
      if (this.$sidebar.showSidebar) {
        this.$sidebar.displaySidebar(false);
      }
    },
  },
};
</script>
<style lang="scss"></style>

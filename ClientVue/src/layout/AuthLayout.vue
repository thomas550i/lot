<template>
  <Header/>
  <div>
      <router-view>
          
      </router-view>
  </div>
  <Footer/>
</template>
<script>
import Footer from "./Footer.vue"
import Header from "./header.vue"
export default {
  
  name: "auth-layout",
  components:{
      Footer:Footer,
      Header:Header
  },
  data(){    
      return{
          year: new Date().getFullYear(),
      showMenu: false,
      }
  },
}
</script>
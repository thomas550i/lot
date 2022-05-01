<template>
  <div>
      <Header/>
          <router-view>
              
          </router-view>
      <Footer/>
  </div>
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
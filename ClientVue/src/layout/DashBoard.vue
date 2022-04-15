<template>
  <div>
    <Header/>
  <div class="container-fluid">
  <div class="row">
    <div class="col-3 col-md-2" style="margin-top:5%;margin-bottom:5%">
      <ul>
          <li>
              <router-link :to="'/Transaction'"><a>Transaction</a></router-link>
          </li>
          <li style="margin-top:20%;">
              <router-link :to="'/PurchaseTickets'"><a>PurchasedTickets</a></router-link>
          </li>
      </ul>
    </div>
    <div class="col-9 col-md-9">
      <router-view>
              
      </router-view>
    </div>
  </div>
</div> 
    <Footer/>
  </div>
</template>

<script>
import Footer from "./Footer.vue";
import Header from "./header.vue";
export default {
  name:"DashBoard",
  components:{
      Footer:Footer,
      Header:Header
  },
  data(){
    return{

    }
  },
  created(){

  },
  methods(){

  }
}
</script>

<style>
  .HalfContainer{
    width: 50%;
    height: 100%;
    margin: 0px;
    padding: 10px;
  }
  .position-right{

  }
</style>
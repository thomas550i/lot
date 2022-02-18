<template>
  <div class="row justify-content-center">
    <div class="col-lg-5 col-md-7">
      <div class="card bg-secondary shadow border-0">
        <div class="card-header bg-transparent">
          <div class="text-muted text-center mt-2 mb-3">
            <h1>DATA LOT</h1>
          </div>
         
        </div>
        <div class="card-body px-lg-5 py-lg-5">
          <div class="text-center text-muted mb-4">
            <small>sign in with credentials</small>
          </div>
          <form role="form">
            <base-input
              formClasses="input-group-alternative mb-3"
              placeholder="Email"
              addon-left-icon="ni ni-email-83"
              v-model="model.email"
            >
            </base-input>

            <base-input
              formClasses="input-group-alternative mb-3"
              placeholder="Password"
              type="password"
              addon-left-icon="ni ni-lock-circle-open"
              v-model="model.password"
            >
            </base-input>

            <input type="checkbox" style="width: 15px;height: 15px;margin-right: 5px;" v-model="remomber"/>
            <span class="text-muted"> Remember me</span>
            <div class="text-center">
              <!-- <base-button type="primary" class="my-4" @click="login">Sign in</base-button> -->
             <button type="button"  class="my-4 btn btn-primary" v-on:click="login">Sign in</button>
            </div>
          </form>
        </div>
      </div>
      <div class="row mt-3">
        <!-- <div class="col-6">
          <a href="#" class="text-light"><small>Forgot password?</small></a>
        </div> -->
        <div class="col-6 text-right">
          <router-link to="/register" class="text-light"
            >
            <!-- <small>Create new account</small> -->
            </router-link
          >
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import helper from "../helper.js"
export default {
  name: "login",
  data() {
    return {
      UserInfo:{},
      remomber:false,
      model: {
        email: "",
        password: "",
      },
    };
  },
  created(){
    var udata = localStorage.getItem("UUID");
    if (udata===null || udata==""){
      //do nothing
    } else {
      var t = JSON.parse(atob(udata))
      this.model.email=t.email
      this.model.password = t.password
      this.remomber = true
    }
  },
 methods: {
    login() {
      console.log("zxcasd");
     
      const article = {"Username":this.model.email,"Password":btoa(this.model.password)};
  this.axios.post(helper.SERVICEURL+"users/login", article)
    .then(response => {
      if(response.data.Success){
        helper.UserInfo = response.data.Data
        helper.UR = response.data.Data[0].Roalval
        helper.setUserinfo(response.data.Data[0])

        if ( this.remomber ) {
          var t = JSON.stringify({"email":this.model.email,"password":this.model.password})
          localStorage.setItem("UUID",btoa(t))
        } else {
          localStorage.removeItem("UUID")
        }

        window.location.href = '#/dashboard';
      } else {
         this.$swal.fire({
  icon: 'error',
  title: 'Oops...',
  text: response.data.Message,
})
      }
      
    });


      
    },
  },
  
};
</script>
<style></style>

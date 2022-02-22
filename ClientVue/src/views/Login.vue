<template>
   <div class="modal-dialog modal-lg" role="document">
          <div class="modal-content">
              
              <div class="modal-header">
                  <button type="button" class="close" data-dismiss="modal">
                      <span aria-hidden="true">
                          <i class="icofont icofont-close-line"></i>
                      </span>
                  </button>
                  <h4 class="modal-title" id="myModalLabel">
                      Authorization 
                      <span>
                          required
                      </span>
                  </h4>
              </div>
              
              <div class="modal-body">

                  <!-- Authirize form -->
                  <div class="row auth-form">
                      <div class="col-md-4">

                          <!-- Nav -->
                          <div class="asside-nav no-bg">
                              <ul class="nav-vrt border">
                                  <li class="active">
                                      <a href="#" class="btn-material">Privacy policy</a>
                                  </li>
                                  <li>
                                      <a href="#" class="btn-material">Terms and conditions</a>
                                  </li>
                                  <li>
                                      <a href="#" class="btn-material">FAQ</a>
                                  </li>
                              </ul>
                          </div>
                      </div>

                      <div class="col-md-5 col-md-offset-1 form-fields">
                          <form>
                              <div class="form-group">
                                  <label for="exampleInputEmail1">Email address</label>
                                  <input type="email" class="form-control" id="exampleInputEmail1" placeholder="Email">
                              </div>
                              <div class="form-group">
                                  <label for="exampleInputPassword1">Password</label>
                                  <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
                              </div>
                              <div class="checkbox padding">
                                  <input type="checkbox" id="inputCheckBox">
                                  <label for="inputCheckBox">
                                      <span class="checkbox-input">
                                      <span class="off">off</span>
                                      <span class="on">on</span>
                                      </span>
                                      remember password
                                  </label>
                              </div>
                              <span class="sdw-wrap">
                                  <button type="submit" class="sdw-hover btn btn-material btn-yellow btn-lg ripple-cont">Login</button>
                                  <span class="sdw"></span>
                              </span>

                              <ul class="addon-login-btn">
                                  <li>
                                      <a href="#">register</a>
                                  </li>
                                  <li>or</li>
                                  <li>
                                      <a href="#">restore password</a>
                                  </li>
                              </ul>
                          </form>
                      </div>
                  </div>
                  <!-- / Authirize form -->
              </div>
          </div>
      </div>
</template>
<script>
export default {
  data(){
    return{
      Mail:"",
      Password:"",

    }
  },
  computed(){

  },
  created(){
      console.log("PASSING INSDE THE COND")
  },
  methods:{
      CloseModal(){
        document.getElementById("myModal").style.display = "none";
      },
  },
}
</script>

<style>

</style>
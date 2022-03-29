<template>
<div>
<nav class="navbar navbar-default">
  <div class="container">     
                <!-- Brand and toggle get grouped for better mobile display -->
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">
                <img src="images/main-brand.png" alt="" class="brand">
            </a>
        </div>

        <!-- Collect the nav links, forms, and other content for toggling -->
        <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
            
            <!-- Top panel / search / phone -->
            <div class="top-panel">
                <div v-if="!isLogin" class="btn-cols">
                    
                    <ul class="list-btn-group">
                      <li><a href="#Login" id="openmodal" data-toggle="modal" data-target="#myModal">Sign In</a></li>
                        <router-link :to="'/Register'"><li><a style="margin-left:5px;">Sign Up</a></li></router-link>
                    </ul>
                </div>
                <div v-if="isLogin" class="btn-cols">
                    
                    <ul class="list-btn-group">
                      <li><a @click="Logout">Logout</a></li>
                    </ul>
                </div>
            </div>
            <ul v-if="!isLogin" class="nav navbar-nav navbar-right info-panel">
                <li class="profile">
                    <span class="wrap">
                        
                        <!-- Image -->
                        <span class="image bg-white">
                            
                            <!-- New message badge -->
                            <span class="badge bg-blue hidden-xs hidden-sm">5</span>
                            
                            <span class="icon">
                                <i class="icofont icofont-user-alt-4 text-blue"></i>
                            </span>

                            <!--img src="images/profile/profile-img.jpg" alt=""-->
                        </span>
                        <!-- Info -->
                        <span class="info">
                            <!-- Name -->
                            <span class="name text-uppercase">Guest User</span>
                        </span>
                    </span>
                </li>
                
                <!-- Cart -->
                <li class="cart">
                    
                    <a href="#" class="cart-icon hidden-xs" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
                        
                        <span class="badge bg-blue"></span>
                        
                        <i class="icofont icofont-cart-alt"></i>
                    </a>
                    
                    <a href="#" class="visible-xs" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
                        <i class="icofont icofont-cart-alt"></i>
                        Shopping cart
                    </a>
                    
                    <ul class="dropdown-menu">
                        <li  v-for="(ExistingShows,index) in ShoppedItems" :key="index">
                            <div class="wrap">
                                <div class="caption">
                                    <span class="comp-header st-1 text-uppercase">
                                        <span>
                                            {{ExistingShows.TimeShows}}    
                                        </span>
                                        <span>
                                            {{ExistingShows.ShowHours}}
                                        </span>

                                    </span>
                                    <span class="price">
                                            <span class="text-grey-dark">$</span>
                                            {{ExistingShows.Price}} <small class="text-grey-dark">.00</small>
                                    </span>
                                </div>
                                
                                <!-- Remove btn -->
                                <span class="remove-btn bg-blue">
                                    <i class="icofont icofont-bucket"></i>
                                </span>
                            </div>
                        </li>
                        <li v-if="helper.ListofCart.length==0">
                            <div class="caption">
                                    <span class="comp-header st-1 text-uppercase" style="position:relative;left:75px;">
                                        No Items in Cart
                                    </span>
                            </div>
                        </li>        
                        <li class="more-btn sdw" v-if="helper.ListofCart.length>0">
                            <a href="card-page-step-1.html" class="btn-material btn-primary">
                                View order <i class="icofont icofont-check-circled"></i>
                            </a>
                        </li>
                                    
                                      
                    </ul>
                </li>
            </ul> 
            <UserProfile v-if="isLogin"/>
            <!-- <ul class="nav navbar-nav">
                <li class="active">
                    <a href="index.html">
                        home
                    </a>
                </li>
                <li>
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                        categories <i class="icofont icofont-curved-down"></i>
                    </a>
                    <ul class="dropdown-menu">
                        <li><a href="#">Man line</a></li>
                        <li><a href="#">Woman</a></li>
                        <li><a href="#">Jewerly</a></li>
                        <li><a href="#">Electronics</a></li>
                        <li><a href="#">Clothes</a></li>
                    </ul>
                </li>
                <li>
                    <a href="index.html" class="dropdown-toggle" data-toggle="dropdown">
                        new
                    </a>
                </li>
                <li>
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                        pages <i class="icofont icofont-curved-down"></i>
                    </a>
                    <ul class="dropdown-menu">
                        <li><a href="shop-list.html">Shop category</a></li>
                        <li><a href="shop-item.html">Shop item</a></li>
                        <li><a href="card-page-step-1.html">Shopping card. Step 1</a></li>
                        <li><a href="card-page-step-2.html">Shopping card. Step 2</a></li>
                        <li><a href="card-page-step-3.html">Shopping card. Step 3</a></li>
                        <li><a href="register-page.html">Register page</a></li>
                        <li><a href="blog-item.html">Item blog</a></li>
                    </ul>
                </li>
            </ul> -->
        
        </div><!-- /.navbar-collapse -->
  </div><!-- /.container-fluid --> 
</nav>
<div class="modal fade" id="myModal">
            <div class="modal-dialog modal-lg" role="document">
                <div class="modal-content">
                   
                    <div class="modal-header">
                        <button type="button" class="close" id="modalclose" data-dismiss="modal">
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
                                <form v-if="!OTPcreated">
                                     
                                     <div class="form-group">
                                        <label for="emailid">Email</label>
                                        <input id="emailid" name="emailid" v-model="Email" v-validate="'email|required'" :class="{'form-control': true, 'danger': errors.has('email') }" type="text" placeholder="Email">
                                        <i v-show="errors.has('emailid')" class="fa fa-warning"></i>
                                     </div>
                                     <span v-show="errors.has('emailid')" class="text-danger text-center" style="font-size:small">{{ "Email ID Must Be Valid" }} <br/></span>
                                     
                                    <div class="form-group">
                                        <label for="exampleInputPassword1">Password</label>
                                        <input type="password" name="password" id="exampleInputPassword1" v-model="Password" v-validate="'required'" :class="{'form-control': true, 'danger': errors.has('password') }" placeholder="Password">
                                        <i v-show="errors.has('password')" class="fa fa-warning"></i>
                                    </div>
                                    <span v-show="errors.has('password')" class="text-danger text-center" style="font-size:small" >{{ "Check Password Length" }}</span>
                                    
                                    
                                    <span class="sdw-wrap">
                                        <button type="button" @click="checkLogin()" class="sdw-hover btn btn-material btn-yellow btn-lg ripple-cont">Login</button>
                                        <span class="sdw"></span>
                                    </span>

                                    <ul class="addon-login-btn">
                                        <li>
                                            <router-link :to="{name:'home'}"><li><a>register</a></li></router-link>
                                            
                                        </li>
                                        <li>or</li>
                                        <li>
                                            <a href="#">restore password</a>
                                        </li>
                                    </ul>
                                </form>
                                <form v-if="OTPcreated">
                                    <div class="form-group">
                                        <label for="OTPidlogin">OTP</label>
                                        <input id="OTPidlogin" name="OTPidlogin" v-model="OTPUSER" v-validate="'required'" :class="{'form-control': true, 'danger': errors.has('OTPidlogin') }" type="text" placeholder="OTP">
                                        <i v-show="errors.has('OTPidlogin')" class="fa fa-warning"></i>
                                     </div>
                                     <span v-show="errors.has('OTPidlogin')" class="text-danger text-center" style="font-size:small">{{ "OTP is Required" }} <br/></span>
                                     <span class="sdw-wrap">
                                        <button type="button" @click="ProcessOTPuser()" class="sdw-hover btn btn-material btn-yellow btn-lg ripple-cont">Login</button>
                                        <span class="sdw"></span>
                                    </span>
                                </form>
                            </div>
                        </div>
                        <!-- / Authirize form -->
                    </div>
                </div>
            </div>
</div>
</div>
</template>

<script>
import User from "./UserProfile.vue"
export default {
  
  name:"Header",
  components:{
      UserProfile:User
  },
  data(){
    return{
      Email:"",
      Password:"",
      CheckLogout:true,
      OTPcreated:false,
      OTPUSER:"",
      UserName:"",
      isLogin:false,
    }
  },
  
  created(){
      let CheckLogin= this.helper.getuserinfo();
      if(CheckLogin){
          this.isLogin=true
      }
  },
  computed:{
      ShoppedItems(){
          let ListofItems=[]
          if(this.helper.LoginFromAddCart){
            ListofItems = this.helper.ListofCart;   
          }
          return ListofItems
      }
  },
  methods:{
    Logout(){
        this.helper.ClearlocalStorage("UI")
        this.Email=""
        this.Password=""
        this.isLogin=false
        console.log("ERRORS",this.errors.items=[])
        this.OTPUSER=""
        this.$router.push("/")
    },
    checkLogin(){

        this.$validator.validateAll().then((result)=>{
            if(result){
                this.GetLoginDetails().then((x)=>{
                    if(x.message=="Login Successfully"){
                        swal.fire({
                            icon: 'success',
                            title: 'Success',
                            text: "Login Success",
                        });
                        this.isLogin = true
                        this.helper.setUserinfo({username:x.Data.UserName,sessionid:x.Data.SessionID,email:x.Data.Email,gender:x.Data.Gender})
                        this.OTPcreated=false;
                        this.ShowDashBoard("/");

 
                    }else if(x.message=="OTP Already Sended in Mail"){
                        swal.fire({
                            icon: 'warning',
                            title: 'OTP',
                            text: "OTP Sended in Mail kindly check and Enter",
                        })
                        this.OTPcreated=true
                    }else if(x.message=="Invalid Email or Password"){
                        swal.fire({
                            icon: 'error',
                            title: 'Invalid',
                            text: "Invalid Email or Password",
                        })        
                    }
                })
            }
        })
    },
    GetLoginDetails(){
        return new Promise((resolve,reject)=>{
            let parameter = {
                "Username":this.Email,
                "Password":this.Password,
            };
            let Actionurl="users/loginclientuser"
            parameter.Password = btoa(parameter.Password)
            this.axios.post(this.helper.SERVICEURL+Actionurl, parameter).then(response => {
            if(response.data.Success){
                resolve({message:response.data.Message,Data:response.data.Data})
            }  else {
                resolve({message:response.data.Message})
            } 
            });
        })
    },
    ProcessOTPuser(){
        this.helper.ProcessOTPUSER(this.Email,this.OTPUSER).then((x)=>{
            if(x.message=="Success"){
                swal.fire({
                    icon: 'success',
                    title: 'OTP',
                    text: "OTP Verified",
                });
                this.ShowDashBoard("/");

            }else{
                swal.fire({
                    icon: 'error',
                    title: 'OTP',
                    text: "Invalid One Time Password",
                }); 
            }
        })
    },
    ShowDashBoard(path){
        console.log("Savein Addcart")
        document.getElementById("modalclose").click();
        if(!this.helper.LoginFromAddCart){
            this.$router.push(path);
        }else{
            this.helper.SaveCart();
        }
    },
  },
}
</script>

<style>

</style>
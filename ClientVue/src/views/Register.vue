<template>
    <div>
    <Header/>
    <div class="container">
            
            <!-- 
            CONTENT
            =============================================== -->
            <div class="row block none-padding-top">
               
                <div class="col-xs-12">
                    <div class="sdw-block">
                        <div class="wrap bg-white">
                            
                            <!-- Authirize form -->
                            <div class="row head-block">
                                
                                <!-- Header & nav -->
                                <div class="col-md-12">
                                    
                                    <!-- Header -->
                                    <h1 class="header text-uppercase">
                                        New user
                                        <span>
                                            registration
                                        </span>
                                    </h1>
                                </div>
                            </div>
                            <!-- Authirize form -->
                            
                            <div class="row ">
                                
                                <!-- Header & nav -->
                                <div class="col-md-3">
                                    
                                    <!-- Text -->
                                    <p class="text">
                                        Magni labore ratione maiores, laborum quaerat molestiae excepturi. Corporis, necessitatibus earum.
                                    </p>
                                    
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
                                
                                <div class="col-md-9">
                                   
                                    <div class="panel-group" id="accordion">
                                       
                                        <div class="panel panel-default">
                                            <div class="panel-heading" id="profileInfo">
                                                <h4 class="panel-title">
                                                    <a data-toggle="collapse" data-parent="#accordion" href="#bankTransrerColl">
                                                        <span class="panel-indicator"></span>
                                                        Profile info
                                                    </a>
                                                </h4>
                                            </div>
                                            <div id="bankTransrerColl" class="panel-collapse collapse in">
                                                <div class="panel-body">
                                                    <form class="form-horizontal" v-if="OtpCreated==false">
                                                        
                                                        
                                                        
                                                        <div class="form-group pd-none">
                                                            <label for="Firstlogid" class="col-sm-3 control-label text-darkness">Your First Name</label>
                                                            <div class="col-sm-8">
                                                                <input type="text" id="Firstlogid" name="Firstlogid"  v-model="FirstName" v-validate="'required|max:30'" :class="{'form-control': true, 'danger': errors.has('Firstlogid') }" placeholder="FirstName">
                                                            </div>
                                                            <div v-show="errors.has('Firstlogid')" class="row text-danger text-center" style="font-size:small" >
                                                                 {{"First Name is Required"}}
                                                            </div>
                                                        </div>


                                                        <div class="form-group pd-none">
                                                            <label for="lastNameid" class="col-sm-3 control-label text-darkness">Your Last Name</label>
                                                            <div class="col-sm-8">
                                                                <input type="text" id="lastNameid" name="lastNameid"  v-model="LastName" v-validate="'required'" :class="{'form-control': true, 'danger': errors.has('lastNameid') }" placeholder="LastName">
                                                            </div>
                                                            <div v-show="errors.has('lastNameid')" class="text-danger text-center" style="font-size:small" >{{ "Last Name Required" }}</div>
                                                        </div>
                                                        
                                                        
                                                        <div class="form-group pd-none">
                                                            <label for="emailidreg" class="col-sm-3 control-label text-darkness">Enter Your Email</label>
                                                            <div class="col-sm-8">
                                                                <input type="email" name="emailidreg" v-model="Emailreg" v-validate="'required|email'" :class="{'form-control': true, 'danger': errors.has('emailidreg') }" placeholder="Enter your Email" id="emailidreg">
                                                            </div>

                                                            <div v-show="errors.has('emailidreg')" class="text-danger text-center" style="font-size:small">{{ "Email ID Must Be Valid" }} <br/></div>
                                                        </div>
                                                        <div class="form-group pd-none">
                                                            <label for="walletaddress" class="col-sm-3 control-label text-darkness">BitCoin WalletAddress</label>
                                                            <div class="col-sm-8">
                                                                <input type="text" name="walletaddress" v-model="WalletAddress" v-validate="'required'" :class="{'form-control': true, 'danger': errors.has('walletaddress') }" placeholder="Enter BitCoin Wallet Address" id="walletaddress">
                                                            </div>

                                                            <div v-show="errors.has('walletaddress')" class="text-danger text-center" style="font-size:small">{{ "BitCoin Wallet Address Required" }} <br/></div>
                                                        </div>
                                                        <div class="form-group pd-none">
                                                            <label for="Genderid" class="col-sm-3 control-label text-darkness">Gender</label>
                                                            <div class="col-sm-8">
                                                                <select v-model="Gender" class="form-select col-sm" v-validate="'required'" name="Genderid" id="Genderid">
                                                                    <option value="" selected>Select</option>
                                                                    <option value="M">Male</option>
                                                                    <option value="F">FeMale</option>
                                                                </select>
                                                            </div>
                                                            <div v-show="errors.has('Genderid')" class="text-danger text-center" style="font-size:small">{{ "Gender required" }} <br/></div>
                                                        </div>
                                                        <div class="form-group pd-none">
                                                            <label for="passwordregid" class="col-sm-3 control-label text-darkness">Enter your password</label>
                                                            <div class="col-sm-8">
                                                                <div class="row">
                                                                    <div class="col-sm-10">
                                                                        <input  :type="PasswordHide?'password':'text'" name="passwordregid" v-model="Passwordreg" v-validate="'required'" :class="{'form-control': true, 'danger': errors.has('passwordregid') }" placeholder="Enter your Password" id="passwordregid">
                                                                    </div>
                                                                    <div class="col-md-2" style="margin-top:10px">
                                                                        <Icons style="width:50%;height:50%;" :icon="PasswordHide?'eye-slash':'eye'" @click="hideShow()" />
                                                                    </div>
                                                                    <div v-show="errors.has('passwordregid')" class="text-danger text-center" style="font-size:small">{{ "Password is required" }} <br/></div>
                                                                </div>

                                                                <div class="row">
                                                                    <Password :password="Passwordreg"/>
                                                                </div>
                                                            </div>
                                                        </div>
                                                        <br/>
                                                        <div class="form-group pd-none">
                                                            <label for="passwordregid" class="col-sm-3 control-label text-darkness">Confirm your password</label>
                                                            <div class="col-sm-8">
                                                                <div class="row">
                                                                    <div class="col-sm-10">
                                                                        <input  :type="ConfirmPasswordHide?'password':'text'" name="confirmpasswordregid" v-model="ConfirmPasswordreg" v-validate="'required|max:60'" :class="{'form-control': true, 'danger': errors.has('confirmpasswordregid') }" placeholder="Enter your Password" id="passwordregid">
                                                                    </div>
                                                                    <div class="col-md-2" style="margin-top:10px">
                                                                        <Icons style="width:50%;height:50%;" :icon="ConfirmPasswordHide?'eye-slash':'eye'" @click="ConfirmhideShow()" />
                                                                    </div>
                                                                    <div v-show="errors.has('confirmpasswordregid')" class="text-danger text-center" style="font-size:small">{{ "Password is required" }} <br/></div>
                                                                </div>
                                                                <div class="row">
                                                                    <Password :password="ConfirmPasswordreg"/>
                                                                </div>
                                                                <div class="row" v-if="ConfirmandPassnotequal">{{"Confirm Password Does not match"}}</div>
                                                            </div>
                                                        </div>
                                                        <span class="sdw-wrap" style="margin-left:30%;">
                                                            <button @click="ProcessRegister()" type="button" class="sdw-hover btn btn-material btn-yellow btn-lg ripple-cont">Submit</button>
                                                            <span class="sdw"></span>
                                                        </span>
                                                    </form>
                                                    <form class="form-horizontal" v-if="OtpCreated">
                                                        <div class="form-group pd-none">
                                                            <label for="OtpSubmit" class="col-sm-3 control-label text-darkness">One Time Password</label>
                                                            <div class="col-sm-8">
                                                                <input type="text" id="OtpSubmit" name="OtpSubmit"  v-model="OtpValue" v-validate="'required|max:30'" :class="{'form-control': true, 'danger': errors.has('OtpSubmit') }" placeholder="Enter Your OTP">
                                                            </div>
                                                            <div v-show="errors.has('OtpSubmit')" class="row text-danger text-center" style="font-size:small" >
                                                                 {{"OTP is Required"}}
                                                            </div>
                                                        </div>
                                                        <span class="sdw-wrap" style="margin-left:30%;">
                                                            <button @click="ProcessotpUser()" type="button" class="sdw-hover btn btn-material btn-yellow btn-lg ripple-cont">Submit</button>
                                                            <span class="sdw"></span>
                                                        </span>
                                                    </form>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                
            </div>
            <!-- END: CONTENT -->
      
            
        </div>
        <Footer/>
    </div>
</template>
<script>
import Footer from "../layout/Footer.vue"
import Header from "../layout/header.vue"
export default {
  name: "Register",
  components:{
      Footer:Footer,
      Header:Header,
  },
  data() {
    return {
      FirstName: "",
      OtpCreated:false,
      OtpValue:"",
      LastName:"",
      WalletAddress:"",
      Emailreg: "",
      Gender:"",
      Passwordreg: "",
      PasswordHide:true,
      ConfirmPasswordreg:"",
      ConfirmPasswordHide:true,
      ConfirmandPassnotequal:false
    };
  },
  created(){
      console.log("PASSING INSDE ");
  },
  methods:{
    hideShow(){
        this.PasswordHide=!this.PasswordHide
    },
    ConfirmhideShow(){
        this.ConfirmPasswordHide=!this.ConfirmPasswordHide
    },
    ProcessRegister(){
        console.log("PASSING INSIDE ");
        let CheckDuplicateMail = (x)=>{
            return new Promise((resolve,reject)=>{
                let parameter = {
                "Email":this.Emailreg,
            };
            let Actionurl="users/checkduplicatemail"
            this.axios.post(this.helper.SERVICEURL+Actionurl, parameter).then(response => {
            if(response.data.Success){
                resolve({message:response.data.Message})
            }  else {
               resolve({message:response.data.Message})
            }     
          });
            })
        }
        let BasicValidation=()=>{
            let check = true
            if(!(this.ConfirmPasswordreg==this.Passwordreg)){
                this.ConfirmandPassnotequal = true
                check=false
            }

            if(check){
                this.$validator.validateAll().then((result)=>{
                    if(result){
                        CheckDuplicateMail(this.Emailreg).then((x)=>{
                            if(x.message!=undefined && x.message=="Email ID Already Exists"){
                                swal.fire({
                                    icon: 'error',
                                    title: 'Email Duplicate',
                                    text: x.message,
                                })
                                console.log("PASSING INSIDE THE CHECKDUPLICATE MAIL")
                            }else if(x.message!=undefined && x.message=="Already Exists but not Completed"){
                                swal.fire({
                                    icon: 'warning',
                                    title: 'OTP MAIL',
                                    text: x.message,
                                }).then(()=>{
                                    this.OtpCreated=true
                                })
                            }else{
                                this.SaveUserdata()
                            }
                        })
                    }else{
                        swal.fire({
                            icon: 'error',
                            title: 'Oops..',
                            text: "Check Exisiting Fields ...",
                        })
                    }
                })
            }
            
        }
        BasicValidation();
    },
    SaveUserdata(){
            let parameter = {
                "Username":this.Emailreg,
                "Password":this.Passwordreg,
                "Firstname":this.FirstName,
                "Lastname":this.LastName,
                "Gender":this.Gender,
                "WalletAddress":this.WalletAddress,
            };
            let Actionurl="users/saveclientuser"
            parameter.Password = btoa(parameter.Password)
            this.axios.post(this.helper.SERVICEURL+Actionurl, parameter).then(response => {
            if(response.data.Success){
                this.OtpCreated=true
            }  else {
               swal.fire({
                    icon: 'error',
                    title: 'Oops..',
                    text: response.data.Message,
                  })
            }     
          });
    },
    ProcessotpUser(){
        this.helper.ProcessOTPUSER(this.Emailreg,this.OtpValue).then((x)=>{
            if(x.message=="Success"){
                swal.fire({
                icon: 'success',
                title: 'OTP SUCCESS',
                text: "OTP APPROVED",
                })
                this.$router.push("/");
            }else{
                swal.fire({
                    icon: 'error',
                    title: 'OTP',
                    text: x.message,
                  })
            }
        })
    }

  },
};
</script>
<style>

</style>
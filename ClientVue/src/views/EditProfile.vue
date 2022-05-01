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
                                        Edit User
                                    </h1>
                                </div>
                            </div>
                            <!-- Edit Form -->
                            
                            <div class="row ">
                                
                                <!-- Header & nav -->
                                <div class="col-md-3">
                                    <div class="Images" >
                                        <img width="85%"  height="70%">
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
                                                    <form class="form-horizontal">
                                                        <div class="form-group pd-none">
                                                            <label for="UserName" class="col-sm-3 control-label text-darkness">User Name</label>
                                                            <div class="col-sm-8">
                                                                <input type="text" id="UserName" name="UserName"  v-model="UserName" v-validate="'required|max:30'" :class="{'form-control': true, 'danger': errors.has('UserName') }" placeholder="User Name">
                                                            </div>
                                                            <div v-show="errors.has('UserName')" class="row text-danger text-center" style="font-size:small" >
                                                                 {{"User Name is Required"}}
                                                            </div>
                                                        </div>
                                                        <div class="form-group pd-none">
                                                            <label for="EmailId" class="col-sm-3 control-label text-darkness">Email ID</label>
                                                            <div class="col-sm-8">
                                                                <input type="text" id="EmailId" name="EmailId" :disabled="true" v-model="EmailID" v-validate="'required|max:30'" :class="{'form-control': true, 'danger': errors.has('EmailId') }" placeholder="Email ID">
                                                            </div>
                                                            <div v-show="errors.has('EmailId')" class="row text-danger text-center" style="font-size:small" >
                                                                 {{"Email ID is Required"}}
                                                            </div>
                                                        </div>
                                                        <div class="form-group pd-none">
                                                            <label for="WalletAddress" class="col-sm-3 control-label text-darkness">BitCoin WalletAddress</label>
                                                            <div class="col-sm-8">
                                                                <input type="text" id="WalletAddress" name="WalletAddress" v-model="WalletAddress" v-validate="'required'" :class="{'form-control': true, 'danger': errors.has('WalletAddress') }" placeholder="BitCoin WalletAddress">
                                                            </div>
                                                            <div v-show="errors.has('WalletAddress')" class="row text-danger text-center" style="font-size:small" >
                                                                 {{"BitCoin WalletAddress is Required"}}
                                                            </div>
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
                                                            <label for="passwordregid" class="col-sm-3 control-label text-darkness">Change Password <p>(Optional)</p></label>
                                                            <div class="col-sm-8">
                                                                <div class="row">
                                                                    <div class="col-sm-10">
                                                                        <input  :type="PasswordHide?'password':'text'" name="passwordregid" v-model="Password" :class="{'form-control': true, 'danger': errors.has('passwordregid') }" placeholder="Change Password" id="passwordregid">
                                                                    </div>
                                                                    <div class="col-md-2" style="margin-top:10px">
                                                                        <Icons style="width:50%;height:50%;" :icon="PasswordHide?'eye-slash':'eye'" @click="hideShow()" />
                                                                    </div>

                                                                </div>

                                                                <div class="row">
                                                                    <Password :password="Password"/>
                                                                </div>
                                                            </div>
                                                        </div>
                                                        <span class="sdw-wrap" style="margin-left:30%;">
                                                            <button @click="Validate()" type="button" class="sdw-hover btn btn-material btn-yellow btn-lg ripple-cont">Save</button>
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
  name: "EditProfile",
  components:{
      Footer:Footer,
      Header:Header,
  },
  data() {
    return {
      UserName:"",
      EmailID:"",
      Gender:"",
      WalletAddress:"",
      Password:"",
      PasswordHide:true,
    };
  },
  created(){
    console.log("passed")
      let GetUserData = this.helper.getuserinfo();
      if(GetUserData){
        this.UserName = GetUserData.username;
        this.EmailID = GetUserData.email;
        this.Gender = GetUserData.gender;
        this.WalletAddress = GetUserData.walletaddress
      }
  },
  methods:{
    hideShow(){
        this.PasswordHide=!this.PasswordHide
    },
    Validate(){
      this.$validator.validateAll().then((result)=>{
        if(result){
          this.save();
        }
      })
    },
    save(){
      let parameters={"FullName":this.UserName,"Email":this.EmailID,"Gender":this.Gender,"Password":btoa(this.Password),"WalletAddress":this.WalletAddress}
      this.helper.AjaxPostSessionID(parameters,"users/edituserprofile",this).then((x)=>{
        if(x.status=="Success"){
          swal.fire({
            icon: 'success',
            title: 'Saved',
            text: 'Saved SuccessFully',
          }).then((save)=>{
            
            let updatelocalstorage = {"username":this.UserName,"email":this.EmailID,"gender":this.Gender}
            this.helper.updateuserinfo(updatelocalstorage)
            window.location="/"
          })
        }else{
          swal.fire({
            icon: 'error',
            title: 'Failed',
            text: 'Try again after some time',
          })
          console.log("error",x.message)
        }
      })
    },
  },
};
</script>
<style scoped>
  .Images{
    position: relative;
    left: 5px;
    float: right;
    border-radius: 5px;
  }
</style>
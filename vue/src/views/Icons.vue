<template>
  <div>
    <base-header type="gradient-success" class="pb-6 pb-8 pt-5 pt-md-8">
      <!-- Card stats -->
      
    </base-header>

    <div class="container-fluid mt--9">
      <div class="row">
        <div class="col">
          <div class="card shadow">
            <div class="card-header bg-transparent">
              <h3 style="float: left;">Users</h3>
              <div class="col text-right">
                <button type="button" class="btn btn-primary" v-if="!isadduser" @click="isadduser=true" >ADD NEW USER</button>
                 <button type="button" class="btn btn-danger" v-if="isadduser" @click="isadduser=false" ><i class="far fa-times-circle"></i></button>
              </div>
            </div>
            <div class="card-body">
              <table class="table table-dark" v-if="!isadduser" >
                <thead class="thead-dark">
                <tr>
                  <th>First Name</th>
                  <th>Last Name</th>
                  <th>User Name</th>
                  <th>Gender</th>
                  <th>Roal</th>
                  <th>Status</th>
                  <th>Action</th>
                </tr>
                </thead>
                <tr v-for="n in userlist" :key="n.Username">
                  <td>{{n.Firstname}}</td>
                  <td>{{n.Lastname}}</td>
                  <td>{{n.Username}}</td>
                  <td><span v-if="n.Gender=='M'">Male</span><span v-if="n.Gender=='F'">Female</span></td>
                  <td>{{n.Roalval}}</td>
                  <td><span v-if="n.IsActive" class="">Active</span> <span v-if="!n.IsActive" class="">In-Active</span></td>
                  <td><button class="btn btn-sm btn-info" type="button" @click="edit(n)"><i class="fas fa-edit"></i></button> <button class="btn btn-sm btn-danger" type="button" @click="del(n.Username)"><i class="fas fa-trash"></i></button></td>
                </tr>
              
              </table>

              <div v-if="isadduser">
                 <form>

                     <div class="form-row">
    <div class="form-group col-md-6">
      <label for="inputEmail4">First Name</label>
      <input type="text"   v-model="u.Firstname" class="form-control" name="firstname" placeholder="Enter First Name">
    </div>
    <div class="form-group col-md-6">
      <label for="inputPassword4">Last Name</label>
      <input type="text" v-model="u.Lastname" class="form-control" id="inputPassword4" placeholder="Enter Last Name">
    </div>
  </div>

           <div class="form-row">
    <div class="form-group col-md-6">
      <label for="inputEmail4">Email id (Email Id will be Username)</label>
      <input :disabled="Mode=='UPDATE' ? '' : disabled"  v-model="u.Username" type="email" class="form-control" id="inputEmail4" placeholder="Enter Email Id">
    </div>
    <div class="form-group col-md-6">
      <label for="inputPassword4">Gender</label>
     <select v-model="u.Gender" class="form-control">
       <option value="M">Male</option>
         <option value="F">FeMale</option>
     </select>
    </div>
  </div>

 <div class="form-row">
    <div class="form-group col-md-6">
      <label for="inputEmail4">Roal</label>
     <select v-model="u.Roalval" class="form-control">
       <option value="1">Data-Entry-Operator</option>
         <option value="2">Quality-Checker</option>
         <option value="3">Super-Admin</option>
     </select>
    </div>
    <div class="form-group col-md-6">
      <label for="inputPassword4">Password</label>
     <input type="text" v-model="u.Password" class="form-control" id="Password" placeholder="Enter Password">
    </div>
  </div>

   <div class="form-row">
  
    <div class="form-group col-md-6">
      <label for="inputPassword4">Status</label>
     <select v-model="u.IsActive" class="form-control">
       <option value="true">Active</option>
         <option value="false">In-Active</option>
        
     </select>
    </div>
  </div>
 
 <button type="button" class="btn btn-default" style="float: left;" v-if="isadduser && Mode=='NEW'" @click="save(true)" >Save & add</button> 
 <button type="button" class="btn btn-success" style="float: right;" v-if="isadduser && Mode=='NEW'" @click="save(false)" >Save & close</button>
 <button type="button" class="btn btn-default" style="float: right;" v-if="isadduser && Mode=='UPDATE'" @click="save(false)" >Save & close</button>
</form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <transition name="bounce">
              <div id="maxalert" v-if="maxalert"> {{max}} <i class="far fa-times-circle" @click="maxalert=false"></i></div>
              </transition>
  </div>
</template>

<script>
import helper from "../helper.js"
export default {
  data() {
    return {
      maxalert:false,
      max:0,
      Mode:"NEW",
      userlist:[],
      isadduser:false,
      u:{Firstname:"",Lastname:"",Username:"",Gender:"M",Roalval:"2",Password:"",IsActive:"true"},
      icons: [
      ],
    };
  },
  
   created: function () {
     this.loadusers();
  },
  methods: {
    validate(){
      //eslint-disable-next-line
      const regexp = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/
        if (this.u.Firstname==""){
          return this.error("Firstname is required!")
        }
        if (this.u.Lastname==""){
          return this.error("Lastname is required!")
        }
        if (this.u.Username==""){
          return this.error("Email is required!")
        } else if (!this.u.Username.match(regexp)){
          return this.error("Email is not valid!")
        }
        if (this.u.Password==""){
          return this.error("Password is required !")
        } else if (this.u.Password.length<6){
          return this.error("Password requires 6 characters minimum !")
        }
       return true
    },
    error(text){
      var _this = this
      if (this.maxalert){
         this.maxalert = false
      }
       this.max = text
         this.maxalert = true
         setTimeout(function(){
                    _this.maxalert = false;
                }, 3000);
                return false;
    },
    edit(n){
      this.Mode="UPDATE"
      this.u = n
      if (this.u.Roalval=="Data-Entry-Operator")
        this.u.Roalval=1
      if (this.u.Roalval=="Quality-Checker")
        this.u.Roalval=2
      if (this.u.Roalval=="Super-Admin")
        this.u.Roalval=3
      
      var tp = atob(this.u.Password)
     this.u.Password = tp
      this.isadduser=true
    },
    del(username){
      const article = {"Username":username};
       
        this.axios.post(helper.SERVICEURL+"users/deleteuser", article)
          .then(response => {
            if(response.data.Success){
                this.$swal.fire({
                    icon: 'success',
                    title: 'Good Job',
                    text: response.data.Message,
                  })
                   this.loadusers()
                  
            }  else {
               this.$swal.fire({
                    icon: 'error',
                    title: 'Oops..',
                    text: response.data.Message,
                  })
            }     
          });
    },
    loadusers(){
      this.Mode="NEW"
       const article = {};
        this.axios.post(helper.SERVICEURL+"users/getusers", article)
          .then(response => {
            if(response.data.Success){
              this.userlist= response.data.Data
            }      
          });

    },
    save(flag){
     if(!this.validate())
     {
       return 
     }
console.log("savesavesavesave")
       const article = this.u;
       article.Password = btoa(article.Password)
       article.Roalval =article.Roalval+""
       var Actionurl= "users/saveuser"
       if (this.Mode=="UPDATE")
            Actionurl = "users/updateuser"

        this.axios.post(helper.SERVICEURL+Actionurl, article)
          .then(response => {
            if(response.data.Success){
                this.$swal.fire({
                    icon: 'success',
                    title: 'Good Job',
                    text: response.data.Message,
                  })
                   this.u={Firstname:"",Lastname:"",Username:"",Gender:"M",Roalval:"2",Password:"",IsActive:"true"}
                  if (!flag) { 
                    this.isadduser=flag
                    this.loadusers()
                  }
            }  else {
               this.$swal.fire({
                    icon: 'error',
                    title: 'Oops..',
                    text: response.data.Message,
                  })
            }     
          });

    },
    onCopy(el) {
      var test = document.getElementById(el);
      test.select();
      document.execCommand("copy");
    },
  },
};
</script>

<style></style>

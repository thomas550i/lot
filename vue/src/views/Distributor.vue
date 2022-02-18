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
              <h3 style="float: left;">Distributor</h3>
              <div class="col text-right">
                <button type="button" class="btn btn-primary" v-if="!isadd" @click="isadd=true" >ADD NEW DISTRIBUTOR</button>
                <button type="button" class="btn btn-danger" v-if="isadd" @click="isadd=false" ><i class="far fa-times-circle"></i></button>
              </div>
            </div>
            <div class="card-body">
              <table class="table table-dark" v-if="!isadd" >
                <thead class="thead-dark">
                <tr>
                  <th>Distributor Id</th>
                  <th>Firstname </th>
                  <th>Lastname </th>
                  <th>Address </th>
                  <th>City </th>
                  <th>Email </th>
                 <th>Status</th>
                  <th>Action</th>
                </tr>
                </thead>
                <tr v-for="n in list" :key="n.ID">
                  <td>{{n.DistributorID}}</td>
                <td>{{n.Firsrname}}</td>
                <td>{{n.Lastname}}</td>
                <td>{{n.Address}}</td>
                <td>{{n.City}}</td>
                <td>{{n.Email}}</td>

                  <td><span v-if="n.IsActive" class="">Active</span> <span v-if="!n.IsActive" class="">In-Active</span></td>
                  <td><button class="btn btn-sm btn-info" type="button" @click="edit(n)"><i class="fas fa-edit"></i></button> <button class="btn btn-sm btn-danger" type="button" @click="del(n.ID)"><i class="fas fa-trash"></i></button></td>
                </tr>
              
              </table>

              <div v-if="isadd">
                 <form>

          <div class="form-row">
    <div class="form-group col-md-6">
      <label for="inputEmail4">Distributor ID</label>
      <input type="text" :disabled="Mode=='UPDATE' ? '' : disabled"   v-model="u.DistributorID" class="form-control" name="DistributorID" placeholder="Enter Distributor ID">
      
    </div>

     <div class="form-group col-md-6">
      <label for="inputEmail4">First name</label>
      <input type="text"  v-model="u.Firsrname" class="form-control" name="firsrname" placeholder="Enter firsr name">
      
    </div>
  
  </div>

   <div class="form-row">
    <div class="form-group col-md-6">
      <label for="inputEmail4">Last name</label>
      <input type="text"  v-model="u.Lastname" class="form-control" name="lastname" placeholder="Enter last name">
      
    </div>

     <div class="form-group col-md-6">
      <label for="inputEmail4">Address</label>
      <textarea type="text"  v-model="u.Address" class="form-control" name="Address" placeholder="Enter Address"> </textarea>
      
    </div>
  
  </div>

  <div class="form-row">
    <div class="form-group col-md-6">
      <label for="inputEmail4">City</label>
      <input type="text"  v-model="u.City" class="form-control" name="lastname" placeholder="Enter City">
      
    </div>

     <div class="form-group col-md-6">
      <label for="inputEmail4">Email</label>
      <input type="text"  v-model="u.Email" class="form-control" name="Email" placeholder="Enter Email">
      
    </div>
  
  </div>

   <div class="form-row">
    <div class="form-group col-md-6">
      <label for="inputEmail4">Mobile</label>
      <input type="text" @keypress="isNumber($event)"  v-model="u.Mobile" class="form-control" name="lastname" placeholder="Enter Mobile No">
      
    </div>

     <div class="form-group col-md-6">
      <label for="inputEmail4">Joined</label>
  
     <datepicker
   :lang="'en'" 
    v-model="u.Joined" 
  />
  
  

    </div>
  
  </div>

                     <div class="form-row">
   

     <div class="form-group col-md-6">
      <label for="inputPassword4">Status</label>
      <div >
       <label class="switch " style="
    float: left;
">
          <input type="checkbox" v-model="u.IsActive"  class="success">
          <span class="slider"></span>
        </label>
        </div>
    
     <!-- <select v-model="u.IsActive" class="form-control">
       <option value="true">Active</option>
         <option value="false">In-Active</option>
        
     </select> -->
    </div>
  
  </div>


 
 <button type="button" class="btn btn-default" style="float: left;" v-if="isadd && Mode=='NEW'" @click="save(true)" >Save & add</button> 
 <button type="button" class="btn btn-success" style="float: right;" v-if="isadd && Mode=='NEW'" @click="save(false)" >Save & close</button>
 <button type="button" class="btn btn-default" style="float: right;" v-if="isadd && Mode=='UPDATE'" @click="save(false)" >Save & close</button>
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
      list:[],
      isadd:false,
      u:{Firsrname:"",Mobile:"",Lastname:"",Address:"",City:"",Email:"",DistributorID:"",IsActive:true,Joined:new Date()},
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
        if (this.u.DistributorID==""){
          return this.error("DistributorID is required!")
        }
        if (this.u.Firsrname==""){
          return this.error("Firstname is required!")
        }
        if (this.u.Lastname==""){
          return this.error("Lastname is required!")
        }
        if (this.u.Address==""){
          return this.error("Address is required!")
        } 
        if (this.u.City==""){
          return this.error("City is required !")
        }
         if (this.u.Email==""){
          return this.error("Email is required!")
        } else if (!this.u.Email.match(regexp)){
          return this.error("Email is not valid!")
        } 
        if (this.u.Mobile==""){
          return this.error("Mobile number is required!")
        } else if (this.u.Mobile.length!=10){
          return this.error("Mobile number is not valid!")
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
      isNumber (evt){
    const keysAllowed= ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'];
    const keyPressed = evt.key;
    
    if (!keysAllowed.includes(keyPressed)) {
           evt.preventDefault()
    }
},
    edit(n){
      this.Mode="UPDATE"
      this.u = n
      this.isadd=true
    },
    del(ID){
      const article = {"ID":ID};
       
        this.axios.post(helper.SERVICEURL+"distributor/delete", article)
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
        this.axios.post(helper.SERVICEURL+"distributor/get", article)
          .then(response => {
            if(response.data.Success){
              this.list= response.data.Data
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
       var Actionurl= "distributor/save"
       if (this.Mode=="UPDATE")
            Actionurl = "distributor/update"

        this.axios.post(helper.SERVICEURL+Actionurl, article)
          .then(response => {
            if(response.data.Success){
                this.$swal.fire({
                    icon: 'success',
                    title: 'Good Job',
                    text: response.data.Message,
                  })
                   this.u={Time:"",Lastname:"",Username:"",Gender:"M",Roalval:"2",Password:"",IsActive:true}
                  if (!flag) { 
                    this.isadd=flag
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

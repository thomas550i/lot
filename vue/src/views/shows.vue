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
              <h3 style="float: left;">Shows</h3>
              <div class="col text-right">
                <button type="button" class="btn btn-primary" v-if="!isadd" @click="isadd=true" >ADD NEW SHOW</button>
                <button type="button" class="btn btn-danger" v-if="isadd" @click="isadd=false" ><i class="far fa-times-circle"></i></button>
              </div>
            </div>
            <div class="card-body">
              <table class="table table-dark" v-if="!isadd" >
                <thead class="thead-dark">
                <tr>
                  <th>Show Id</th>
                  <th>Show Time</th>
                 <th>Status</th>
                  <th>Action</th>
                </tr>
                </thead>
                <tr v-for="n in list" :key="n.ID">
                  <td>{{n.ID}}</td>
                  <td>{{n.Time}}</td>
                  <td><span v-if="n.IsActive" class="">Active</span> <span v-if="!n.IsActive" class="">In-Active</span></td>
                  <td><button class="btn btn-sm btn-info" type="button" @click="edit(n)"><i class="fas fa-edit"></i></button> <button class="btn btn-sm btn-danger" type="button" @click="del(n.ID)"><i class="fas fa-trash"></i></button></td>
                </tr>
              
              </table>

              <div v-if="isadd">
                 <form>

                     <div class="form-row">
    <div class="form-group col-md-6">
      <label for="inputEmail4">Show Time</label>
      <input type="text"  v-model="u.Time" class="form-control" name="Time" placeholder="Enter Show Time">
      
    </div>

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
      u:{Time:"",Lastname:"",Username:"",Gender:"M",Roalval:"2",Password:"",IsActive:true},
      icons: [
      ],
    };
  },
   created: function () {
     this.loadusers();
  },
  methods: {
    edit(n){
      this.Mode="UPDATE"
      this.u = n
      this.isadd=true
    },
    del(ID){
      const article = {"ID":ID};
       
        this.axios.post(helper.SERVICEURL+"shows/delete", article)
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
        this.axios.post(helper.SERVICEURL+"shows/get", article)
          .then(response => {
            if(response.data.Success){
              this.list= response.data.Data
            }      
          });

    },
     validate(){
        if (this.u.Time==""){
          return this.error("Show Time is required!")
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
    save(flag){
      if(!this.validate())
     {
       return 
     }
console.log("savesavesavesave")
       const article = this.u;
       var Actionurl= "shows/save"
       if (this.Mode=="UPDATE")
            Actionurl = "shows/update"

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

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
              <h3 style="float: left;">Approved Entry</h3>
              <div class="col text-right">
                <button type="button" class="btn btn-primary" v-if="!isadd" @click="movenewentry" >ADD NEW ENTRY</button>
                <button type="button" class="btn btn-danger" v-if="isadd" @click="isadd=false" ><i class="far fa-times-circle"></i></button>
              </div>
            </div>
            <div class="card-body">
              <table class="table table-dark" v-if="!isadd" >
                <thead class="thead-dark">
                <tr>
                  <th>Distributor Id</th>
                  <th>Date </th>
                  <th>Show Time </th>
                  <th>Approved by </th>
                 <th>Status</th>
                  <th>Action</th>
                </tr>
                </thead>
                <tr v-for="n in list" :key="n.ID">
                  <td>{{n.Distributorid}}</td>
                <td>{{getdate(n.Date)}}</td>
                <td>{{n.Showid}}</td>
                <td>{{n.Approvedby}}</td>
                <td>{{n.Status}}</td>
             

               
                  <td><button class="btn btn-sm btn-info" type="button" @click="edit(n)"><i class="fas fa-eye"></i></button> </td>
                </tr>
              
              </table>

              <div v-if="isadd">
                

                       <div class="row">
    <div class="col">
       <datepicker
   :lang="'en'" 
   :disabled="true"
    
    v-model="entrydate" 
  />
    </div>
    <div class="col">
     <select v-model="Showid" :disabled="true" class="form-control" style="width: 50%;"  >
         <option v-for="n in showmaster" :key="n.Time"  :value="n.ID"  >{{n.Time}}</option>
     </select>
    </div>
    <div class="col">
     <select v-model="Distributorid" :disabled="true" class="form-control"  style="width: 70%;"  >
         <option v-for="n in Distributormaster" :key="n.Firsrname" :value="n.DistributorID"  >{{n.Firsrname}} - {{n.DistributorID}}</option>
     </select>
    </div>
  </div>

   <div class="row" style="overflow: auto;margin-top:20px">
       <table border='1' style='border-collapse:collapse'>
      <tbody>
          <tr style="text-align: center;font-weight: bold;background: aliceblue;">
                <td>A</td>
                <td>Set A</td>
                <td>B</td>
                <td>Set B</td>
                <td>C</td>
                <td>Set C</td>

                <td>AB</td>
                <td>Set AB</td>
                <td>AC</td>
                <td>Set AC</td>
                <td>BC</td>
                <td>Set BC</td>

                 <td>ABC (FULL)</td>
                <td>Set ABC (FULL)</td>
                <td>ACB (HALF)</td>
                <td>Set ACB (HALF)</td>
                <td>BOX (FULL)</td>
                <td>Set BOX (FULL)</td>
                <td>BOX (HALF)</td>
                <td>Set BOX (HALF)</td>
          </tr>
         <tr v-for="(d,idx) in dailydata" :key="d.ID">

                <td><input type="number"  max="9"  @keypress="isNumber($event)" @change="onblurev(idx,'A',9)" class="cutomeinput" v-model="d.A" /></td>
                <td><input type="number"  max="99999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.SetA" @change="onblurev(idx,'SetA',99999)"/></td>
                <td><input type="number"  max="9" @keypress="isNumber($event)" class="cutomeinput" v-model="d.B" @change="onblurev(idx,'B',9)"/></td>
                <td><input type="number"  max="99999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.SetB" @change="onblurev(idx,'SetB',99999)"/></td>
                <td><input type="number"  max="9" @keypress="isNumber($event)" class="cutomeinput" v-model="d.C" @change="onblurev(idx,'C',9)"/></td>
                <td><input type="number"  max="99999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.SetC" @change="onblurev(idx,'SetC',99999)"/></td>
                <td><input type="number"  max="99" @keypress="isNumber($event)" class="cutomeinput" v-model="d.AB" @change="onblurev(idx,'AB',99)"/></td>
                <td><input type="number"  max="99999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.SetAB" @change="onblurev(idx,'SetAB',99999)"/></td>
                <td><input type="number"  max="99" @keypress="isNumber($event)" class="cutomeinput" v-model="d.AC" @change="onblurev(idx,'AC',99)"/></td>
                <td><input type="number"  max="99999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.SetAC" @change="onblurev(idx,'SetAC',99999)"/></td>
                <td><input type="number"  max="99" @keypress="isNumber($event)" class="cutomeinput" v-model="d.BC" @change="onblurev(idx,'BC',99)"/></td>
                <td><input type="number"  max="99999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.SetBC" @change="onblurev(idx,'SetBC',99999)"/></td>
                <td><input type="number"  max="999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.ABCFULL" @change="onblurev(idx,'ABCFULL',999)"/></td>
                <td><input type="number"  max="99999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.SetABCFULL" @change="onblurev(idx,'SetABCFULL',99999)"/></td>
                <td><input type="number"  max="999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.ABCHALF" @change="onblurev(idx,'ABCHALF',999)"/></td>
                <td><input type="number"  max="99999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.SetABCHALF" @change="onblurev(idx,'SetABCHALF',99999)"/></td>
                <td><input type="number"  max="999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.BOXFULL" @change="onblurev(idx,'BOXFULL',999)"/></td>
                <td><input type="number"  max="99999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.SetBOXFULL" @change="onblurev(idx,'SetBOXFULL',99999)"/></td>
                <td><input type="number"  max="999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.BOXHALF" @change="onblurev(idx,'BOXHALF',999)"/></td>
                <td><input type="number"  max="99999" @keypress="isNumber($event)" class="cutomeinput" v-model="d.SetBOXHALF" @change="onblurev(idx,'SetBOXHALF',99999)"/></td>


                  
               
        </tr>
      </tbody>
    </table>
  </div>



              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import helper from "../helper.js"

export default {
  data() {
    return {
      Distributorid:"",
      Showid:"",
      entrydate:"",
      Mode:"NEW",
      dailydata:[],
      Distributormaster:[],
      showmaster:[],
      list:[],
      isadd:false,
      u:{Firsrname:"",Lastname:"",Address:"",City:"",Email:"",DistributorID:"",IsActive:true,Joined:new Date()},
      icons: [
      ],
    };
  },
   created: function () {
     this.loadusers();
     this.getshowmaster();
     this.getDistributormaster();
  },
  methods: {
     getshowmaster(){
           this.axios.post(helper.SERVICEURL+"shows/get", {})
          .then(response => {
            if(response.data.Success){
              this.showmaster= response.data.Data
            }      
          });
      },
       getDistributormaster(){
            this.axios.post(helper.SERVICEURL+"distributor/get", {})
          .then(response => {
            if(response.data.Success){
              this.Distributormaster= response.data.Data
            }      
          });
      },
      onblurev(idx,key,max){
      var _this =this
        if (this.dailydata[idx][key]>max){
          this.dailydata[idx][key]=max
        
         this.max = max
         this.maxalert = true
         setTimeout(function(){
                    _this.maxalert = false;
                }, 2000);
        }
    },
      isNumber (evt){
    const keysAllowed= ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'];
    const keyPressed = evt.key;
    
    if (!keysAllowed.includes(keyPressed)) {
           evt.preventDefault()
    }
},
getdate(d){
  var tdate = new Date(d)
  return tdate.getDate()+"-"+(tdate.getMonth()+1)+"-"+tdate.getFullYear()
},
movenewentry(){
  window.location = "#/dataentry"
},
    edit(n){
      this.Mode="NEW"
     this.Distributorid=n.Distributorid
     this.Showid=n.Showidint+""
     this.entrydate= new Date(n.Date)
     this.loadata()
      
    },
    loadata(){
      
      this.Mode="NEW"
      this.dailydata=[]
      if (this.Distributorid=="" || this.Distributorid===undefined){
        return false
      }
       if (this.Showid=="" || this.Showid===undefined){
        return false
      }
       const article = {"Date":this.entrydate,"Distributorid":this.Distributorid,"Showid":this.Showid+""};
        this.axios.post(helper.SERVICEURL+"dailydata/get", article)
          .then(response => {
            if(response.data.Success){
              this.setdata(response.data.Data)
            }      
          });
    },
    setdata(d){
      var self = this
      if (d.length>0) {
        d.forEach(element => {

          for (const [key, value] of Object.entries(element)) {
          if (value==0) {
            element[key] =null
          }
        }

            self.dailydata.push(element)
        });
        
      } 

        for(var i = 0; i < 50; i++){
           this.dailydata.push({
             "ID":-1,
             "Date":this.entrydate,
             "Distributorid":this.Distributorid,
             "Showid":this.Showid+"",
              A :null,
              SetA:null,
              B:null,
              SetB:null,
              C:null,
              SetC:null,
              AB :null,
              SetAB:null,
              AC:null,
              SetAC:null,
              BC:null,
              SetBC:null,
              ABCFULL:null,
              SetABCFULL:null,
              ABCHALF:null,
              SetABCHALF:null,
              BOXFULL:null,
              SetBOXFULL:null,
              BOXHALF:null,
              SetBOXHALF:null,
              Status:"User-Submitted"
           })
       }

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
       const article = {"Appliedby":helper.UserInfo.Username,"Status":"Approved"};
        this.axios.post(helper.SERVICEURL+"dailydata/dailyapproval", article)
          .then(response => {
            if(response.data.Success){
              this.list= response.data.Data
            }      
          });

    },
    save(flag){
      
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

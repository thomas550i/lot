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
              <h3 style="float: left;">Data Entry</h3>
              <transition name="bounce">
              <div id="maxalert" v-if="maxalert">Can't Enter Value More Than <strong> {{max}} </strong> <i class="far fa-times-circle" @click="maxalert=false"></i></div>
              </transition>
              <div class="col text-right">
                  <button type="button" class="btn btn-primary" v-if="isadd" @click="save(false)" :disabled="OverallStatus=='Approved'" ><i class="far fa-save"></i> Save</button>
                   <button type="button" class="btn btn-success" v-if="isadd" @click="save(true)" :disabled="OverallStatus=='Approved'" ><i class="fab fa-atlassian"></i> Save and Submit</button>
                   <p v-if="OverallStatus=='Approved'" class="text-danger" style="padding-right: 10px;">* Approved Entry Can't Be Update </p>
              </div>
            </div>
            <div class="card-body">
          
              <div v-if="isadd">
                   <div class="row">
    <div class="col">
       <datepicker
   :lang="'en'" 
    :change="loadusers"
    v-model="entrydate" 
  />
    </div>
    <div class="col">
     <select v-model="Showid" class="form-control" style="width: 50%;"  @change="loadusers">
         <option v-for="n in showmaster" :key="n.Time"  :value="n.ID"  >{{n.Time}}</option>
     </select>
    </div>
    <div class="col">
     <select v-model="Distributorid" class="form-control"  style="width: 70%;" @change="loadusers" >
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
// import CustomInput from "../components/custominput.vue"

export default {
  data() {
    return {
      maxalert:false,
      max:0,
      OverallStatus:"",
      Mode:"NEW",
      entrydate:new Date(),
      Showid:"",
      Distributorid:"",
      dailydata:[],
      showmaster:[],
      Distributormaster:[],
      isadd:true,
      u:{Firsrname:"",Lastname:"",Address:"",City:"",Email:"",DistributorID:"",IsActive:true,Joined:new Date()},
      icons: [
      ],
    };
  },
  //  components: { CustomInput },
   created: function () {
     
     this.getshowmaster();
     this.getDistributormaster();
  },
   watch : {
               entrydate:function(newAge, oldAge,) {
                 if (newAge!=oldAge){
                   this.loadusers()
                 }
               },
   },
  methods: {
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
        const keysAllowed= ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
        const keyPressed = evt.key;
        
        if (!keysAllowed.includes(keyPressed)) {
              evt.preventDefault()
        }
      },
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
      this.OverallStatus=""
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
        this.getdailystatus()
    },
    getdailystatus(){
      
      this.Mode="NEW"
      this.dailydata=[]
      if (this.Distributorid=="" || this.Distributorid===undefined){
        return false
      }
       if (this.Showid=="" || this.Showid===undefined){
        return false
      }
       const article = {"Date":this.entrydate,"Distributorid":this.Distributorid,"Showid":this.Showid+""};
        this.axios.post(helper.SERVICEURL+"dailydata/getdailystatus", article)
          .then(response => {
            if(response.data.Success){
              if(response.data.Data.length>0){
                this.OverallStatus = response.data.Data[0].Status
              }
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
    },
    save(flag){
      
console.log("savesavesavesave")
       var article = [];

       this.dailydata.forEach(element => {
          var allempty = true
          for (const [key, value] of Object.entries(element)) {
          if ((value!=0 && value!=null) && (key!='ID' && key!='Date' && key!='Showid' && key!='Distributorid' &&  key!='Status')) {
            allempty = false
            break
          }
        }
        if (!allempty){
          article.push(element)
        }

       })

       if (article.length<=0){
         this.$swal.fire({
                    icon: 'error',
                    title: 'Oops..',
                    text: "Add Something To Save",
                  })
         return false;
       }
       var Actionurl= "dailydata/save"
       if (this.Mode=="UPDATE")
            Actionurl = "dailydata/update"

        this.axios.post(helper.SERVICEURL+Actionurl, article)
          .then(response => {
            if(response.data.Success){
                   this.u={Time:"",Lastname:"",Username:"",Gender:"M",Roalval:"2",Password:"",IsActive:true}
                  if (!flag) { 
                    this.$swal.fire({
                    icon: 'success',
                    title: 'Good Job',
                    text: "Data Saved but Not Yet Submitted To QC",
                  })
                    //this.isadd=flag
                    this.newentry()
                  } else {
                    this.submitforapproval()
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
    newentry(){
       this.OverallStatus=""
      this.entrydate=new Date()
      this.Showid=""
      this.Distributorid=""
      this.dailydata=[]
    },
    submitforapproval(){
      console.log("this.helper.UserInfo==>",helper.UserInfo)
      const article = {"Date":this.entrydate,"Distributorid":this.Distributorid,"Showid":this.Showid+"","Appliedby":helper.UserInfo.Username};
       this.axios.post(helper.SERVICEURL+"dailydata/submitforapproval", article)
          .then(response => {
            if(response.data.Success){
                this.$swal.fire({
                    icon: 'success',
                    title: 'Good Job',
                    text: response.data.Message,
                  })
                  this.newentry()
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

<style>
.cutomeinput{
    width: 50px;
    height: 50px;
    border: none;
}
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

/* Firefox */
input[type=number] {
  -moz-appearance: textfield;
}
#maxalert {
  position: fixed;
    padding: 20px;
    z-index: 10000;
    color: #fff;
    background: #bf0909;
    top: 100px;
    right: 10px;
    border-radius: 10px;
}


.bounce-enter-active {
  animation: bounce-in .5s ease-out both;
}

.bounce-leave-active {
  animation: bounce-in .5s reverse ease-in both;
}

@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.25);
  }
  100% {
    transform: scale(1);
  }
}
</style>

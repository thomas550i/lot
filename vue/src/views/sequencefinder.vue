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
              <h3 style="float: left;">Sequence finder</h3>
              <div class="col text-right">
              
                <button type="button" class="btn btn-danger" v-if="isadd" @click="isadd=false" ><i class="far fa-times-circle"></i></button>
              </div>
            </div>
            <div class="card-body">
           

              <div>
                

                       <div class="row">
    <div class="col">
       <datepicker
   :lang="'en'" 
   
    
    v-model="entrydate" 
  />
    </div>
    <div class="col">
     <select v-model="Showid"  class="form-control" style="width: 50%;"  >
         <option v-for="n in showmaster" :key="n.Time"  :value="n.ID"  >{{n.Time}}</option>
     </select>
    </div>
    <div class="col">
      <button type="button" class="btn btn-primary"  @click="findnumber(false)" >Find Sequence Number</button>
    </div>
  </div>

   <div class="row" style="overflow: auto;margin-top:20px">
       <div class="col" v-if="winningnumber!=''">
                         <table class="table table-dark" >
                <thead class="thead-dark">
                    <tr>
                  <th style="font-size: 15px;color: #fff;">Set A</th>
                  <th></th>
                 <th></th>
                </tr>
                <tr>
                  <th>Number</th>
                  <th>Sales</th>
                 <th>Prize</th>
                </tr>
                </thead>
                <tr v-for="(item, key) in fulldata" :key="key">
                  <td v-if="item.Setname=='A'">{{item.Setnumber}}</td>
                  <td v-if="item.Setname=='A'">{{item.Setsales}}</td>
                  <td v-if="item.Setname=='A'">{{getprinze(item,"A")}}</td>
                </tr>
              
              </table>

       </div>
       <div class="col" v-if="winningnumber!=''">
              <table class="table table-dark" >
                <thead class="thead-dark">
                    <tr>
                  <th style="font-size: 15px;color: #fff;">Set B</th>
                  <th></th>
                 <th></th>
                </tr>
                <tr>
                  <th>Number</th>
                  <th>Sales</th>
                 <th>Prize</th>
                </tr>
                </thead>
                <tr v-for="(item, key) in fulldata" :key="key">
                  <td v-if="item.Setname=='B'">{{item.Setnumber}}</td>
                  <td v-if="item.Setname=='B'">{{item.Setsales}}</td>
                  <td v-if="item.Setname=='B'">{{getprinze(item,"B")}}</td>
                </tr>
              
              </table>
       </div>
       <div class="col">
                <div class="winn" v-if="winningnumber!=''">
                    
                    <table>
                        <tr>
                            <td><div style="font-size: 20px;">Best Sequence Number</div>
                            <div style="font-size:80px;font-weight: bold;">{{winningnumber}}</div></td>
                            <td><i style="font-size: 80px;" class="fas fa-trophy"></i></td>
                            </tr>
                    </table>
                    <div>Total Winning Prize: {{lowprice}}</div>

                </div>

                <div class="winnsecand" v-if="winningnumber!=''">
                    
                    <table>
                        <tr>
                            <td><div style="font-size: 20px;">Change Sequence Number for Result</div></td>
                            </tr>
                            <tr>
                                <td><input class="form-control" v-model="scandnumber" /></td>
                            <td><button type="button"  class="btn btn-info" @click="findsecandnumber"> Check</button></td>
                            </tr>
                    </table>
                    

                </div>
       </div>

  </div>
  <div class="row" style="margin-top:10px">
       <div class="col" v-if="winningnumber!=''">
                         <table class="table table-dark" >
                <thead class="thead-dark">
                    <tr>
                  <th style="font-size: 15px;color: #fff;">Set c</th>
                  <th></th>
                 <th></th>
                </tr>
                <tr>
                  <th>Number</th>
                  <th>Sales</th>
                 <th>Prize</th>
                </tr>
                </thead>
                <tr v-for="(item, key) in fulldata" :key="key">
                  <td v-if="item.Setname=='C'">{{item.Setnumber}}</td>
                  <td v-if="item.Setname=='C'">{{item.Setsales}}</td>
                  <td v-if="item.Setname=='C'">{{getprinze(item,"C")}}</td>
                </tr>
              
              </table>

       </div>

        <div class="col" v-if="winningnumber!=''">
                         <table class="table table-dark" >
                <thead class="thead-dark">
                    <tr>
                  <th style="font-size: 15px;color: #fff;">Set AB</th>
                  <th></th>
                 <th></th>
                </tr>
                <tr>
                  <th>Number</th>
                  <th>Sales</th>
                 <th>Prize</th>
                </tr>
                </thead>
                <tr v-for="(item, key) in fulldata" :key="key">
                  <td v-if="item.Setname=='AB'">{{item.Setnumber}}</td>
                  <td v-if="item.Setname=='AB'">{{item.Setsales}}</td>
                  <td v-if="item.Setname=='AB'">{{getprinze(item,"AB")}}</td>
                </tr>
              
              </table>

       </div>

        <div class="col" v-if="winningnumber!=''">
                         <table class="table table-dark" >
                <thead class="thead-dark">
                    <tr>
                  <th style="font-size: 15px;color: #fff;">Set BC</th>
                  <th></th>
                 <th></th>
                </tr>
                <tr>
                  <th>Number</th>
                  <th>Sales</th>
                 <th>Prize</th>
                </tr>
                </thead>
                <tr v-for="(item, key) in fulldata" :key="key">
                  <td v-if="item.Setname=='BC'">{{item.Setnumber}}</td>
                  <td v-if="item.Setname=='BC'">{{item.Setsales}}</td>
                  <td v-if="item.Setname=='BC'">{{getprinze(item,"BC")}}</td>
                </tr>
              
              </table>

       </div>
  </div>

   <div class="row" style="margin-top:10px">
       <div class="col" v-if="winningnumber!=''">
                         <table class="table table-dark" >
                <thead class="thead-dark">
                    <tr>
                  <th style="font-size: 15px;color: #fff;">Set AC</th>
                  <th></th>
                 <th></th>
                </tr>
                <tr>
                  <th>Number</th>
                  <th>Sales</th>
                 <th>Prize</th>
                </tr>
                </thead>
                <tr v-for="(item, key) in fulldata" :key="key">
                  <td v-if="item.Setname=='AC'">{{item.Setnumber}}</td>
                  <td v-if="item.Setname=='AC'">{{item.Setsales}}</td>
                  <td v-if="item.Setname=='AC'">{{getprinze(item,"AC")}}</td>
                </tr>
              
              </table>

       </div>

        <div class="col" v-if="winningnumber!=''">
                         <table class="table table-dark" >
                <thead class="thead-dark">
                    <tr>
                  <th style="font-size: 15px;color: #fff;">Set ABCFULL</th>
                  <th></th>
                 <th></th>
                </tr>
                <tr>
                  <th>Number</th>
                  <th>Sales</th>
                 <th>Prize</th>
                </tr>
                </thead>
                <tr v-for="(item, key) in fulldata" :key="key">
                  <td v-if="item.Setname=='ABCFULL'">{{item.Setnumber}}</td>
                  <td v-if="item.Setname=='ABCFULL'">{{item.Setsales}}</td>
                  <td v-if="item.Setname=='ABCFULL'">{{getprinze(item,"ABCFULL")}}</td>
                </tr>
              
              </table>

       </div>

        <div class="col" v-if="winningnumber!=''">
                         <table class="table table-dark" >
                <thead class="thead-dark">
                    <tr>
                  <th style="font-size: 15px;color: #fff;">Set ABCHALF</th>
                  <th></th>
                 <th></th>
                </tr>
                <tr>
                  <th>Number</th>
                  <th>Sales</th>
                 <th>Prize</th>
                </tr>
                </thead>
                <tr v-for="(item, key) in fulldata" :key="key">
                  <td v-if="item.Setname=='ABCHALF'">{{item.Setnumber}}</td>
                  <td v-if="item.Setname=='ABCHALF'">{{item.Setsales}}</td>
                  <td v-if="item.Setname=='ABCHALF'">{{getprinze(item,"ABCHALF")}}</td>
                </tr>
              
              </table>

       </div>
  </div>

  <div class="row" style="margin-top:10px">
       <div class="col" v-if="winningnumber!=''">
                         <table class="table table-dark" >
                <thead class="thead-dark">
                    <tr>
                  <th style="font-size: 15px;color: #fff;">Set BOXFULL</th>
                  <th></th>
                 <th></th>
                </tr>
                <tr>
                  <th>Number</th>
                  <th>Sales</th>
                 <th>Prize</th>
                </tr>
                </thead>
                <tr v-for="(item, key) in fulldata" :key="key">
                  <td v-if="item.Setname=='BOXFULL'">{{item.Setnumber}}</td>
                  <td v-if="item.Setname=='BOXFULL'">{{item.Setsales}}</td>
                  <td v-if="item.Setname=='BOXFULL'">{{getprinze(item,"BOXFULL")}}</td>
                </tr>
              
              </table>

       </div>

        <div class="col" v-if="winningnumber!=''">
                         <table class="table table-dark" >
                <thead class="thead-dark">
                    <tr>
                  <th style="font-size: 15px;color: #fff;">Set BOXHALF</th>
                  <th></th>
                 <th></th>
                </tr>
                <tr>
                  <th>Number</th>
                  <th>Sales</th>
                 <th>Prize</th>
                </tr>
                </thead>
                <tr v-for="(item, key) in fulldata" :key="key">
                  <td v-if="item.Setname=='BOXHALF'">{{item.Setnumber}}</td>
                  <td v-if="item.Setname=='BOXHALF'">{{item.Setsales}}</td>
                  <td v-if="item.Setname=='BOXHALF'">{{getprinze(item,"BOXHALF")}}</td>
                </tr>
              
              </table>

       </div>

        
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
        d1:"",
        d2:"",
        d3:"",
        scandnumber:"",
        ID:"",
        fulldata:[],
        winningnumber:"",
        lowprice:"",
      Distributorid:"",
      Showid:"",
      entrydate:new Date(),
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
      getprinze(sn,set){
          var rval =0
        //   console.log("set->",set)
        //   console.log("sn->",sn)
        //   console.log("d1->",this.d1)
          if (set=="A" && sn.Setnumber==this.d1){
                rval = sn.Setsales*100
          }
          if (set=="B" && sn.Setnumber==this.d2){
                rval = sn.Setsales*100
          }
          if (set=="C" && sn.Setnumber==this.d3){
                rval = sn.Setsales*100
          }
        
        var narray = sn.Setnumber.toString()
          if (set=="AB" && narray==(this.d1+this.d2)){
                rval = sn.Setsales*700
          }
           if (set=="BC" && narray==(this.d2+this.d3)){
                rval = sn.Setsales*700
          }
           if (set=="AC" && narray==(this.d1+this.d3)){
                rval = sn.Setsales*700
          }

        if (set=="ABCFULL" && narray==(this.d1+this.d2+this.d3)){
                rval = sn.Setsales*24000
          }
        if (set=="ABCHALF" && narray==(this.d1+this.d2+this.d3)){
                rval = sn.Setsales*12000
          }

          return rval
      },
      findsecandnumber(){
          if (this.scandnumber.length!=3){
                this.$swal.fire({
                    icon: 'error',
                    title: 'Oops..',
                    text: "Enter 3 Digit Number.",
                  })
                  return false
          } else {
              this.findnumber(true)
          }
      },
      findnumber(includecustom){
          if (!includecustom){
              this.scandnumber = ""
          }
           this.axios.post(helper.SERVICEURL+"dailydata/findnumber", {"Date":this.entrydate,"Showid":this.Showid+'',"CustomNumber":this.scandnumber})
          .then(response => {
            if(response.data.Success){
              this.fulldata= response.data.Data.RData
              this.winningnumber = response.data.Data.Winningnumber
              this.d1 = this.winningnumber[0]
              this.d2 = this.winningnumber[1]
              this.d3 = this.winningnumber[2]
              this.lowprice = response.data.Data.Lowprice
            }      
          });
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
      this.ID = n.ID
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
       const article = {"Status":"User-Submitted"};
        this.axios.post(helper.SERVICEURL+"dailydata/dailyapproval", article)
          .then(response => {
            if(response.data.Success){
              this.list= response.data.Data
            }      
          });

    },
    Approve(test){
        console.log(test)
      var flag = true


      
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
                    text: response.data.Message,
                  })
                    //this.isadd=flag
                    this.loadusers()
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
     submitforapproval(){
      console.log("this.helper.UserInfo==>",helper.UserInfo)
      const article = {"ID":this.ID,"Approvedby":helper.UserInfo.Username};
       this.axios.post(helper.SERVICEURL+"dailydata/approveall", article)
          .then(response => {
            if(response.data.Success){
                this.$swal.fire({
                    icon: 'success',
                    title: 'Good Job',
                    text: response.data.Message,
                  })
                  this.isadd=false
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
    onCopy(el) {
      var test = document.getElementById(el);
      test.select();
      document.execCommand("copy");
    },
  },
};
</script>

<style>
.winn {
    padding:20px;
    background-color:#3b5998;
    color: #fff;
    text-shadow:1px 1px 2px rgb(0 0 0 / 50%)
}

.winnsecand{
    margin-top: 10px;
    padding:20px;
    background-color:#efefef;
    color: #302828;
   
}
</style>

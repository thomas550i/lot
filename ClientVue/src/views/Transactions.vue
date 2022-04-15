<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-12">
        <table class="table">
          <thead> 
            <tr>
              <th scope="col">Date</th>
              <th scope="col">Number of Tickets</th>
              <th scope="col">TotalAmount</th>
              <th scope="col">Status</th>
              <th scope="col">Action</th>
            </tr>
          </thead>
          <tbody >
            <tr v-for="(ListInside,index1) in list" :key="index1">
              <!-- <p>{{ListInside}}</p> -->
              <td>{{ListInside.Date}}</td>
              <td>{{ListInside.NumberofTickets}}</td>
              <td>{{ListInside.TotalAmount}}</td>
              <td>{{ListInside.Status}}</td>
              <td><button type="button" @click="ShowTableFunc(ListInside.Transactid)" class="btn btn-primary">View Shows</button></td>
            </tr>
            <tr style="height:200px;">
              <td v-if="list.length==0" class="text-center" style="padding-top:75px;" colspan="4">No Transaction To Show</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name:"Transaction",
  data(){
    return{
      CallLoadTableFlag:false,
      Username:"",
      list:[
      ],
      elements:['one',' two','three'],
      ShowTable:false,
    }
  },
  created(){
    this.LoadTable().then((x)=>{
      if(x.success){
        this.list = x.data
        this.CallAgainLoadTable(x.data)
      }
    })
  },
  methods:{
    ShowTableFunc(Transactid){
      let ID = Transactid
      var GetShowDetails=()=>{
        return new Promise((resolve,reject)=>{
          this.helper.AjaxPostSessionID({"Transactid":ID},"users/gettransactionbyid",this).then((x)=>{
              console.log("x------",x);
              if(x.data.length>0){
                  resolve({"success":true,data:x.data})
              }else{
                resolve({"success":false,data:[]})
              }
          })
        });
      }
      let ConvertToString=(x)=>{
        let form=""
        x.forEach((y)=>{
          form+=`<tr><td>${y.TimeShows}</td><td>${y.ShowHours}</td><td>${y.ShowType}</td><td>${y.Price}</td><td>${y.TicketNumber}</td><td>${y.PaymentStatus}</td></tr>`
        })
        return form
      }

      GetShowDetails().then((result)=>{
        if(result.data.length>0){
          let Rows = ConvertToString(result.data);
          swal.fire({
            title: 'Show Details',
            customClass: 'my-swal-wide',
            html: `<table class="table">
                <thead>
                  <tr">
                    <th style="text-align:center">DateofShow</th>
                    <th style="text-align:center">ShowHour</th>
                    <th style="text-align:center">ShowType</th>
                    <th style="text-align:center">Price</th>
                    <th style="text-align:center">TicketNumber</th>
                    <th style="text-align:center">PaymentStatus</th>
                  </tr>
                </thead>
                <tbody>
                  `+Rows+`
                </tbody>
              </table>`,
            confirmButtonText: 'OK',
          })
        }
      })
      
    },
    LoadTable(){
      return new Promise((resolve,reject)=>{
        let GetUserInfo = this.helper.getuserinfo();
        console.log(this.helper.getuserinfo())
        if(GetUserInfo){
          this.Username=GetUserInfo.email
          this.helper.AjaxPostSessionID({"MailId":this.Username},"users/gettransactions",this).then((x)=>{
              console.log("x------",x);
              if(x.data.length>0){
                  resolve({"success":true,data:x.data})
              }else{
                resolve({"success":false,data:[]})
              }
          })
        }

      })
    },
    CallAgainLoadTable(data){
      let flagcount=false
      if(data.length>0){
        data.forEach((x)=>{
          if(x.Status=="created"||x.Status=="pending"||x.Status=="delayed"){
            flagcount=true
          }
        })
        if(flagcount){
          setTimeout(() => {
              this.LoadTable().then((x)=>{
                if(x.success){
                  this.list = x.data
                  this.CallAgainLoadTable(x.data)
                }
              })
            }, 10000);
        }
      }
    }
  },
}
</script>

<style>
  .my-swal-wide{
    width:  70%;
    height:80%;
  }
</style>
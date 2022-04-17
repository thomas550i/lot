<template>
  <div class="container-fluid">
    <div class="row" style="margin-bottom:5%;">
      <div class="col-12">
        <button type="button" @click="CallTableAgain" style="float:right;" class="btn btn-primary">Refresh</button>
      </div>
    </div>
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
              <td><button type="button" @click="ShowTableFunc(ListInside.Transactid)" class="btn btn-primary">View Tickets</button></td>
            </tr>
            <tr v-if="list.length==0" style="height:200px;">
              <td v-if="list.length==0" class="text-center" style="padding-top:75px;" colspan="4">No Transaction To Show</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div class="row">
          <div class="col-12">
            <nav aria-label="Page navigation example">
              <ul class="pagination CenterIt">
                <li class="page-item">
                  <a class="page-link" href="#" aria-label="Previous">
                    <span aria-hidden="true" @click="FilterTable(PageBegin,0)">&laquo;</span>
                  </a>
                </li>
                <li :class="ShowActive==index2?'page-link active':'page-link'" v-for="(values,index2) in Total" :key="values">
                  <a class="page-link" @click="FilterTable(index2+1,index2)" >{{index2+1}}</a>
                </li>
                <li class="page-item">
                  <a class="page-link" @click="FilterTable(PageEnd,Total-1)" href="#" aria-label="Next">
                    <span aria-hidden="true">&raquo;</span>
                  </a>
                </li>
              </ul>
            </nav>
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
      Total:0,
      PageBegin:1,
      PageEnd:0,
      Page:1,
      Begin:0,
      End:0,
      ShowActive:0,
      TempList:[],
    }
  },
  created(){
    this.LoadTable().then((x)=>{
      if(x.success){
        this.TempList = x.data
        if(x.data.length>0){
          this.Total = Math.ceil(x.data.length/10)
          if(this.Total<=0){
            this.Total=1

          }
          this.FilterTable(this.Page,0);
          console.log(this.Total)
        }
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
    CallTableAgain(){
      this.LoadTable().then((x)=>{
      if(x.success){
          this.list = x.data
          // this.CallAgainLoadTable(x.data)
        }
      })
    },
    FilterTable(Page,index){
      this.ShowActive=index
      let ListLegth = this.TempList.length;
      if(Page==1){  
        this.Begin=0
        this.End=Page*10
      }else{
        this.Begin=this.End
        this.End=this.End+10
      }
      if(this.End<=ListLegth){
        this.list = this.TempList.slice(this.Begin,this.End);
      }else{
        this.list = this.TempList.slice(this.Begin,this.TempList.length);
      }
    },
    // CallAgainLoadTable(data){
    //   let flagcount=false
    //   if(data.length>0){
    //     data.forEach((x)=>{
    //       if(x.Status=="created"||x.Status=="pending"||x.Status=="delayed"){
    //         flagcount=true
    //       }
    //     })
    //     if(flagcount){
    //       setTimeout(() => {
    //           this.LoadTable().then((x)=>{
    //             if(x.success){
    //               this.list = x.data
    //               this.CallAgainLoadTable(x.data)
    //             }
    //           })
    //         }, 10000);
    //     }
    //   }
    // }
  },
}
</script>

<style>
  .my-swal-wide{
    width:  70%;
    height:80%;
  }
  .CenterIt{
    position: relative;
    left: 40%;
  }
</style>
<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-12">
        <div class="row">
          <div class="col">
            <table class="table">
          <thead> 
            <tr>
              <th scope="col">Type of Show</th>
              <th scope="col">Show Date</th>
              <th scope="col">Show Hour</th>
              <th scope="col">Price</th>
              <th scope="col">Ticket Number</th>
              <th scope="col">Ticket Status</th>
            </tr>
          </thead>
          <tbody >
            <tr class="moveleft" v-for="(ListInside,index1) in list" :key="index1">
              <!-- <p>{{ListInside}}</p> -->
              <td>{{ListInside.ShowType}}</td>
              <td>{{ListInside.TimeShows}}</td>
              <td>{{ListInside.ShowHours}}</td>
              <td>{{ListInside.Price}}</td>
              <td>{{ListInside.TicketNumber}}</td>
              <td><div><Icons class="IconButtonClass btn-danger btn-number IconButton"  :icon="ListInside.TicketStatus?'fas fa-crown':'fas fa-crown'"></Icons></div></td>
            </tr>
            <tr style="height:200px;">
              <td v-if="list.length==0" class="text-center" style="padding-top:75px;" colspan="6">Purchase Ticket To Show</td>
            </tr>
          </tbody>
        </table>
          </div>
        </div>
        <div class="row">
          <div class="col">
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
    </div>
  </div>  
</template>

<script>
export default {
  name:"PurchaseTickets",
  data(){
    return{
      list:[],
      UserName:"",
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
    });
  },
  methods:{
    LoadTable(){
      return new Promise((resolve,reject)=>{
        let GetUserInfo = this.helper.getuserinfo();
        console.log(this.helper.getuserinfo())
        if(GetUserInfo){
          this.Username=GetUserInfo.email
          this.helper.AjaxPostSessionID({"MailId":this.Username},"users/getpurchasedtickets",this).then((x)=>{
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
  }
}
</script>

<style>
.moveleft td{
  position: relative;
  left: 15px;

}
.IconButtonClass{
    background-color: #ccc;
    border-radius: 3px;
    text-transform: uppercase;
    font-size: .75em;
    min-height: 25px;
    padding: 18px 24px 15px;
    outline: 0;
    text-decoration: none;
}
.CenterIt{
    position: relative;
    left: 40%;
}
</style>
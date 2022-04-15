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
        <!-- <div class="row">
          <div class="col">
            <nav aria-label="Page navigation example">
              <ul class="pagination justify-content-end">
                <li class="page-item disabled">
                  <a class="page-link" href="#" tabindex="-1" aria-disabled="true">Previous</a>
                </li>
                <li class="page-item"><a class="page-link" href="#">1</a></li>
                <li class="page-item"><a class="page-link" href="#">2</a></li>
                <li class="page-item"><a class="page-link" href="#">3</a></li>
                <li class="page-item">
                  <a class="page-link" href="#">Next</a>
                </li>
              </ul>
            </nav>
          </div>
        </div> -->
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
    }
  },
  created(){
    this.LoadTable().then((x)=>{
      if(x.success){
        this.list = x.data
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
</style>
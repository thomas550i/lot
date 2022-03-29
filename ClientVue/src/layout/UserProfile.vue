<template>
            <ul class="nav navbar-nav navbar-right info-panel">
                <li class="profile">
                    <span class="wrap">
                        
                        <!-- Image -->
                        <span class="image bg-white">
                            
                            <!-- New message badge -->
                            <span class="badge bg-blue hidden-xs hidden-sm">5</span>
                            
                            <span class="icon">
                                <i class="icofont icofont-user-alt-4 text-blue"></i>
                            </span>

                            <!--img src="images/profile/profile-img.jpg" alt=""-->
                        </span>
                        <!-- Info -->
                        <span class="info">
                            <!-- Name -->
                            <span class="name text-uppercase">{{UserName}}</span>
                            <router-link :to="'/edit'"><a>Edit Profile</a></router-link>
                        </span>
                    </span>
                </li>
                
                <!-- Cart -->
                <li class="cart" @click="LoadCart">
                    
                    <a href="#" class="cart-icon hidden-xs" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
                        
                        <span class="badge bg-blue">{{NumberofItems}}</span>
                        
                        <i class="icofont icofont-cart-alt"></i>
                    </a>
                    
                    <a href="#" class="visible-xs" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
                        <i class="icofont icofont-cart-alt"></i>
                        Shopping cart
                    </a>
                    
                    <ul class="dropdown-menu DropDownListClass">
                        <li  v-for="(ExistingShows,index) in ListofCart" :key="index">
                            <div class="wrap">
                                <div class="caption">
                                    <span class="comp-header st-1 text-uppercase">
                                        <span style="margin-top:2px;">
                                            {{ExistingShows.TimeShows}} Time {{ExistingShows.ShowHours}}    
                                        </span>
                                        <span style="margin-top:2px;">
                                            {{ExistingShows.ShowType}}
                                        </span>
                                        <span style="margin-top:2px;">
                                            {{ExistingShows.TicketNumber}}
                                        </span>
                                    </span>
                                    <span class="price">
                                            <span class="text-grey-dark">$</span>
                                            {{ExistingShows.Price}} <small class="text-grey-dark">.00</small>
                                    </span>
                                    <span v-if="ExistingShows.Expired">
                                        <h6 style="color:red;">{{"Show Expired"}}</h6>
                                    </span>
                                </div>
                                
                                <!-- Remove btn -->
                                <span @click="DeleteCartItem(index,ExistingShows.ID)" class="remove-btn bg-blue" style="cursor:pointer;">
                                    <i class="icofont icofont-bucket"></i>
                                </span>
                            </div>
                        </li>
                        <li v-if="ListofCart.length==0">
                            <div class="caption">
                                    <span class="comp-header st-1 text-uppercase" style="position:relative;left:75px;">
                                        No Items in Cart
                                    </span>
                            </div>
                        </li>        
                        <li class="more-btn sdw" @click="CheckOut" style="cursor:pointer;" v-if="ListofCart.length>0">
                            <a class="btn-material btn-primary">
                                    Check Out <i class="icofont icofont-check-circled"></i>
                            </a>
                            <!-- <span v-if="Spinneron" class="loader">.</span> -->
                        </li>             
                    </ul>
                </li>
            </ul>  
</template>

<script>
export default {
  name:"UserProfile",
  data(){
    return {
      UserName:"Guest User",
      SessionID:"",
      Mail:"",
      ProfileImg:"",
      ListofCart:[],
      Spinneron:false,
      NumberofItems:0,
    }
  },
  props:{
    DataItem:{
      type:Object,
    }
  },
  created(){
    let UserData = this.helper.getuserinfo();
    if(UserData){
      this.UserName = UserData.username;
      this.SessionID = UserData.sessionid;
      this.Mail = UserData.email
      this.LoadCart();
    }
  },
  
  methods:{
    edit(){
        console.log("edit")
    },
    shoppingCart(){
      console.log("SESSIONID",this.SessionID)
      let Actionurl="users/shoppingcart"
      let parameter={
        "options":"nothing"
      }
      this.helper.AjaxPostSessionID(parameter,Actionurl,this).then((x)=>{
          console.log("x --- ",x)
      })
    },
    LoadCart(){
        let UserData = this.helper.getuserinfo();
        if(UserData){
            if(UserData.email){
               this.helper.GetExisitingCart(UserData.email).then((x)=>{
                     this.ListofCart = x.CartItems
                     this.NumberofItems = x.CartItems.length
               })
            }
        }
    },
    DeleteCartItem(index,ID){
        let paramters = {"Email":this.Mail,"ID":ID}
        this.helper.AjaxPostSessionID(paramters,"users/deletecartitem").then((result)=>{
            if(result.success){
                this.ListofCart.splice(index,1)
                this.NumberofItems = this.ListofCart.length
            }
        });
    },
    CheckOut(){
        if(this.ListofCart.length>0){
            let count=0;
            this.ListofCart.forEach((x)=>{
                if(x.Expired){
                    count++
                }
            })
            if(count>0){                                                                 
                swal.fire({
                        title: 'Do You Want to Remove Expired Ticket To Proceed CheckOut?',
                        showDenyButton: true,
                        confirmButtonText: 'Yes',
                        denyButtonText: `No`,
                }).then((result) => {
                    if (result.isConfirmed) {
                        this.RemoveExpiredTicket(this.Mail).then((x)=>{
                            console.log(x);
                            if(x.success){
                                //
                            }else{
                                swal.fire('Error in Removing Tickets contact administrator', '', 'error')
                            }
                        })
                    } else if (result.isDenied) {
                        swal.fire('Remove Expired Tickets TO CheckOut', '', 'info')
                    }
                })
            }
        }
    },
    RemoveExpiredTicket(MailId){
        return new Promise((resolve,reject)=>{
            this.helper.AjaxPostSessionID({"MailId":MailId},"users/removeexpiredtickets").then((result)=>{
                this.helper.GetExisitingCart(MailId).then((x)=>{
                     this.ListofCart = x.CartItems
                     this.NumberofItems = x.CartItems.length
                     resolve({"success":true})
               })
            });
        });
    },
  },
}
</script>

<style>
.DropDownListClass{
    height: 225px;
    overflow-y: scroll;
    position: absolute;
    right: -16%;
}
/* .button {
  position: relative;
  padding: 8px 16px;
  background:#00a0ea;
  border: none;
  outline: none;
  border-radius: 2px;
  cursor: pointer;
}

.button:active {
  background: #00a0ea;
}

.button__text {
  font: bold 20px "Quicksand", san-serif;
  color: #ffffff;
  transition: all 0.2s;
}

.button--loading .button__text {
  visibility: hidden;
  opacity: 0;
}

.button--loading::after {
  content: "";
  position: absolute;
  width: 16px;
  height: 16px;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  margin: auto;
  border: 4px solid transparent;
  border-top-color: #ffffff;
  border-radius: 50%;
  animation: button-loading-spinner 1s ease infinite;
}

@keyframes button-loading-spinner {
  from {
    transform: rotate(0turn);
  }

  to {
    transform: rotate(1turn);
  }
} */
</style>
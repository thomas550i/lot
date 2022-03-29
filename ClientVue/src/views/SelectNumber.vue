<template>
  <span>
    <div class="container">
          <div class="row">
              <div class="col-xs-12">
                  
                  <ol class="breadcrumb bg-blue">
                      <li><a href="#">Home</a></li>
                      <li class="active"><a><router-link :to="'/'">Upcoming Shows</router-link></a></li>
                  </ol>
                  
              </div>
          </div>
      </div>
      <div class="container">
            <!-- 
            STEPS
            =============================================== -->
            <div class="row block none-padding-top">
                <div class="col-xs-12">
                   
                    <ul class="steps row">
                        <li class="hidden-xs col-xs-12 col-sm-4 col-md-4 col-lg-3">
                            <div class="icon number bg-blue">
                                1
                            </div>
                            <span>
                                Select 
                            </span>
                            Your Favoriate Lot
                            
                            <span class="dir-icon">
                                <i class="icofont icofont-stylish-right text-yellow"></i>
                            </span>
                        </li>

                        <li class="col-sm-4 col-md-4 col-lg-3">
                            <div class="icon number bg-blue">
                                2
                            </div>
                            <span>
                                Enter 
                            </span>
                            your Desired Number     
                        </li>
                    </ul>
                </div>
            </div>
            <!-- END: STEPS -->
            
            <!-- 
            CONTENT
            =============================================== -->
            <div class="row block none-padding-top">
                <div class="col-xs-12">
                    
                    <div class="product-list">
                        <div class="wrap bg-white">
                           
                            <!-- Header -->
                            <div class="list-header text-uppercase">

                                <div class="product" style="margin-left:5px;width:29%;">
                                    Selected SLOT
                                </div>

                                <div class="price text-center" style="margin-left:5px; width:36%">
                                    <h4>{{SlotName}}</h4>
                                </div>
                            </div><!-- / Header -->
                            
                            <!-- List body -->
                            <div class="list-body" v-if="DataFormed" style="height:120px;">
                                
                                <!-- Item -->
                                <div class="item">
                                    
                                    <!-- Checkbox -->
                                    
                                    <div class="TimeValue" :style="TimeStyle">
                                        <span>{{TimeShows}}</span>
                                        <span>{{ShowHours}}</span>
                                    </div>                                    
                                    <div class="ColInputsValue" :style="SelectType=='SingleDigit'?'left:36%;':''" v-if="SelectType=='SingleDigit'||'DoubleDigit'||'TripleDigit'">
                                        <p>
                                          </p><div class="input-group" id="Singledigit">
                                              <span  @mousedown="SingleDigitChange('minus')" class="input-group-btn">
                                                    <Icons class="btn btn-danger btn-number IconButton"  :icon="'fa-solid fa-circle-minus'"></Icons>
                                              </span>
                                              <input @blur="CheckNumber('SingleDigit')" type="text" name="quant[2]" class="form-control input-number text-center" v-model="SingleDigit" min="1" max="9">
                                              <span @mousedown="SingleDigitChange('plus')" class="input-group-btn">
                                                <Icons class="btn btn-success btn-number IconButton"  :icon="'fa-solid fa-circle-plus'"></Icons>
                                              </span>
                                              <span v-if="ErrorVal">Cannot be 0</span>
                                          </div>
                                      <p></p>
                                    </div>
                                    <div class="ColInputsValue" v-if="ShowDoubleDigit">
                                        <p>
                                          </p><div class="input-group" id="doubledigit">
                                              <span  @mousedown="DoubleDigitChange('minus')" class="input-group-btn">
                                                    <Icons class="btn btn-danger btn-number IconButton"  :icon="'fa-solid fa-circle-minus'"></Icons>
                                              </span>
                                              <input @blur="CheckNumber('DoubleDigit')" type="text" name="quant[2]" class="form-control input-number text-center" v-model="DoubleDigit" min="1" max="9">
                                              <span @mousedown="DoubleDigitChange('plus')" class="input-group-btn">
                                                <Icons class="btn btn-success btn-number IconButton"  :icon="'fa-solid fa-circle-plus'"></Icons>
                                              </span>
                                          </div>
                                      <p></p>
                                    </div>
                                    <div class="ColInputsValue" v-if="SelectType=='TripleDigit'"> 
                                      <p>
                                          </p><div class="input-group" id="Tripledigit">
                                              <span  @mousedown="TripleDigitChange('minus')" class="input-group-btn">
                                                    <Icons class="btn btn-danger btn-number IconButton"  :icon="'fa-solid fa-circle-minus'"></Icons>
                                              </span>
                                              <input @blur="CheckNumber('TripleDigit')" type="text" class="form-control input-number text-center" v-model="TripleDigit" min="1" max="9">
                                              <span @mousedown="TripleDigitChange('plus')" class="input-group-btn">
                                                <Icons class="btn btn-success btn-number IconButton"  :icon="'fa-solid fa-circle-plus'"></Icons>
                                              </span>
                                          </div>
                                      <p></p>
                                    </div>
                                    <div class="ColInputsValue"> 
                                      <p>
                                          </p><div class="input-group" :style="SelectType=='SingleDigit'?'left:72%':''">
                                              <button type="button" @click="AddInCart" class="btn btn-success">Add To Cart</button>
                                          </div>
                                      <p></p>
                                    </div>                                
                            </div><!-- / List body -->
                            </div>
                            <!-- Footer -->

                            <div class="list-footer bg-blue">

                              <router-link :to="'/'"> 
                                      <button class="btn btn-default btn-material ripple-cont">Continue Shopping</button>
                              </router-link>
                              <button @click="CheckOut" class="btn btn-success btn-material ripple-cont">CheckOut</button>
                            </div><!-- / Footer -->
                        </div>
                    </div>
                </div>
            </div>
            <!-- END: CONTENT -->
        </div>
  </span>
</template>

<script>
export default {
  data(){
    return{
      TimeShows:"",
      ShowHours:"",
      Value:"",
      SlotName:"",
      TimeStyle:"",
      ErrorVal:false,
      SingleDigit:1,
      DoubleDigit:1,
      TripleDigit:1,
      DataID:0,
      ShowDoubleDigit:false,
      DataFormed:false,
      SelectType:"SingleDigit",
    }
  },
  created(){
    console.log("Passing the selectNumber",this.helper.getuserinfo())
    let ExistingValue = this.helper.getuserinfo();
    if(ExistingValue){
      this.DataID = ExistingValue.lotDataID
      this.TimeShows = ExistingValue.lotDataTimeShow
      this.ShowHours = ExistingValue.lotDataShowHour
      this.SelectType = ExistingValue.SelectedType
      if(this.SelectType=="DoubleDigit"){
        this.TimeStyle = "right:49%;"
        this.SlotName = "Double Digit Slot"
        this.ShowDoubleDigit = true
      }else if(this.SelectType=="SingleDigit"){
        this.TimeStyle = "right:31%;"
        this.SlotName = "Single Digit Slot"
      }else{
        this.SlotName = "Triple Digit Slot"
        this.ShowDoubleDigit = true
      }
      this.DataFormed = true
    }

  },
  watch: {
  SingleDigit(newVal) {
    let re = /[^0-9]/gi;
    if(!(newVal>=1&&newVal<=9)){
      this.$set(this,'SingleDigit', newVal.replace(re, ''));
    }
  },
  DoubleDigit(newVal) {
    let re = /[^0-9]/gi;
    if(!(newVal>=1&&newVal<=9)){
      this.$set(this,'DoubleDigit', newVal.replace(re, ''));
    }
  },
  TripleDigit(newVal) {
    let re = /[^0-9]/gi;
    if(!(newVal>=1&&newVal<=9)){
      this.$set(this,'TripleDigit', newVal.replace(re, ''));
    }
  }
  },
  methods:{
    SingleDigitChange(value){
      if(value=="minus" && this.SingleDigit>=0 && this.SingleDigit<=9){
        if(this.SingleDigit!=1){
          this.SingleDigit -=1
        } 
      }else if(value=="plus" && this.SingleDigit>=0 && this.SingleDigit<=9){
        if(this.SingleDigit!=9 && this.SingleDigit<=9){
          this.SingleDigit +=1 
        }
      }else{
        if(this.SingleDigit>9){
          this.SingleDigit=9
        }
        if(this.SingleDigit<0){
          this.SingleDigit=1
        }
      }
    },
    DoubleDigitChange(value){
      if(value=="minus" && this.DoubleDigit>=0 && this.DoubleDigit<=9){
        if(this.DoubleDigit!=1){
          this.DoubleDigit -=1
        } 
      }else if(value=="plus" && this.DoubleDigit>=0 && this.DoubleDigit<=9){
        if(this.DoubleDigit!=9){
          this.DoubleDigit +=1 
        }
      }else{
        if(this.DoubleDigit>9){
          this.DoubleDigit=9
        }
        if(this.DoubleDigit<0){
          this.DoubleDigit=1
        }
      }
    },
    TripleDigitChange(value){
      if(value=="minus" && this.TripleDigit>=0 && this.TripleDigit<=9){
        if(this.TripleDigit!=1){
          this.TripleDigit -=1
        } 
      }else if(value=="plus" && this.TripleDigit>=0 && this.TripleDigit<=9){
        if(this.TripleDigit!=9){
          this.TripleDigit +=1 
        }
      }else{
        if(this.TripleDigit>9){
          this.TripleDigit=9
        }
        if(this.TripleDigit<0){
          this.TripleDigit=1
        }
      }
    },
    CheckNumber(value){
      if(value=="SingleDigit"){
        if(this.SingleDigit>9){
          this.SingleDigit=9
        }
        if(this.SingleDigit<0){
          this.SingleDigit=1
        }
      }else if(value=="DoubleDigit"){
        if(this.DoubleDigit>9){
          this.DoubleDigit=9
        }
        if(this.DoubleDigit<0){
          this.DoubleDigit=1
        }
      }else{
        if(this.TripleDigit>9){
          this.TripleDigit=9
        }
        if(this.TripleDigit<0){
          this.TripleDigit=1
        }
      }
    },
    AddInCart(){
      let GetUserinfo = this.helper.getuserinfo();
      if(GetUserinfo){
        let TicketNumber=this.GetTicketNumber();
        let PriceAmount = this.GetPriceAmount();
        let ShoppedItem = {"TimeShows":this.TimeShows,"ShowHours":this.ShowHours,"ShowType":this.SelectType,"Showid":this.DataID,"TicketNumber":TicketNumber,"Price":PriceAmount}
        if(GetUserinfo.sessionid){
          console.log("After firstsave")
          this.helper.SaveCart(ShoppedItem);
        }else{
          document.getElementById("openmodal").click();
          this.helper.MoveTemporyCart(ShoppedItem);
          this.helper.LoginFromAddCart=true
        }
      }
    },
    CheckOut(){

    },
    GetTicketNumber(){
      let TicketNumber=0
      if(this.SelectType=="SingleDigit"){
        TicketNumber = Number(String(this.SingleDigit))
      }else if(this.SelectType=="DoubleDigit"){
        TicketNumber = Number(String(this.SingleDigit)+String(this.DoubleDigit))
      }else{
        TicketNumber = Number(String(this.SingleDigit)+String(this.DoubleDigit)+String(this.TripleDigit))
      }
      return TicketNumber
    },
    GetPriceAmount(){
      let PriceAmount=0
      if(this.SelectType=="SingleDigit"){
        PriceAmount = 25
      }else if(this.SelectType=="DoubleDigit"){
        PriceAmount = 50
      }else{
        PriceAmount = 100
      }
      return PriceAmount
    }
  }

}
</script>

<style scoped>
  .ColInputsValue{
    display: inline-grid;
    position: relative;
    float: left;
    width:175px;
    left:25%;
    margin-right:2%;

  }
  .IconButton{
    width: 35px;
    height: 38px;
    padding: 0px;
  }
  .RestButton{
    display: inline-grid;
    position: relative;
    float: left;
  }
  .TimeValue{
    display: inline-grid;
    position: relative;
    right: 67%;
    text-align: center;
  }
</style>
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
                        <li class="active col-xs-12 col-sm-4 col-md-4 col-lg-3">
                            <div class="icon number bg-blue">
                                1
                            </div>
                            <span>
                                Select 
                            </span>
                            Your Favoriate Lot
                            
                            <span class="dir-icon hidden-xs">
                                <i class="icofont icofont-stylish-right text-yellow"></i>
                            </span>
                        </li>

                        <li class="hidden-xs col-sm-4 col-md-4 col-lg-3">
                            <div class="icon number bg-grey">
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

                                <div class="product">
                                    Selected SLOT
                                </div>

                                <div class="price hidden-xs hidden-sm">
                                    Type1
                                </div>

                                <div class="qnt hidden-xs hidden-sm">
                                    Type2
                                </div>

                                <div class="total hidden-xs hidden-sm">
                                    Type3
                                </div>

                                <div class="rmv hidden-xs hidden-sm">
                                    Remove
                                </div>
                            </div><!-- / Header -->
                            
                            <!-- List body -->
                            <div class="list-body">
                                
                                <!-- Item -->
                                <div class="item">
                                    
                                    <!-- Checkbox -->
                                    
                                    <div class="product">
                                        <img src="images/shop/img-01.jpg" alt="">
                                        
                                        <span class="comp-header st-8 text-uppercase">
                                            {{TimeShows}}
                                            <span>
                                                {{ShowHours}}
                                            </span>
                                            
                                        </span>
                                    </div>

                                    <div class="price hidden-xs">
                                        <button @click="SingleDigit" class="sale-badge item-badge text-uppercase bg-green" style="display:block;text-align:center;">
                                            <span style="position:relative;bottom:10px;">Single Digit</span>
                                        </button>
                                    </div>

                                    <div class="price hidden-xs">
                                        <button @click="DoubleDigit()" class="sale-badge item-badge text-uppercase" style="display:block;text-align:center;background-color:rgba(254, 77, 249, 0.71)">
                                            <span style="position:relative;bottom:10px;">Double Digit</span>
                                        </button>
                                    </div>

                                    <div class="price hidden-xs">
                                        <button @click="TripleDigit()" class="sale-badge item-badge text-uppercase bg-green" style="display:block;text-align:center;background-color:rgba(243, 177, 23, 0.71)">
                                            <span style="position:relative;bottom:10px;">Triple Digit</span>
                                        </button>
                                    </div>
                                     <router-link :to="'/'"> 
                                    <div class="rmv text-center">

                                        <button class="remove-btn">
                                          
                                            <i class="icofont icofont-close-line"></i>
                                        </button>
                                    </div>
                                    </router-link>
                                </div>
                                
                            </div><!-- / List body -->
                            
                            <!-- Footer -->
                            <div class="list-footer bg-blue">
                                
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
  name:"SelectSlot",
  data(){
    return{
      ID:0,
      TimeShows:"",
      ShowHours:"",
      TypeofLot:3,
    }
  },
  methods:{
      SingleDigit(){
          this.helper.updateuserinfo({"SelectedType":"SingleDigit"})
          this.$router.push("/SelectNumber")
      },
      DoubleDigit(){
          this.helper.updateuserinfo({"SelectedType":"DoubleDigit"})
          this.$router.push("/SelectNumber")
      },
      TripleDigit(){
          this.helper.updateuserinfo({"SelectedType":"TripleDigit"})
          this.$router.push("/SelectNumber")
      }
  },
  created(){
    console.log("PASSING INSDE THE CREATED BUYNOW")
    let GetSelectedData = this.helper.getuserinfo()
    if(GetSelectedData){
      this.TimeShows = GetSelectedData.lotDataTimeShow;
      this.ShowHours = GetSelectedData.lotDataShowHour;
      this.ID = GetSelectedData.lotDataID;
    }
  },
  
}
</script>

<style>

</style>
<template>
    <div v-if="dataformed" class="container block">
         <div class="row">
          <div class="col-xs-12">
              <div class="block-header text-uppercase">
                  <h2 class="header">Upcoming Shows</h2>
              </div>
          </div>
        </div>
        <div class="row">
                <div class="col-xs-12">
                  <div class="owl-carousel owl-default features nav-top-left owl-loaded owl-drag">
                      <div class="owl-stage-outer">
                          <div class="owl-stage" :style="'transform: translate3d('+Translate+'px, 0px, 0px); transition: all 0.25s ease 0s; width: 3000px;'">
                            <span v-for="(Items,Index) in temp" :key="Index">
                            <div v-if="Items.display" :class="Items.display?'owl-item active':'owl-item'" style="width: 270px;">
                                  <div class="shop-item hover-sdw timer" data-timer-date="2018, 2, 5, 0, 0, 0">

                                    <div class="wrap">
                                      <div class="row">
                                        <div class="comp-header st-4 text-uppercase">
                                          {{Items.ShowHour}}
                                          <br/>
                                          {{Items.TimeShows}}
                                          <!-- Badge -->
                                          <span class="sale-badge item-badge text-uppercase bg-green" style="display:block;">
                                                {{Items.ShowHour}}
                                          </span>
                                        </div>
                                      </div>

                                      <div class="row" style="margin-top:50px;">
                                        <div class="info">
                                          <router-link :to="'/SelectSlot'"><button class="icon-card" @click="savelocal(Items)" style="margin:10px;background-color:#00a0ea;color:#fff;"> Buy Now
                                            <i class="icofont icofont-cart-alt"></i>
                                        </button></router-link>
                                          
                                        </div>
                                      </div>
                                      
                                   </div>
                                </div>
                            </div>
                            </span>
                            <div  v-if="temp.length==0"><h2>No shows Available</h2></div>
                          </div>
                      </div>
                      <div class="owl-nav">
                        <div class="owl-prev" @click="prev">
                          <span class="arrow-left icofont icofont-curved-left"></span>
                        </div>
                        <div class="owl-next" @click="next">
                          <span class="arrow-right icofont icofont-curved-right"></span>
                        </div> 
                      </div>
                      <div class="owl-dots">
                        <span v-for="(dot,indexdot) in Dot" :key="indexdot">
                          <div :class="dot.active?'owl-dot active':'owl-dot'" >
                          </div>
                        </span>
                      </div>   
                </div>
                </div>
        </div> 
    </div>
    
    
</template>

<script>
export default {
  name: "UpcomingShow",
  data(){
    return{
      temp:[],
      begin:0,
      end:4,
      Dot:[
        {dot:0,active:true},
      ],
      dataformed:false,
      ActiveDot:0,
      List:[
        
      ],
      Translate:0,
    }
  },
  created(){
    console.log("change");
    this.getdailyshows().then((x)=>{
      if(x.Data!=null){
        if(x.Data.length>0){
          this.List=x.Data
          this.temp = this.List

          if(this.List.length>0){
            let count = 0,temporary=[];
            for(let i=0;i<this.List.length;i++){
              if(count==10){
                count=0
                if(temporary.length>0){
                  temporary.push({dot:i,active:false})
                }else{
                  temporary.push({dot:i,active:true})
                }
              }
              count++
            }
            if(temporary.length>0){
              temporary[0].active=true
              this.Dot = temporary
            }else{
              this.Dot[0].dot=this.List.length-1
            }
            this.end = this.Dot[0].dot
            for(let i=this.begin;i<=this.end;i++){
              this.List[i].display = true
            }
            this.dataformed=true;
          }else{

            this.dataformed = true
          }
        }else{
            this.dataformed = true
        }
      }else{
            this.dataformed = true
      }
    })

    
  },
  methods:{
    next(){
        
        if(this.begin<this.List.length && this.end<this.List.length){
          if(this.Translate>-1750 && this.temp.length>0){
          this.Translate-=250
          }else if(this.temp.length==0){
          console.log("TEMPEMPTY",this.temp,this.begin,this.end)
          this.Translate=0
          }else{
            console.log("this.Translate",this.Translate,this.temp,this.begin,this.end)
            this.Translate=0
          }
          
          if(this.end<10 && this.begin<=0){
            this.end=10
            this.begin=1
          }else{
            this.begin = this.end
            this.end+=10
          }
          this.List.forEach((x)=>{
            x.display=false
          })
          if(this.end>this.List.length){
            this.end = this.List.length
          }
          for(let i=this.begin;i<this.end;i++){
            this.List[i].display=true
          }
          this.Dot.forEach((x)=>{
            if(x.dot==this.begin){
              x.active = true
            }else{
              x.active = false
            }
          })
          let flag=false
          this.Dot.forEach((x)=>{
            if(x.active){
              flag = true
            }
          })
          if(!flag){
            this.Dot[0].active=true
          }
          this.temp = this.List.slice(this.begin,this.end)    
        }else{
          this.Translate=-250
          this.begin=0
          let end = 0
          this.List.forEach((x)=>{
            x.display=false
          })
          if(this.List.length>10){
            this.end=10
          }else{
            this.end = this.List.length
          }
          for(let i=0;i<this.end;i++){
            this.List[i].display=true

          }
          this.temp = this.List.slice(this.begin,this.end)    
        }
    },
    prev(){
      if(this.begin<this.List.length){
          if(this.Translate<0 && this.temp.length>5){
            this.Translate+=250
            }else if(this.temp.length==0){
            console.log("TEMPEMPTY",this.temp,this.begin,this.end)
            this.Translate=0
            }else{
              console.log("this.Translate",this.Translate,this.temp,this.begin,this.end)
              this.Translate=0
            }
            
            if(this.end>10){
              this.end-=10
            }else{
              this.end=10
            }
            if(this.begin>10){
              this.begin-=10
            }else{
              this.begin=0
            }
            this.List.forEach((x)=>{
              x.display=false
            })
            if(this.end>this.List.length){
              this.end = this.List.length
            }
            for(let i=this.begin;i<this.end;i++){
              this.List[i].display=true
            }
            this.Dot.forEach((x)=>{
              if(x.dot==this.begin){
                x.active = true
              }else{
                x.active = false
              }
            })
            let flag=false
            this.Dot.forEach((x)=>{
              if(x.active){
                flag = true
              }
            })
            if(!flag){
              this.Dot[0].active=true
            }
            this.temp = this.List.slice(this.begin,this.end)    
        }
    },
    getdailyshows(){
      return new Promise((resolve,reject)=>{
        let parameter = {

        };
        let Actionurl="users/getdailyshows"
        parameter.Password = btoa(parameter.Password)
        this.axios.post(this.helper.SERVICEURL+Actionurl, parameter).then(response => {
        if(response.data.Success){
            resolve({message:response.data.Message,Data:response.data.Data})
        }  else {
            resolve({message:response.data.Message})
        } 
        });
      })
    },
    savelocal(items){
      this.helper.updateuserinfo({"lotDataID":items.ID,"lotDataTimeShow":items.TimeShows,"lotDataShowHour":items.ShowHour});
    }
  },
}
</script>

<style scoped>
  .shop-item .item-badge {
    position: absolute;
    top: 30px;
    right: 20px;
}

.comp-header>span {
    display: block;
    font-weight: 500;
}
.bg-green {
    background-color: #3ab522;
}
.item-badge {
    color: #fff;
    font-size: .875em;
    font-weight: 700;
    width: 70px;
    height: 70px;
    display: block;
    border-radius: 50%;
    text-align: center;
    padding-top: 27px;
}
.comp-header {
    display: inline-block;
    font-weight: 800;
    line-height: 1.286;
    letter-spacing: .02em;
}
.text-uppercase {
    text-transform: uppercase;
}
</style>
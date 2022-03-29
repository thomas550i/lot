import { faSortAlphaDownAlt } from "@fortawesome/free-solid-svg-icons";

import Axios from "axios"
export default {
    SERVICEURL:"http://localhost:8111/",
    ListofCart:[],
    NumberofCartItems:0,
    LoginFromAddCart:false,

    setUserinfo(user){
        localStorage.setItem("UI", JSON.stringify(user));
    },
    getuserinfo(){
        var d = localStorage.getItem("UI");
        if (d===null || d==""){
            return false
        } else {
            return JSON.parse(d)
        }
    },
    updateuserinfo(parameter){
        let GetUserData = this.getuserinfo();
        if(parameter && GetUserData){
            let keys = Object.keys(parameter)
            if(keys){
                keys.forEach((x)=>{
                    GetUserData[x] = parameter[x]
                });
            }
        }else{
            GetUserData = parameter;
        }
        this.setUserinfo(GetUserData)
    },
    ProcessOTPUSER(Email,OTP){
        console.log("EMAIL",Email,"OTP",OTP);
        return new Promise((resolve,reject)=>{
            let parameter = {
                "Username":Email,
                "OTP":OTP,
            };
            let Actionurl="users/processotpuser"
            Axios.post(this.SERVICEURL+Actionurl, parameter).then(response => {
            if(response.data.Success){
                resolve({message:"Success"})
            }  else {
                resolve({message:response.data.Message})
            }     
          });
        })
    },
    ClearlocalStorage(key){
        localStorage.removeItem(key);
    },
    AjaxPostSessionID(parameters,url,self){
        let UserData=this.getuserinfo();
        let HeadersParam={}
        if(UserData){
            HeadersParam["SessionID"] = UserData.sessionid
            HeadersParam["UserName"] = UserData.username
        }
        return new Promise((resolve,reject)=>{
            Axios.post(this.SERVICEURL+url, parameters,{headers:HeadersParam}).then(response => {
            if(response.data.Success){
                resolve({message:"Success",data:response.data.Data,status:"Success"});
            }else if(response.data.Message=="SessionExpired" && !response.data.Success){
                window.swal.fire({
                    icon: 'warning',
                    title: 'Session',
                    text: "Session Expired",
                }).then((x)=>{
                    this.ClearlocalStorage("UI")
                    window.location="/"
                })
            }else {
                resolve({message:response.data.Message,data:[],status:"Failed"});
            }     
          }).catch((error)=>{
            resolve({message:error,status:"error"});
          })
        })
    },
    MoveTemporyCart(value){
        this.ListofCart.push(value);
    },
    GetExisitingCart(username){
        return new Promise((resolve,reject)=>{
            this.AjaxPostSessionID({"Usermail":username},"users/exisitingcart",this).then((result)=>{
                if(result.data.length>0){
                    this.NumberofCartItems = result.data.length;
                    this.ListofCart = result.data;
                    resolve({CartItems:result.data,TotalCartItems:result.data.length})
                }else{
                    resolve({CartItems:[],TotalCartItems:0})
                }
            });
        });
    },
    SaveCart(value){
        let GetUserinfo = this.getuserinfo()
        if(value){
            value["UserMail"] = GetUserinfo.email
            this.AjaxPostSessionID(value,"users/addinshoppingcart",this).then((result)=>{
                 window.swal.fire({
                    icon: 'success',
                    title: 'Added to Cart',
                    text: "Successfully Added to Cart",
                }).then((x)=>{
                    this.$router.push("/SelectNumber");
                })
                this.GetExisitingCart(GetUserinfo.email);
            });

        }else{
            let Parameters = this.ListofCart[0]
            Parameters["UserMail"] = GetUserinfo.email
            this.AjaxPostSessionID(Parameters,"users/addinshoppingcart",this).then((result)=>{
                console.log(result)
                 window.swal.fire({
                    icon: 'success',
                    title: 'Added to Cart',
                    text: "Successfully Added to Cart",
                }).then((x)=>{
                    this.$router.push("/SelectNumber");
                })
                this.GetExisitingCart(GetUserinfo.email);
            });
        }
        

    }
}
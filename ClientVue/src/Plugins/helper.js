import { faSortAlphaDownAlt } from "@fortawesome/free-solid-svg-icons";

import Axios from "axios"
export default {
    SERVICEURL:"http://localhost:8111/",
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
        if(parameter){
            let keys = Object.keys(parameter)
            if(keys){
                keys.forEach((x)=>{
                    GetUserData[x] = parameter[x]
                });
            }
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
                resolve({message:"Success",data:response.Data,status:"Success"});
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
    }
}
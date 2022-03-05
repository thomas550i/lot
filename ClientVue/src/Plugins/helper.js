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
    SessionExpire(){
       window.swal.fire({
            icon: 'warning',
            title: 'OTP',
            text: "Session Expired",
        }) 
        this.ClearlocalStorage("UI");
        this.$router.push("/");
    },
    ClearlocalStorage(key){
        localStorage.removeItem(key);
    },
}
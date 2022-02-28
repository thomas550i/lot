import Axios from "axios"
export default {
    SERVICEURL:"http://localhost:8111/",
    setUserinfo(user){
        user.sstime = new Date(new Date(new Date().setHours(new Date().getHours() + 2)))
        localStorage.setItem("UI", JSON.stringify(user));
    },
    checklogin(){
        var d = localStorage.getItem("UI");
        if (d===null || d==""){
           window.location.href = "http://localhost:8080/#/login"
           return false
            
        } else {
            var uinfo = JSON.parse(d)
            var time = new Date(uinfo.sstime)
            
            var current = new Date()
            console.log(time.getTime())
            console.log(current.getTime())
            if (time.getTime()<=current.getTime()){
                window.location.href = "http://localhost:8080/#/login";
                return false
            } else {
                this.UR = uinfo.Roalval
                this.UserInfo =uinfo
            }
        }
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
    UserInfo:"",
    UR:""
}
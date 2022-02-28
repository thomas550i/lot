package handlers

import (
	c "DataLotApi/class"
	orm "DataLotApi/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/lib/gomail-master"
)

var db *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type authuser struct {
		Username string
		Password string
	}
	var t authuser
	err := decoder.Decode(&t)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}
	log.Println(t.Username)

	db = orm.Db()
	type userinfo struct {
		Lastname  string
		Firstname string
		Username  string
		Gender    string
		Roalval   string
		SessionId string
	}
	repos := []userinfo{}
	rows, err := db.Query(`select 
	lastname,
	firstname,
	username,
	gender,
	roles.roalval
	 from users LEFT JOIN roles ON users.roalid = roles.roalid 
	 where username='` + t.Username + `' and password='` + t.Password + `' `)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := userinfo{}
		err = rows.Scan(
			&repo.Lastname,
			&repo.Firstname,
			&repo.Username,
			&repo.Gender,
			&repo.Roalval,
		)
		if err != nil {
			return
		}
		repos = append(repos, repo)
	}
	err = rows.Err()
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}
	if len(repos) <= 0 {
		c.ResFail(w, nil, "Authentication failed. please check your username/password")
		return
	} else {
		repos[0].SessionId = setSession(repos[0].Username)
		c.ResSuccess(w, repos, "")
		return
	}

	// out, err := json.Marshal(repos)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, string(out))
}
func setSession(u string) string {

	return string(rand.Intn(200))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	//....
	// p := c.Users{}
	// c.GetPlayload(r, w, &p)
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&p)
	// if err != nil {
	// 	c.ResFail(w, nil, "Not a valid Input")
	// 	return
	// }

	db = orm.Db()
	repos := []c.Users{}
	rows, err := db.Query(`select 
	lastname,
	firstname,
	username,
	gender,
	roles.roalval,
	Password,
	users.isactive
	 from users LEFT JOIN roles ON users.roalid = roles.roalid`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := c.Users{}
		err = rows.Scan(
			&repo.Lastname,
			&repo.Firstname,
			&repo.Username,
			&repo.Gender,
			&repo.Roalval,
			&repo.Password,
			&repo.IsActive,
		)
		if err != nil {
			return
		}
		repos = append(repos, repo)
	}
	err = rows.Err()
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}

	c.ResSuccess(w, repos, "")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Users{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	ckrows, ckerr := db.Query(`DELETE FROM users WHERE username='` + p.Username + `'`)
	if ckerr != nil {
		c.ResFail(w, nil, ckerr.Error())
		return
	} else {
		c.ResSuccess(w, nil, "Deleted Successfully")
	}
	defer ckrows.Close()
}


func SaveClientUser(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)
	Simpletoken:=""
	for i:=0;i<7;i++{
		Simpletoken+=strconv.Itoa(rand.Intn(7))
	}
	
	fmt.Println(r.FormValue("Lastname"))
	repos:=[]string{}

	p := c.ClientUsers{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	MailBody:=`
	<h4>Welcome to DataLot `+p.Firstname +""+p.Lastname +` Your Token Number <h1>`+Simpletoken+`<h1> Enjoy Datalot !<h4>`
	
	if p.Username!=""{
		if Sendmail,errprs:=SendMailRegisterUser("thomas550i@gmail.com",p.Username,"Welcome To DATALOT",MailBody,"");!Sendmail{
			fmt.Println("SMTP ERROR",errprs)
			c.ResFail(w, nil, "FAILED IN SMTP TO SEND MAIL")
			return	
		}
		rows, err := db.Query(`INSERT INTO users (firstname, lastname, username, gender, roalid, Password,IsActive,mailtoken)
		VALUES ('` + p.Firstname + `', '` + p.Lastname + `', '` + p.Username + `', '` + p.Gender + `', '`+"4"+`', '` + p.Password + `','` + "0" + `','` + Simpletoken + `');`)
		if err != nil {
			return
		}
		defer rows.Close()
	}
	c.ResSuccess(w, repos, "Saved Successfully")
}

func SendMailRegisterUser(FromMail string,ToMail string,Subject string,MailBodys string,AttachPath string)(bool,error){
	MailSendFlag:=true
	MailConnection := gomail.NewDialer("smtp-relay.sendinblue.com", 587, "thomas550i@gmail.com", "RFU1JwAKjOSEbQBy")
	m := gomail.NewMessage()
	m.SetHeader("From",FromMail)
	m.SetHeader("To", ToMail)
	m.SetHeader("Subject", Subject)
	m.SetBody("text/html", MailBodys)
	if AttachPath!=""{
	m.Attach(AttachPath)
	}
	var errors error
	if err := MailConnection.DialAndSend(m); err != nil {
	MailSendFlag=false
	errors=err
	}
	return MailSendFlag,errors
	
}

func CheckClientDuplicateMail(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)
	Msg:="No Exisiting Email"
	type Email struct{
		Email string
		IsActive string
	}
	p:=Email{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	repos:=[]Email{}
	db = orm.Db()
	rows, err := db.Query(`select username,isactive from users where username='` + p.Email + `' `)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := Email{}
		err = rows.Scan(
			&repo.Email,
			&repo.IsActive,
		)
		if err != nil {
			return
		}
		repos = append(repos, repo)
	}
	
	err = rows.Err()
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}
	
	if len(repos)> 0 {
		for _,value:=range repos{
			if value.IsActive=="false"{
				Msg="Already Exists but not Completed"
				break
			}
		}
		if Msg=="No Exisiting Email"{
			c.ResFail(w, nil, "Email ID Already Exists")
			return
		}
	}

	
	
	c.ResSuccess(w, repos, Msg)

}

func ProcessOtpUser(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)
	type CheckOTP struct{
		Username string
		OTP string
	}
	p := CheckOTP{}
	checkflag:=false
	
	repos:=[]CheckOTP{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}
	db = orm.Db()
	rows, err := db.Query(`select username,mailtoken from users where username='` + p.Username + `' and mailtoken='`+p.OTP+`'`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := CheckOTP{}
		err = rows.Scan(
			&repo.Username,
			&repo.OTP,
		)
		if err != nil {
			return
		}
		repos = append(repos, repo)
		checkflag = true
	}

	
	err = rows.Err()
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}

	if checkflag{
		rowsupdate, err2 := db.Query(`UPDATE users SET  isactive='`+ "1" + `' WHERE username = '` + p.Username + `' and  mailtoken= '` + p.OTP + `';`)
		if err2 != nil {
			c.ResFail(w, nil, err2.Error())
			return
		}
		defer rowsupdate.Close()
	}else{
		c.ResFail(w, nil, "INVALID TOKEN")
		return
	}

	c.ResSuccess(w,p,"Success")

}

func Saveuser(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Users{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	repos := []c.Users{}

	ckrows, ckerr := db.Query(`Select username from users where username='` + p.Username + `'`)
	if ckerr != nil {
		c.ResFail(w, nil, ckerr.Error())
		return
	}
	defer ckrows.Close()
	if ckrows.Next() {
		c.ResFail(w, nil, "Username already Exist!")
		return
	}
	if p.IsActive == "true" {
		p.IsActive = "1"
	} else {
		p.IsActive = "0"
	}
	rows, err := db.Query(`INSERT INTO users (firstname, lastname, username, gender, roalid, Password,IsActive)
	VALUES ('` + p.Firstname + `', '` + p.Lastname + `', '` + p.Username + `', '` + p.Gender + `', '` + p.Roalval + `', '` + p.Password + `','` + p.IsActive + `');`)
	if err != nil {
		return
	}
	defer rows.Close()

	c.ResSuccess(w, repos, "Saved Successfully")
}

func Updateuser(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Users{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	repos := []c.Users{}

	// ckrows, ckerr := db.Query(`Select username from users where username='` + p.Username + `'`)
	// if ckerr != nil {
	// 	c.ResFail(w, nil, ckerr.Error())
	// 	return
	// }
	// defer ckrows.Close()
	// if ckrows.Next() {
	// 	c.ResFail(w, nil, "Username already Exist!")
	// 	return
	// }

	if p.IsActive == "true" {
		p.IsActive = "1"
	} else {
		p.IsActive = "0"
	}
	rows, err := db.Query(`UPDATE users
	SET  firstname='` + p.Firstname + `', lastname='` + p.Lastname + `', username='` + p.Username + `', gender='` + p.Gender + `', roalid='` + p.Roalval + `', Password='` + p.Password + `',IsActive = '` + p.IsActive + `'
	WHERE username = '` + p.Username + `';`)
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}
	defer rows.Close()

	c.ResSuccess(w, repos, "Updated Successfully")
}

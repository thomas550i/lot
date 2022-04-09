package handlers

import (
	c "DataLotApi/class"
	orm "DataLotApi/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	coinbase "github.com/coinbase-commerce-go-main"

	"github.com/lib/gomail-master"
)

var db *sql.DB


func Pupdate(w http.ResponseWriter, r *http.Request) {
	 bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
   bodyString := string(bodyBytes)
	 w.Header().Set("Content-Type", "application/json")
	 f, err := os.Create("/root/go/src/logs/log.txt")
	 if err != nil {
				fmt.Println(err)
				return
		}
		l, err := f.WriteString(bodyString)
		if err != nil {
				fmt.Println(err)
				f.Close()
				return
		}
		fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
	 fmt.Fprintf(w, bodyString)
}

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








////////////////////////CLIENT APIS ///////////////////////////////////////////////



func LoginClientUser(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)
	Message:="Login Successfully"
	type LoginUser struct{
		UserName string
		FirstName string
		LastName string
		Gender string
		Email string 
		Password string
		IsActive string
		OTP string
	} 
	result:=make(map[string]interface{}) 
	p := LoginUser{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}
	
	repos:=[]LoginUser{}
	db = orm.Db()
	rows, err := db.Query(`select lastname,firstname,username,gender,mailtoken,isactive from users where username='` + p.UserName + `' and password='` + p.Password + `' `)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := LoginUser{}
		err = rows.Scan(
			&repo.LastName,
			&repo.FirstName,
			&repo.Email,
			&repo.Gender,
			&repo.OTP,
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
	if len(repos)>0{
		if repos[0].IsActive=="false"{
			result["OTP"] = repos[0].OTP
			Message="OTP Already Sended in Mail"
		}else{
			result["UserName"]=repos[0].FirstName+""+repos[0].LastName
			result["Email"] = repos[0].Email
			result["Gender"] = repos[0].Gender
			sessionid,errors:=CheckSessioninLoginTime(repos[0].Email)
			if errors!=nil{
				c.ResFail(w, sessionid, err.Error())
				return
			}
			result["SessionID"] = sessionid
		}
	}else{
		Message = "Invalid Email or Password"
	}
	
	c.ResSuccess(w,result, Message)
}

func EditUserProfile(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)
	type EditUser struct{
		FullName string
		Password string
		Email string
		Gender string
		UserName string
		FirstName string
		LastName string
	}
	p:=EditUser{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err!=nil{
		c.ResFail(w, nil, err.Error())
		return
	}
	if !c.CheckSettionid(&w, r){
		c.ResFail(w, nil, "SessionExpired")
		return
	}

	// repos:=[]EditUser{}
	// rows, err := db.Query(`select firstname,lastname from users where username='` + p.Email + `'`)
	// if err != nil {
	// 	c.ResFail(w, nil, err.Error())
	// 	return
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	repo := EditUser{}
	// 	err = rows.Scan(
	// 		&repo.FirstName,
	// 		&repo.LastName,
	// 	)
	// 	if err != nil {
	// 		c.ResFail(w, nil, err.Error())
	// 		return
	// 	}
	// 	repos = append(repos, repo)
	// }
	// ExisitingUserName:=repos[0].FirstName+repos[0].LastName
	// err = rows.Err()
	// if err != nil {
	// 	c.ResFail(w, nil, err.Error())
	// 	return
	// }

	GetUserName:=strings.Split(p.FullName," ")
	NewUserName:=GetUserName[0]
	LastName:=""
	if len(GetUserName)>1{
		LastName=strings.Join(GetUserName[1:],"")
	}
	if strings.TrimSpace(LastName)!=""{
		NewUserName+=LastName
	}
	if p.Password==""{
		rowsupdate, err2 := db.Query(`UPDATE users SET firstname='`+ GetUserName[0] + `',lastname='`+ LastName + `',gender='`+ p.Gender + `' WHERE username = '` + p.Email + `';`)
		if err2 != nil {
			c.ResFail(w, nil, err2.Error())
			return
		}
		defer rowsupdate.Close()
	}else{
		rowsupdate, err2 := db.Query(`UPDATE users SET firstname='`+ GetUserName[0] + `',lastname='`+ LastName + `',gender='`+ p.Gender + `',password='`+ p.Password + `' WHERE username = '` + p.Email + `';`)
		if err2 != nil {
			c.ResFail(w, nil, err2.Error())
		return
		}
		defer rowsupdate.Close()
	}
	c.ResSuccess(w,nil, "SUCCESS")
}

func ShoppingCart(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)
	
	type Shoppingcart struct{
		options string
	}
	fmt.Println("SESSION ID",r.Header.Get("User-Agent"))
	p := Shoppingcart{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}
	if !c.CheckSettionid(&w, r){
		c.ResFail(w, nil, "SessionExpired")
		return
	}
	c.ResSuccess(w,nil, "SUCCESS")

}
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	type shoppingcart struct{
		ID int
	}
	p:=shoppingcart{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	if !c.CheckSettionid(&w, r){
		c.ResFail(w, nil, "SessionExpired")
		return
	}

	db = orm.Db()
	ckrows, ckerr := db.Query(`DELETE FROM shoppingcart WHERE id='` + fmt.Sprint(p.ID) + `'`)
	if ckerr != nil {
		c.ResFail(w, nil, ckerr.Error())
		return
	} else {
		c.ResSuccess(w, nil, "Deleted Successfully")
	}
	defer ckrows.Close()
}


func AddToCart(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)
	
	type Shoppingcart struct{
		UserMail string
		Price int
		ShowHours string
		TimeShows string
		ShowType string
		TicketNumber int
		Showid int
	}
	fmt.Println("SESSION ID",r.Header.Get("User-Agent"))
	p := Shoppingcart{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}
	if !c.CheckSettionid(&w, r){
		c.ResFail(w, nil, "SessionExpired")
		return
	}

	rows, err := db.Query(`INSERT INTO shoppingcart (username, showtype, showid, price,paymentstatus,ticketnumber,showhours,timeshows,expired)
		VALUES ('` + p.UserMail + `', '` + p.ShowType + `', '` + fmt.Sprint(p.Showid) + `','` + fmt.Sprint(p.Price) + `','` + fmt.Sprint("created") + `', '` + fmt.Sprint(p.TicketNumber) + `', '` + fmt.Sprint(p.ShowHours) + `', '` + fmt.Sprint(p.TimeShows) + `','` + fmt.Sprint(false) + `');`)
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return 
	}
	defer rows.Close()

	c.ResSuccess(w,nil, "SUCCESS")

}

func RemoveExpiredTickets(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)
	type shoppingcart struct{
		MailId string
	}
	p:=shoppingcart{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	if !c.CheckSettionid(&w, r){
		c.ResFail(w, nil, "SessionExpired")
		return
	}

	db = orm.Db()
	ckrows, ckerr := db.Query(`DELETE FROM shoppingcart WHERE username='` + fmt.Sprint(p.MailId) + `' AND expired='true' AND ordertoken is null`)
	if ckerr != nil {
		c.ResFail(w, nil, ckerr.Error())
		return
	}
	defer ckrows.Close()
	
	c.ResSuccess(w,nil,"SUCCESS")
}

func GetExisitingCart(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)
	
	type Shoppingcart struct{
		UserMail string
		Username string
		ShowType string
		Showid int
		TicketNumber int
		Price int
		ShowHours string
		TimeShows string
		ID int
		Expired bool
	}
	fmt.Println("SESSION ID",r.Header.Get("User-Agent"))
	p := Shoppingcart{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}
	if !c.CheckSettionid(&w, r){
		c.ResFail(w, nil, "SessionExpired")
		return
	}


	repos:=[]Shoppingcart{}
	db = orm.Db()
	rows, err := db.Query(`select id,username,showtype,showid,price,ticketnumber,showhours,timeshows from shoppingcart where username='` + p.UserMail + `' and ordertoken is null`)
	if err != nil {
		c.ResFail(w, err, "Error")
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := Shoppingcart{}
		err = rows.Scan(
			&repo.ID,
			&repo.Username,
			&repo.ShowType,
			&repo.Showid,
			&repo.Price,
			&repo.TicketNumber,
			&repo.ShowHours,
			&repo.TimeShows,

		)
		if err != nil {
			c.ResFail(w, err, "Error")
			return
		}
		repos = append(repos, repo)
	}
	
	err = rows.Err()
	if err != nil {
		c.ResFail(w, err, "Error")
		return
	}
	for index,rows:=range repos{
		if !ReturnValidDailyShows(rows.Showid){
			repos[index].Expired=true
		}
	}

	c.ResSuccess(w,repos,"SUCCESS")
}


func ReturnValidDailyShows(Showid int)bool{
	Valid:=true
	time.LoadLocation("Asia/Calcutta")
	PresentTime:=time.Now().Add(time.Minute * 30)
	FormattedTime:=strings.Split(PresentTime.Format("Mon Jan 02 2006 15:02:02 GMT+0530 (India Standard Time)"),"GMT")
	StringTime:=FormattedTime[0]
	type Shows struct{
		ID int
		Time  string
		IsActive bool
		IsspecialEvent bool
		bannerimg string
		date time.Time
	}
	repos:=[]Shows{}
	db := orm.Db()
	rows, err := db.Query(`select id,Time,IsActive,IsspecialEvent,bannerimg,date from shows where date>='` + StringTime + `' AND id='`+fmt.Sprint(Showid)+`'`)
	if err != nil {
		fmt.Println("err",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := Shows{}
		err = rows.Scan(
			&repo.ID,
			&repo.Time,
			&repo.IsActive,
			&repo.IsspecialEvent,
			&repo.bannerimg,
			&repo.date,
		)
		if err != nil {
			fmt.Println(err)
			fmt.Println()
		}
		repos = append(repos, repo)
	}
	err = rows.Err()
	if err != nil {
			fmt.Println(err)
	}
	if len(repos)==0{
		Valid = false
		rowsupdate, err2 := db.Query(`UPDATE shoppingcart SET expired='`+ "true" +`'WHERE showid = '` + fmt.Sprint(Showid) + `';`)
		if err2 != nil {
			
		}
		defer rowsupdate.Close()		
	}
	return Valid
}


func CheckSessioninLoginTime(EmailID string)(int,error){
	type SessionMaster struct{
		UserName string
		SessionId int
	}
	var SessionId int

	repos:=[]SessionMaster{}
	db = orm.Db()
	rows, err := db.Query(`select sessionid from sessionmaster where email='` + EmailID + `'`)
	if err != nil {
		return SessionId,err
	}
	defer rows.Close()
	for rows.Next() {
		repo := SessionMaster{}
		err = rows.Scan(
			&repo.SessionId,
		)
		if err != nil {
			return SessionId,err
		}
		repos = append(repos, repo)
	}
	
	err = rows.Err()
	if err != nil {
		return SessionId,err

	}



	Layout:="2006-01-02 15:04:05"
	PresentTime:=time.Now()
	ExpireTime:=time.Now().Add(time.Minute * 10)


	if len(repos)==0{
		StringSessionID:=""
		for i:=0;i<9;i++{
			StringSessionID+=strconv.Itoa(rand.Intn(10))
		}
		
		
		
		fmt.Println("PRESENT TIME ",PresentTime,PresentTime.Format(Layout))
		fmt.Println("ExpireTime TIME ",ExpireTime,ExpireTime.Format(Layout))
		
		
		rows, err := db.Query(`INSERT INTO sessionmaster (email, sessionid, presenttime, expiretime)
		VALUES ('` + EmailID + `', '` + StringSessionID + `', '` + PresentTime.Format(Layout) + `', '` + ExpireTime.Format(Layout) + `');`)
		if err != nil {
			return SessionId,err
		}
		defer rows.Close()
		SessionId,_= strconv.Atoi(StringSessionID)
	}else{
		fmt.Println("PRESENT UPDATE TIME ",PresentTime,PresentTime.Format(Layout))
		fmt.Println("ExpireTime UPDATE TIME ",ExpireTime,ExpireTime.Format(Layout))


		rowsupdate, err2 := db.Query(`UPDATE sessionmaster SET presenttime='`+ PresentTime.Format(Layout) + `',expiretime='`+ ExpireTime.Format(Layout) + ` ' WHERE email = '` + EmailID + `';`)
		if err2 != nil {
			
			return SessionId,err
		}
		defer rowsupdate.Close()
		SessionId = repos[0].SessionId
	}



	return SessionId,err

}


func ProceedToCheckOut(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)	
	type Proceed struct{
		MailId string
	}
	p:=Proceed{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}
	if !c.CheckSettionid(&w, r){
		c.ResFail(w, nil, "SessionExpired")
		return
	}
	code:=""
	type ShowTickets struct{
		ID int
		Username string
		ShowId int
		Price int
		PaymentStatus string
		ExpiredStatus bool
	}
	repos:=[]ShowTickets{}
	db := orm.Db()
	rows, err := db.Query(`select id,username,showid,price,paymentstatus,expired from shoppingcart where username='` + p.MailId + `' AND ordertoken IS NULL`)
	if err != nil {
		fmt.Println("err",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := ShowTickets{}
		err = rows.Scan(
			&repo.ID,
			&repo.Username,
			&repo.ShowId,
			&repo.Price,
			&repo.PaymentStatus,
			&repo.ExpiredStatus,
		)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, repo)
	}
	err = rows.Err()
	if err != nil {
			fmt.Println(err)
	}
	if len(repos)>0{
		fmt.Print(len(repos))
		TotalAmount:=0.0
		showids:=""
		for _,value:=range repos{
			if showids!=""{
				showids+=","+fmt.Sprint(value.ID)
			}else{
				showids=fmt.Sprint(value.ID)
			}
			TotalAmount+=float64(value.Price)
		}
		client := coinbase.Client("c91e3b19-00b6-473c-82f5-d3ed1ee6815a")
		charge, err := client.Charge.Create(coinbase.ChargeParam{
			Name:        "Buy Ticket",
			Description: "Place Your Orders",
			Local_price: coinbase.Money{
				Amount:   fmt.Sprint(TotalAmount), //amount to be paid
				Currency: "USD",
			},
			Pricing_type: "fixed_price",
			Redirect_url: "http://localhost:8080/#/",   //success page
			Cancel_url:   "http://localhost:8080/#/", //cancel page
		})
		if err!=nil{
			fmt.Println("Error ----- ",err)
		}
		if value,ok:=charge.(map[string]interface{});ok{
				if Datavalue,ok:=value["data"].(map[string]interface{});ok{
					code = Datavalue["code"].(string)
				}
				PresentTime:=time.Now().Format("2006-01-02 15:04:05")
				fmt.Println(PresentTime)
				rows, err := db.Query(`INSERT INTO transactions (username, token, totalamount, responsedata, paymentstatus, showids,timeofdate)
				VALUES ('` + p.MailId + `', '` + code + `', '` + fmt.Sprint(TotalAmount) + `', '` + fmt.Sprint(charge) + `', '`+"created"+`', '` + showids + `', '` + PresentTime + `');`)
				if err != nil {
					fmt.Println("Errors ----- ",err)
					c.ResFail(w, nil, "Error in Backend")
					return
				}
				defer rows.Close()
				updateordertokennumber(code)

			}
	}
	c.ResSuccess(w,code, "SUCCESS")
}

func GetTransactionById(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)	
	type Transactbyid struct{
		Transactid int
	}
	p:=Transactbyid{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}
	if !c.CheckSettionid(&w, r){
		c.ResFail(w, nil, "SessionExpired")
		return
	}
	type TransactionByid struct{
		ID int
		Token string
		Showids string
	}
	repos:=[]TransactionByid{}
	db := orm.Db()
	rows, err := db.Query(`select token,id,showids from transactions where id='` + fmt.Sprint(p.Transactid) + `'`)
	if err != nil {
		fmt.Println("err",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := TransactionByid{}
		err = rows.Scan(
			&repo.Token,
			&repo.ID,
			&repo.Showids,
		)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, repo)
	}
	err = rows.Err()
	if err != nil {
			fmt.Println(err)
	}
	type ListofShoppingCart struct{
		ShowType string
		Price int
		PaymentStatus string
		ShowHours string
		TimeShows string
		TicketNumber string
	}
	listofCart:=[]ListofShoppingCart{}
	if len(repos)>0{
		for _,value:=range repos{	
			db := orm.Db()
			rows, err := db.Query(`select showtype,price,paymentstatus,showhours,timeshows,ticketnumber from shoppingcart where ordertoken='` + fmt.Sprint(value.Token) + `'`)
			if err != nil {
				fmt.Println("err",err)
			}
			defer rows.Close()
			for rows.Next() {
				repo := ListofShoppingCart{}
				err = rows.Scan(
					&repo.ShowType,
					&repo.Price,
					&repo.PaymentStatus,
					&repo.ShowHours,
					&repo.TimeShows,
					&repo.TicketNumber,
				)
				if err != nil {
					fmt.Println(err)
				}
				listofCart = append(listofCart, repo)
			}
			err = rows.Err()
			if err != nil {
					fmt.Println(err)
			}	
		}
	}
	c.ResSuccess(w,listofCart, "SUCCESS")
}


func GetTransactions(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)	
	type Transact struct{
		MailId string
	}
	p:=Transact{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}
	if !c.CheckSettionid(&w, r){
		c.ResFail(w, nil, "SessionExpired")
		return
	}
	type Transaction struct{
		ID int
		Username string
		TotalAmount int
		PaymentStatus string
		Showids string
		Token string
		TimeofDate string
	}
	repos:=[]Transaction{}
	db := orm.Db()
	rows, err := db.Query(`select id,username,totalamount,paymentstatus,showids,token,timeofdate from transactions where username='` + p.MailId + `'`)
	if err != nil {
		fmt.Println("err",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := Transaction{}
		err = rows.Scan(
			&repo.ID,
			&repo.Username,
			&repo.TotalAmount,
			&repo.PaymentStatus,
			&repo.Showids,
			&repo.Token,
			&repo.TimeofDate,
		)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, repo)
	}
	err = rows.Err()
	if err != nil {
			fmt.Println(err)
	}
	var Final []interface{}
	
	if len(repos)>0{
		for _,value:=range repos{
			CreateTemp:=make(map[string]interface{})
			CreateTemp["Date"] = value.TimeofDate
			CreateTemp["NumberofTickets"] = len(strings.Split(value.Showids,","))
			CreateTemp["TotalAmount"] = value.TotalAmount
			CreateTemp["Status"] = value.PaymentStatus
			CreateTemp["Transactid"] = value.ID
			Final = append(Final, CreateTemp)
		} 
	}

	c.ResSuccess(w,Final, "SUCCESS")
}


func updateordertokennumber(code string){
		rowsupdate, err2 := db.Query(`UPDATE shoppingcart SET  ordertoken='`+ code + `' WHERE ordertoken is null AND expired=false`)
		if err2 != nil {
			fmt.Println("ERROR2 ------ ",err2)
		}
		defer rowsupdate.Close()
}

func GetDailyShows(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)
	type Shows struct{
		ID int
		Time  string
		IsActive bool
		IsspecialEvent bool
		bannerimg string
		date time.Time
	}

	time.LoadLocation("Asia/Calcutta")
	PresentTime:=time.Now().Add(time.Minute *30)
	FormattedTime:=strings.Split(PresentTime.Format("Mon Jan 02 2006 15:02:02 GMT+0530 (India Standard Time)"),"GMT")
	StringTime:=FormattedTime[0]
	
	repos:=[]Shows{}
	db := orm.Db()
	rows, err := db.Query(`select id,Time,IsActive,IsspecialEvent,bannerimg,date from shows where date>='` + StringTime + `'`)
	if err != nil {
		fmt.Println("err",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := Shows{}
		err = rows.Scan(
			&repo.ID,
			&repo.Time,
			&repo.IsActive,
			&repo.IsspecialEvent,
			&repo.bannerimg,
			&repo.date,
		)
		if err != nil {
			fmt.Println(err)
			fmt.Println()
		}
		repos = append(repos, repo)
	}
	err = rows.Err()
	if err != nil {
			fmt.Println(err)
	}
	var FinalRow []interface{}
	for _,value:=range repos{
		ShowsStruct:=make(map[string]interface{})
		ShowsStruct["ShowHour"] = value.Time
		ShowsStruct["TimeShows"] = value.date.Format("02-Jan-2006")
		ShowsStruct["SpecialEvent"] = value.IsspecialEvent
		ShowsStruct["BannerImg"] = value.bannerimg
		ShowsStruct["display"] = true
		ShowsStruct["ID"] = value.ID
		FinalRow = append(FinalRow, ShowsStruct)
	}
	
	
	c.ResSuccess(w,FinalRow, "SUCCESS")
}



func SaveClientUser(w http.ResponseWriter, r *http.Request){
	c.SetupResponse(&w, r)
	Simpletoken:=""
	for i:=0;i<7;i++{
		Simpletoken+=strconv.Itoa(rand.Intn(7))
	}
	
	fmt.Println(r.FormValue("Lastname"))

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
	c.ResSuccess(w, nil, "Saved Successfully")
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

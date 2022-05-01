package class

import (
	orm "DataLotApi/db"
	"encoding/json"
	"strings"

	"fmt"
	"net/http"
	"time"
)

type Return struct {
	Success bool
	Data    interface{}
	Message string
}

type responce struct {
	Name string
}

func GetPlayload(r *http.Request, w http.ResponseWriter, p interface{}) {
	SetupResponse(&w, r)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(p)
	if err != nil {
		ResFail(w, nil, "Not a valid Input")
		return
	}
}

func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	//(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	if (*req).Method == "OPTIONS" {
		return
	}
}


func CheckSettionid(w *http.ResponseWriter, req *http.Request)bool{
	isValidSession:=false
	type GetSessionTime struct{
		Sessionid string
		PresentTime string
		ExpireTime string
	}
	repos:=[]GetSessionTime{}
	Layout:="2006-01-02 15:04:05"
	if req.Header.Get("Sessionid")!=""{
		Sessionid:=req.Header.Get("Sessionid")
		db := orm.Db()
		rows, err := db.Query(`select sessionid,presenttime,expiretime from sessionmaster where sessionid='` + Sessionid + `'`)
		if err != nil {
			fmt.Println("err",err)
		}
		defer rows.Close()
		for rows.Next() {
			repo := GetSessionTime{}
			err = rows.Scan(
				&repo.Sessionid,
				&repo.PresentTime,
				&repo.ExpireTime,
			)
			if err != nil {
				
			}
			repos = append(repos, repo)
		}
		err = rows.Err()
		if err != nil {
			
		}
		

		GetExpireTime:=FormatTime("2006-01-02 15:04:05",repos[0].ExpireTime)
		PresentTime:=time.Now()
		fmt.Println("GetExpireTime",GetExpireTime)
		fmt.Println("PresentTime",PresentTime)
		if PresentTime.Before(GetExpireTime){
			isValidSession =true
			GetExpireTime.Add(time.Minute * 10)
			rowsupdate, err2 := db.Query(`UPDATE sessionmaster SET presenttime='`+ PresentTime.Format(Layout) + `',expiretime='`+ GetExpireTime.Format(Layout) + ` ' WHERE sessionid = '` + repos[0].Sessionid + `';`)
			if err2 != nil {
				
			}
			defer rowsupdate.Close()
		}
	}

	return isValidSession

}


func FormatTime(layout string,Value string)time.Time{
	var NewTime time.Time
	GetTime,_:=time.Parse(layout,strings.TrimSpace(Value))
	loc, _ := time.LoadLocation("Asia/Kolkata")

	NewTime =time.Date(GetTime.Year(),GetTime.Month(),GetTime.Day(),GetTime.Hour(),GetTime.Minute(),GetTime.Second(),GetTime.Nanosecond(),loc)

	return NewTime
}

func ResSuccess(w http.ResponseWriter, Data interface{}, m string) {
	res := Return{}
	res.Success = true
	res.Message = m
	res.Data = Data
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(res)
}

func ResFail(w http.ResponseWriter, Data interface{}, m string) {
	res := Return{}
	res.Success = false
	res.Message = m
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(res)
}

type Users struct {
	Lastname  string
	Firstname string
	Username  string
	Gender    string
	Roalval   string
	IsActive  string
	Password  string
}

type ClientUsers struct{
	Lastname  string
	Firstname string
	Username  string
	Email     string
	Gender    string
	Roalval   string
	IsActive  string
	Password  string
	WalletAddress string
}

type Show struct {
	ID       int64
	Time     string
	IsActive bool
}

type Distributor struct {
	ID            int64
	Firsrname     string
	Lastname      string
	Address       string
	City          string
	Email         string
	DistributorID string
	Mobile        string
	Joined        time.Time
	IsActive      bool
}

type DailyData struct {
	ID            int64
	Date          time.Time
	Showid        string
	Distributorid string
	A             int64
	SetA          int64
	B             int64
	SetB          int64
	C             int64
	SetC          int64
	AB            int64
	SetAB         int64
	AC            int64
	SetAC         int64
	BC            int64
	SetBC         int64
	ABCFULL       int64
	SetABCFULL    int64
	ABCHALF       int64
	SetABCHALF    int64
	BOXFULL       int64
	SetBOXFULL    int64
	BOXHALF       int64
	SetBOXHALF    int64
	Status        string
}

type DailyApproval struct {
	ID            int64
	Date          time.Time
	Showid        string
	Showidint     int64
	Distributorid string
	Appliedby     string
	Approvedby    string
	Status        string
}

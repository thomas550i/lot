package main

import (
	"DataLotApi/api/handlers"
	orm "DataLotApi/db"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	orm.InitDb()
	defer orm.Close()
	http.HandleFunc("/users/login", handlers.IndexHandler)
	http.HandleFunc("/users/getusers", handlers.GetUsers)
	http.HandleFunc("/users/saveuser", handlers.Saveuser)
	http.HandleFunc("/users/deleteuser", handlers.DeleteUser)
	http.HandleFunc("/users/updateuser", handlers.Updateuser)

	http.HandleFunc("/shows/get", handlers.Get_Show)
	http.HandleFunc("/shows/save", handlers.Save_Show)
	http.HandleFunc("/shows/delete", handlers.Delete_Show)
	http.HandleFunc("/shows/update", handlers.Update_Show)

	http.HandleFunc("/distributor/get", handlers.Get_Distributor)
	http.HandleFunc("/distributor/save", handlers.Save_Distributor)
	http.HandleFunc("/distributor/delete", handlers.Delete_Distributor)
	http.HandleFunc("/distributor/update", handlers.Update_Distributor)

	http.HandleFunc("/dailydata/get", handlers.Get_DailyData)
	http.HandleFunc("/dailydata/save", handlers.Save_DailyData)
	http.HandleFunc("/dailydata/delete", handlers.Delete_DailyData)
	http.HandleFunc("/dailydata/update", handlers.Update_DailyData)
	http.HandleFunc("/dailydata/submitforapproval", handlers.Submitforapproval)
	http.HandleFunc("/dailydata/getdailystatus", handlers.Get_DailyStatus)
	http.HandleFunc("/dailydata/dailyapproval", handlers.Get_DailyApproval)
	http.HandleFunc("/dailydata/approveall", handlers.Daily_approveall)
	http.HandleFunc("/dailydata/findnumber", handlers.Get_Findnumber)

	log.Fatal(http.ListenAndServe("0.0.0.0:8111", nil))
}

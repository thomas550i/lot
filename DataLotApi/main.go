package main

import (
	"DataLotApi/api/handlers"
	orm "DataLotApi/db"
	"log"
	"net/http"

	// cors "github.com/rs"

	_ "github.com/lib/pq"
)

func main() {
	orm.InitDb()
	defer orm.Close()
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.Write([]byte("{\"hello\": \"world\"}"))
	// })
	http.HandleFunc("/users/login", handlers.IndexHandler)
	http.HandleFunc("/users/getusers", handlers.GetUsers)
	http.HandleFunc("/users/saveuser", handlers.Saveuser)
	http.HandleFunc("/users/saveclientuser", handlers.SaveClientUser)
	http.HandleFunc("/users/checkduplicatemail", handlers.CheckClientDuplicateMail)
	http.HandleFunc("/users/processotpuser", handlers.ProcessOtpUser)
	http.HandleFunc("/users/loginclientuser", handlers.LoginClientUser)
	http.HandleFunc("/users/shoppingcart", handlers.ShoppingCart)


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
	// handler := cors.Default().Handler(mux)
	// http.ListenAndServe("0.0.0.0:8111", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8111", nil))
}

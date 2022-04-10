package main

import (
	"DataLotApi/api/handlers"
	orm "DataLotApi/db"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/crypto/acme/autocert"
	cors "github.com/rs"

	_ "github.com/lib/pq"
)

func main() {
	orm.InitDb()
	defer orm.Close()
//	var m *autocert.Manager
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write([]byte("{\"hello\": \"world\"}"))
	// })

	//handler := cors.Default().Handler(mux)

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.Write([]byte("{\"hello\": \"world\"}"))
	// })
	mux.HandleFunc("/users/pupdate", handlers.Pupdate)
	mux.HandleFunc("/users/login", handlers.IndexHandler)
	mux.HandleFunc("/users/getusers", handlers.GetUsers)
	mux.HandleFunc("/users/saveuser", handlers.Saveuser)
	mux.HandleFunc("/users/saveclientuser", handlers.SaveClientUser)
	mux.HandleFunc("/users/edituserprofile",handlers.EditUserProfile)
	mux.HandleFunc("/users/checkduplicatemail", handlers.CheckClientDuplicateMail)
	mux.HandleFunc("/users/processotpuser", handlers.ProcessOtpUser)
	mux.HandleFunc("/users/loginclientuser", handlers.LoginClientUser)
	mux.HandleFunc("/users/addinshoppingcart", handlers.AddToCart)
	mux.HandleFunc("/users/exisitingcart", handlers.GetExisitingCart)
	mux.HandleFunc("/users/deletecartitem", handlers.DeleteCartItem)
	mux.HandleFunc("/users/removeexpiredtickets", handlers.RemoveExpiredTickets)
	mux.HandleFunc("/users/proceedtocheckout", handlers.ProceedToCheckOut)
	mux.HandleFunc("/users/gettransactions", handlers.GetTransactions)
mux.HandleFunc("/users/gettransactionbyid", handlers.GetTransactionById)

	mux.HandleFunc("/users/getdailyshows", handlers.GetDailyShows)

	mux.HandleFunc("/users/deleteuser", handlers.DeleteUser)
	mux.HandleFunc("/users/updateuser", handlers.Updateuser)

	mux.HandleFunc("/shows/get", handlers.Get_Show)
	mux.HandleFunc("/shows/save", handlers.Save_Show)
	mux.HandleFunc("/shows/delete", handlers.Delete_Show)
	mux.HandleFunc("/shows/update", handlers.Update_Show)

	mux.HandleFunc("/distributor/get", handlers.Get_Distributor)
	mux.HandleFunc("/distributor/save", handlers.Save_Distributor)
	mux.HandleFunc("/distributor/delete", handlers.Delete_Distributor)
	mux.HandleFunc("/distributor/update", handlers.Update_Distributor)

	mux.HandleFunc("/dailydata/get", handlers.Get_DailyData)
	mux.HandleFunc("/dailydata/save", handlers.Save_DailyData)
	mux.HandleFunc("/dailydata/delete", handlers.Delete_DailyData)
	mux.HandleFunc("/dailydata/update", handlers.Update_DailyData)
	mux.HandleFunc("/dailydata/submitforapproval", handlers.Submitforapproval)
	mux.HandleFunc("/dailydata/getdailystatus", handlers.Get_DailyStatus)
	mux.HandleFunc("/dailydata/dailyapproval", handlers.Get_DailyApproval)
	mux.HandleFunc("/dailydata/approveall", handlers.Daily_approveall)
	mux.HandleFunc("/dailydata/findnumber", handlers.Get_Findnumber)

	hserver := cors.AllowAll().Handler(mux)
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}

	server := &http.Server{
		Addr:    ":443",
		Handler: hserver,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go http.ListenAndServe(":8111", certManager.HTTPHandler(nil))
	err:=server.ListenAndServeTLS("", "")
	if err!=nil{
		fmt.Println("Error  ----- ",err)
	}
}

// func cacheDir() (dir string) {
//         if u, _ := user.Current(); u != nil {
//                 dir = filepath.Join(os.TempDir(), "cache-golang-autocert-"+u.Username)
//                 if err := os.MkdirAll(dir, 0700); err == nil {
//                         return dir
//                 }
//         }
//         return ""
// }

package main

import (
	"DataLotApi/api/handlers"
	orm "DataLotApi/db"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/crypto/acme/autocert"
	cors "github.com/rs"

	_ "github.com/lib/pq"
)

func main() {
	orm.InitDb()
	defer orm.Close()
	var m *autocert.Manager
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
	// http.ListenAndServe("0.0.0.0:8111", hserver)



	hostPolicy := func(ctx context.Context, host string) error {
		// Note: change to your real host
		allowedHost := "www.eodmarket.com"
		if host == allowedHost {
			return nil
		}
		return fmt.Errorf("acme/autocert: only %s host is allowed", allowedHost)
	}

	dataDir := "."
	m = &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: hostPolicy,
		Cache:      autocert.DirCache(dataDir),
	}
	
	srv := &http.Server{
			Addr:         ":443",
			Handler:      hserver,
			TLSConfig:    &tls.Config{GetCertificate: m.GetCertificate},
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
		fmt.Println(srv.ListenAndServeTLS("tls.crt", "tls.key"))
		err := srv.ListenAndServeTLS("/etc/ssl/certs/ca-certificates.crt", "/etc/ssl/private/apache-selfsigned.key")
		if err != nil {
			log.Fatalf("httpsSrv.ListendAndServeTLS() failed with %s", err)
		}
	//log.Fatal(http.ListenAndServe("0.0.0.0:8111", nil))
}

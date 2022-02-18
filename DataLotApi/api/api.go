package api

import (
	"net/http"
	"DataLotApi/api/handlers"
)

http.HandleFunc("/api/index", handlers.IndexHandler)
http.HandleFunc("/api/repo/", handlers.)
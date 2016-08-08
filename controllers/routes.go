package controllers

import (
	"net/http"
	//	"github.com/vibhavp/i58-portal/models"
	"github.com/vibhavp/i58-portal/controllers/api"
	_ "github.com/vibhavp/i58-portal/controllers/api/admin"
	_ "github.com/vibhavp/i58-portal/controllers/api/stats"
)

func init() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./views/draw.html")
	})
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./views/admin.html")
	})
	http.HandleFunc("/api/login", api.Login)

	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir("./views/assets"))))
	http.Handle("/js/", http.StripPrefix("/js/",
		http.FileServer(http.Dir("./views/js"))))
}
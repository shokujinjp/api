package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shokujinjp/shokujinjp-sdk-go/shokujinjp"
)

const location = "Asia/Tokyo"

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"health": "ok"}`)
}

func menuAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	all, err := shokujinjp.GetMenuAllData()
	if err != nil {
		log.Fatal(err)
	}

	jb, err := json.Marshal(all)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(jb))
}

func menuToday(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	today, err := shokujinjp.GetMenuDateData(time.Now())
	if err != nil {
		log.Fatal(err)
	}

	sortedToday := shokujinjp.SortByCategory(today)

	jb, err := json.Marshal(sortedToday)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(jb))

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	router := mux.NewRouter()
	router.Path("/").HandlerFunc(index)

	rMenu := router.PathPrefix("/menu").Subrouter()
	rMenu.Path("/all").HandlerFunc(menuAll)
	rMenu.Path("/today").HandlerFunc(menuToday)

	if err := http.ListenAndServe(":"+port, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

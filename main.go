package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/shokujinjp/shokujinjp-sdk-go/shokujinjp"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"health": "ok"}`)
}

func menuAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	all, err := shokujinjp.GetAllMenuData()
	if err != nil {
		log.Fatal(err)
	}

	jb, err := json.Marshal(all)
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

	router := mux.NewRouter()
	router.Path("/").HandlerFunc(index)

	rMenu := router.PathPrefix("/menu").Subrouter()
	rMenu.Path("/all").HandlerFunc(menuAll)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

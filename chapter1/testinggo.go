package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func testHandler(w http.ResponseWriter, r *http.Request){
	
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/",testHandler)

	http.Handle("/",router)

	fmt.Println("Everything is working fine !!")
}

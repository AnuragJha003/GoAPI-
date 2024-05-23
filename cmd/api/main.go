package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"  //useful for web development 
	"github.com/Anurag003/goapi/internal/handlers" //the handler function apna hi h yeh 
	log "github.com/sirupsen/logrus"//and the log external library imported make sure it is installed before import
)


func main(){

	log.SetReportCaller(true)//setting a logger 
	var r *chi.Mux = chi.NewRouter() //returns a pointer to the mux type and helps in creation of a new router 
	//much like struct 
	handlers.Handler(r)//this r is passed onto the Handler function for processing which is implemented in the internals 

	fmt.Println("Starting GO API service...")

  err := http.ListenAndServe("localhost:8000", r) //starting the http server and listenandserve takes the location or port where we want to start the server 
  //and the r which is passed onto the handler 
  if err != nil {
	  log.Error(err)
  }
	
}
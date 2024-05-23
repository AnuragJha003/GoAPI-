package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"  //useful for web development 
	"github.com/Anurag003/goapi/internal/handlers" //the handler function apna hi h yeh 
	log "github.com/sirupsen/logrus"//and the log external library imported make sure it is installed before import
)


func main(){

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

  err := http.ListenAndServe("localhost:8000", r)
  if err != nil {
	  log.Error(err)
  }
	
}
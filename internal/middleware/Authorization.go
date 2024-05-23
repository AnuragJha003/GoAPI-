package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/avukadin/goapi/api"
	"github.com/avukadin/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var unauthorized=errors.New(fmt.Sprintf("Invalid username or token"))

func Authorization(next http.Handler) http.Handler{ //middleware signature in common 
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){//it takes in an httphandler and returns an http handler in return 
		//responsewriter is used to construct a response and the request is used for returning a request all information 
		var username string=r.URL.Query().Get("username")
		var token=r.Header.Get("Authorization")//taking username and tokwn from the header and url 
		var err error 
		if username== "" {
			api.RequestErrorHandler(w,unauthorized)return//error construction 
		}
		var database *tools.DatabaseInterface  //instantiate a new database interface 
		database,err=tools.NewDatabase()//and create a new database 
		if err!=nil{
			api.InternalErrorHandler(w)
			return 
		}
		var loginDetails *tools.LoginDetails 
		loginDetails=(*database).GetUserLoginDetails(username)//take the logindetails and extract it fromm the database based on the username provided in the url and params 

		if(loginDetails==nil || (token != (*loginDetails).AuthToken)){
			log.Error(unauthorized)
			api.RequestErrorHandler(w,unauthorized)
			return
		}//if some error is there 
		next.ServeHTTP(w,r)//else forward it to the the next middleware in line 
		//Middleware1---> next.servehttp--> middleware2-->....-> Handlerfunction
		//Authorization--> next.servehttp->GetCoinBalance 
	})
}
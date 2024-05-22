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

func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		var username string=r.URL.Query().Get("username")
		var token=r.Header.Get("Authorization")
		var err error 
		if username== "" {
			api.RequestErrorHandler(w,unauthorized)return
		}
		var database *tools.DatabaseInterface
		database,err=tools.NewDatabase()
		if err!=nil{
			api.InternalErrorHandler(w)
			return 
		}
		var loginDetails *tools.LoginDetails 
		loginDetails=(*database).GetUserLoginDetails(username)

		if(loginDetails==nil || (token != (*loginDetails).AuthToken)){
			log.Error(unauthorized)
			api.RequestErrorHandler(w,unauthorized)
			return
		}
		next.ServeHTTP(w,r)
	})
}
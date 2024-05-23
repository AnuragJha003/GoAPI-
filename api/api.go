package api 
import(
	"encoding/json"
	"net/http"
)


type CoinBalanceParams struct{
	Username string 
}

type CoinBalanceResponse struct{
	Code int 
	Balance int64
}

type Error struct{
	Code int 
	Message string 
}

func writeError(w http.ResponseWriter, message string, code int){ //writeerror function for returning error in general 
	resp:=Error{
		Code:code,
		Message:message,
	}
	w.Header().Set("Content-Type","application/json")
	w.writeHeader(code)//writing the header in there 

	json.newEncoder(w).Encode(resp)//and encoding the header from the json in there 
}

var(
	RequestErrorHandler=func(w http.ResponseWriter, err error){
		writeError(w,err.Error(),http.StatusBadRequest)
	}//specific error message 
	InternalErrorHandler=func(w http.ResponseWriter){
		writeError(w,"An unexpected eror",http.StatusInternalServerError)
	}//a more general error message in here returned 
)
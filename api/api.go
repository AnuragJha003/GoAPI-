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

func writeError(w http.ResponseWriter, message string, code int){
	resp:=Error{
		Code:code,
		Message:message,
	}
	w.Header().Set("Content-Type","application/json")
	w.writeHeader(code)

	json.newEncoder(w).Encode(resp)
}

var(
	RequestErrorHandler=func(w http.ResponseWriter, err error){
		writeError(w,err.Error(),http.StatusBadRequest)
	}//specific 
	InternalErrorHandler=func(w http.ResponseWriter){
		writeError(w,"An unexpected eror",http.StatusInternalServerError)
	}
)
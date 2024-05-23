package handlers
//for the router get function in here 
import (
	"encoding/json"
	"net/http"

	"github.com/avukadin/goapi/api"
	"github.com/avukadin/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {//response and request 
	var params = api.CoinBalanceParams{}//getting the params from the coibalanceparams 
	var decoder *schema.Decoder = schema.NewDecoder()//decoding it 
	var err error

	err = decoder.Decode(&params, r.URL.Query())//grab the params from the url and set them in the struct for creating a new coinbalance 

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface //new db interface created 
	database, err = tools.NewDatabase()//new db 
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)//token details (coin wala taken)
	//and extracted from the database (getusercoins)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}//setting it in the repsonse 

	w.Header().Set("Content-Type", "application/json")//and header m 
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
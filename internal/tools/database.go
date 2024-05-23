package tools 

import(
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct{
	AuthToken string 
	Username string 
}

type CoinDetails struct{
	Coins int64 
	Username string 
}
//struct consisting of the schema of the datas in the database 
type DatabaseInterface interface{
	GetUserLoginDetails(username string) *LoginDetails
	GetCoinBalance(username string) *CoinDetails
	SetupDatabase() error
}//interface defined along with the data types 

func NewDatabase() (*DatabaseInterface, error){//creating a new database with the interface passed as well
	var database DatabaseInterface=& mockDB{}//new databaseinterface pointed to the address of the mockdb 
	var err error=database.SetupDatabase()//error setup m 
	if err!=nil{
		log.Error(err)
		return nil,err
	}

	return &database,nil//return the database mock db has been created 

}
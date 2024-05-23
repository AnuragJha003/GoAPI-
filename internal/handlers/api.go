package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"  //importing the middleware from the chimiddle library 
	"github.com/Anurag003/goapi/internal/middleware"//yeh toh apna hi h 
)
//small handler rheta toh it is the private variable and function 
func Handler(r *chi.Mux){
	//global middleware 
	r.Use(chimiddle.StripSlashes) //stripslashes end m route slash stripped (use middlleware)

	r.Route("/account",func(router chi.Router){ //path for the route and an anonumous function 
		//middleware for /account route 
		router.Use(middleware.Authorization)//use another middleware which we create 
		router.Get("/coins",GetCoinBalance)//else we get the particular coins path which will be handled by the function attached to it 
	})
}
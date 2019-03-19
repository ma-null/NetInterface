package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/ma-null/NetInterface/handlers"
	
	"net/http"

)

func main() {
	
	router := httprouter.New()
	router.GET("/service/version", handlers.GetVersion)
	router.GET("/service/"+handlers.ApiVersion+"/interfaces", handlers.GetIntefaces)
	router.GET("/service/"+handlers.ApiVersion+"/interface/:name", handlers.GetInteface)
	
	fmt.Println(http.ListenAndServe(":8080", router))
	
}
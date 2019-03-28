package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/ma-null/NetInterface/handlers"
	"net/http"
)

func main() {
	router := httprouter.New()
	mng := handlers.NewNetInterfaceManger()
	router.GET("/service/version", handlers.GetVersion)
	router.GET("/service/"+handlers.ApiVersion+"/interfaces", handlers.GetInterfaces(mng))
	router.GET("/service/"+handlers.ApiVersion+"/interface/:name", handlers.GetInterface(mng))
	
	fmt.Println(http.ListenAndServe(":8080", router))
}
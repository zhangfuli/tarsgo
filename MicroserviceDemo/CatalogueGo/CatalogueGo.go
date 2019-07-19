package main

import (
	"MicroserviceDemo/CatalogueGo/MicroserviceDemo"
	imp2 "MicroserviceDemo/CatalogueGo/imp"
	"github.com/TarsCloud/TarsGo/tars"

)

func main() { //Init servant
	imp := new(imp2.DoCatalogueImp)                               //New Imp
	app := new(MicroserviceDemo.DoCatalogue)                      //New init the A Tars
	cfg := tars.GetServerConfig()                                 //Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".DoCatalogueObj") //Register Servant
	tars.Run()
}

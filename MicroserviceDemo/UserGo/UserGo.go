package main

import (
	"MicroserviceDemo/UserGo/MicroserviceDemo"
	imp2 "MicroserviceDemo/UserGo/imp"
	"github.com/TarsCloud/TarsGo/tars"
)

func main() { //Init servant
	imp := new(imp2.DoUserImp)                               //New Imp
	app := new(MicroserviceDemo.DoUser)                      //New init the A Tars
	cfg := tars.GetServerConfig()                            //Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".DoUserObj") //Register Servant
	tars.Run()
}

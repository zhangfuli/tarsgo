package main

import (
	"MicroserviceDemo/PaymentGo/MicroserviceDemo"
	imp2 "MicroserviceDemo/PaymentGo/imp"
	"github.com/TarsCloud/TarsGo/tars"
)

func main() { //Init servant
	imp := new(imp2.DoPaymentImp)                               //New Imp
	app := new(MicroserviceDemo.DoPayment)                      //New init the A Tars
	cfg := tars.GetServerConfig()                               //Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".DoPaymentObj") //Register Servant
	tars.Run()
}

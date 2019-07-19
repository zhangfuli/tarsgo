package main

import (
	"MicroserviceDemo/PaymentGo/MicroserviceDemo"
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"
)

var comm *tars.Communicator
func main() {
	comm = tars.NewCommunicator()
	obj := "TestApp.HelloGo.SayHelloObj@tcp -h 127.0.0.1 -p 10015 -t 60000"
	app := new(MicroserviceDemo.DoPayment)
	comm.StringToProxy(obj, app)

	health, healthErr := app.Health()
	authorisation, authorisationErr := app.Authorise(80)
	if healthErr != nil {
		fmt.Println(healthErr)
		return
	}
	fmt.Println("resp: ", health)
	if authorisationErr != nil {
		fmt.Println(authorisationErr)
		return
	}
	fmt.Println("resp: ", authorisation)
}

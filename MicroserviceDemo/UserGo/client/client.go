package main

import (
	"MicroserviceDemo/UserGo/MicroserviceDemo"
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"

)

func main() {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("MicroserviceDemo.UserGo.DoUserObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	app := new(MicroserviceDemo.DoUser)
	comm.StringToProxy(obj, app)

	findCustomerCardByIdRet, findCustomerCardByIdErr := app.FindAddressById("57a98ddce4b00679b4a830d1")
	if findCustomerCardByIdErr != nil {
		fmt.Println(findCustomerCardByIdErr)
		return
	}
	fmt.Println("ret: ", findCustomerCardByIdRet)
}

package main

import (
	"MicroserviceDemo/CatalogueGo/MicroserviceDemo"
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"


)

func main() {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("MicroserviceDemo.CatalogueGo.DoCatalogueObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	app := new(MicroserviceDemo.DoCatalogue)
	comm.StringToProxy(obj, app)

	catalogue, catalogueErr := app.Catalogue()
	if catalogueErr != nil {
		fmt.Println(catalogueErr)
		return
	}
	fmt.Println("resp: ", catalogue)

	catalogueId, catalogueIdErr := app.CatalogueId("3395a43e-2d88-40de-b95f-e00e1502085b")
	if catalogueIdErr != nil {
		fmt.Println(catalogueIdErr)
		return
	}
	fmt.Println("resp: ", catalogueId)

	sizeId, sizeIdErr := app.Count()
	if sizeIdErr != nil {
		fmt.Println(sizeIdErr)
		return
	}
	fmt.Println("resp: ", sizeId)

	tags, tagsErr := app.Tags()
	if tagsErr != nil {
		fmt.Println(tagsErr)
		return
	}
	fmt.Println("resp: ", tags)
}

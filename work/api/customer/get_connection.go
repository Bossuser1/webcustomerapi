package customer

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/webcustomerapi/models"
	"fmt"
	"strconv"
	api "github.com/webcustomerapi/restapi/operations/customer"
)


//
//
//
var DoGetConnection= func(p api.GetconnectionbyidParams) middleware.Responder{

	fmt.Println(p.ID)
	//t1:=string(p.ID)
	//t1, err := strconv.Atoi(p.ID) strconv.Itoa(97)
	t1:=strconv.Itoa(int(p.ID))
	t2:="string"

	return api.NewGetconnectionbyidOK().WithPayload(&models.Connection{DateConnect: &t1,IP: &t2})
}

/*
Sniperkit-Bot
- Status: analyzed
*/

//go:generate go run $GOPATH/src/github.com/TIBCOSoftware/flogo-lib/flogo/gen/gen.go $GOPATH
package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/TIBCOSoftware/flogo-contrib/trigger/rest"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/engine"
	"github.com/TIBCOSoftware/flogo-lib/flogo"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"github.com/sniperkit/snk.fork.retgits-flogo-components/activity/addtodate"
	"github.com/sniperkit/snk.fork.retgits-flogo-components/activity/randomnumber"
)

var (
	httpport = os.Getenv("HTTPPORT")
)

func main() {
	// Create a new Flogo app
	app := appBuilder()

	e, err := flogo.NewEngine(app)

	if err != nil {
		logger.Error(err)
		return
	}

	engine.RunEngine(e)
}

func appBuilder() *flogo.App {
	app := flogo.NewApp()

	// Convert the HTTPPort to an integer
	port, err := strconv.Atoi(httpport)
	if err != nil {
		logger.Error(err)
	}

	// Register the HTTP trigger
	trg := app.NewTrigger(&rest.RestTrigger{}, map[string]interface{}{"port": port})
	trg.NewFuncHandler(map[string]interface{}{"method": "GET", "path": "/api/expected-date/:invoiceId"}, Handler)
	trg.NewFuncHandler(map[string]interface{}{"method": "GET", "path": "/swaggerspec"}, SwaggerSpec)

	return app
}

// SwaggerSpec is the function that gets executedto retrieve the SwaggerSpec
func SwaggerSpec(ctx context.Context, inputs map[string]*data.Attribute) (map[string]*data.Attribute, error) {
	// The return message is a map[string]*data.Attribute which we'll have to construct
	response := make(map[string]interface{})
	ret := make(map[string]*data.Attribute)

	fileData, err := ioutil.ReadFile("swagger.json")
	if err != nil {
		ret["code"], _ = data.NewAttribute("code", data.TypeInteger, 500)
		response["msg"] = err.Error()
	} else {
		ret["code"], _ = data.NewAttribute("code", data.TypeInteger, 200)
		var data map[string]interface{}
		if err := json.Unmarshal(fileData, &data); err != nil {
			panic(err)
		}
		response = data
	}

	ret["data"], _ = data.NewAttribute("data", data.TypeAny, response)

	return ret, nil

}

// Handler is the function that gets executed when the engine receives a message
func Handler(ctx context.Context, inputs map[string]*data.Attribute) (map[string]*data.Attribute, error) {
	// Get the ID from the path
	id := inputs["pathParams"].Value().(map[string]string)["invoiceId"]

	// Generate a random number for the amount
	// There are definitely better ways to do this with Go, but this keeps the flow consistent with the UI version
	in := map[string]interface{}{"min": 0, "max": 10}
	out, err := flogo.EvalActivity(&randomnumber.MyActivity{}, in)
	if err != nil {
		return nil, err
	}
	datediff := out["result"].Value().(int)

	// Generate a new date
	// There are definitely better ways to do this with Go, but this keeps the flow consistent with the UI version
	in = map[string]interface{}{"number": datediff, "units": "days"}
	out, err = flogo.EvalActivity(&addtodate.MyActivity{}, in)
	if err != nil {
		return nil, err
	}
	expectedPaymentDate := out["result"].Value().(string)

	// The return message is a map[string]*data.Attribute which we'll have to construct
	response := make(map[string]interface{})
	response["id"] = id
	response["expectedDate"] = expectedPaymentDate

	ret := make(map[string]*data.Attribute)
	ret["code"], _ = data.NewAttribute("code", data.TypeInteger, 200)
	ret["data"], _ = data.NewAttribute("data", data.TypeAny, response)

	return ret, nil
}

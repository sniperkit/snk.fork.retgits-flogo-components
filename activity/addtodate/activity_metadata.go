/*
Sniperkit-Bot
- Status: analyzed
*/

package addtodate

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
    "name": "addtodate",
    "type": "flogo:activity",
    "ref": "github.com/sniperkit/snk.fork.retgits-flogo-components/activity/addtodate",
    "version": "0.1.0",
    "title": "Add to Date",
    "description": "Add a specified number of units to a date",
    "author": "retgits",
    "homepage": "https://github.com/sniperkit/snk.fork.retgits-flogo-components/tree/master/activity/addtodate",
    "inputs":[
      {
        "name": "number",
        "type": "integer"
      },
      {
        "name": "units",
        "type": "string",
        "allowed" : ["years", "months", "days"]
      },
      {
        "name": "date",
        "type": "string"
      }
    ],
    "outputs": [
      {
        "name": "result",
        "type": "string"
      }
    ]
  }`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}

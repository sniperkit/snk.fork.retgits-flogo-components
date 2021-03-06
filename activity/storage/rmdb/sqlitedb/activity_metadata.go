package sqlitedb

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "RMDB - Sqlite3",
  "version": "0.0.1",
  "title": "SQLite DB Operations",
  "type": "flogo:activity",
  "ref": "github.com/sniperkit/snk.fork.retgits-flogo-components/activity/storage/rmdb/sqlitedb",
  "description": "Create, insert, update & delete from SQLite DB v3.x",
  "author": "Rosco Pecoltran <sniperkit@protonmail.com>",
  "input":[
    {
      "name": "database",
      "type": "string",
      "required": true
    },
    {
      "name": "query",
      "type": "string",
      "required": true
    },
    {
      "name": "parameters",
      "type": "params"
    },
    {
      "name": "prefix_path",
      "type": "string",
      "required": true,
      "value": "~"
    }
  ],
  "output": [
    {
      "name": "result",
      "type": "any"
    }
  ]
}

`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
